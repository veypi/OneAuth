//
// task.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-10-19 01:59
// Distributed under terms of the MIT license.
//
use std::time::{Duration, Instant};

use futures_util::StreamExt;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::process;
use std::sync::{Arc, Mutex};
use sysinfo::{Pid, ProcessExt, System, SystemExt};
use tracing::{info, warn};

#[derive(Debug, Clone, Deserialize, Serialize)]
struct SysInfo {
    client: ClientInfo,
    id: String,
    // server: String,
}
#[derive(Debug, Clone, Deserialize, Serialize)]
struct ClientInfo {
    id: i64,
    acc: String,
    name: String,
    host: String,
}

pub fn start_stats_info(url: String) {
    tokio::spawn(async move {
        let pid = process::id();
        info!("star on pid {}", pid);
        let mut s = System::new_all();
        let pid = Pid::from(pid as usize);
        let props = sysinfo::ProcessRefreshKind::everything();
        let url = format!("{}/api/v1/import/prometheus", url);
        let client = reqwest::Client::new();
        let start = Instant::now();
        loop {
            tokio::time::sleep(std::time::Duration::from_secs(1)).await;
            s.refresh_process_specifics(pid, props);
            if let Some(process) = s.process(pid) {
                let stat_str = format!(
                    "srv_cpu{{i=\"oa\"}} {}\nsrv_mem{{i=\"oa\"}} {}\nsrv_start{{i=\"oa\"}} {}",
                    process.cpu_usage(),
                    process.memory(),
                    start.elapsed().as_secs(),
                );
                match client.post(&url).body(stat_str).send().await {
                    Ok(_) => {}
                    Err(e) => {
                        warn!("{}", e);
                    }
                }
            }
        }
    });
}

pub fn start_nats_online(client: async_nats::client::Client) {
    let db: Arc<Mutex<HashMap<i64, ClientInfo>>> = Arc::new(Mutex::new(HashMap::new()));
    {
        let db = db.clone();
        let client = client.clone();
        tokio::spawn(async move {
            let mut sub = client
                .subscribe("$SYS.ACCOUNT.*.CONNECT".to_string())
                .await
                .unwrap();
            while let Some(msg) = sub.next().await {
                let s = String::from_utf8(msg.payload.to_vec()).unwrap();
                info!("{}", s);
                let inf: SysInfo = serde_json::from_slice(&msg.payload.to_vec()).unwrap();
                info!("add {} {}", inf.client.id, inf.client.name);
                let mut db = db.lock().unwrap();
                db.insert(inf.client.id, inf.client);
            }
        });
    }
    {
        let db = db.clone();
        let client = client.clone();
        tokio::spawn(async move {
            let mut sub = client
                .subscribe("$SYS.ACCOUNT.*.DISCONNECT".to_string())
                .await
                .unwrap();
            while let Some(msg) = sub.next().await {
                // let s = String::from_utf8(msg.payload.to_vec()).unwrap();
                let inf: SysInfo = serde_json::from_slice(&msg.payload.to_vec()).unwrap();
                info!("remove {} {}", inf.client.id, inf.client.name);
                let mut db = db.lock().unwrap();
                db.remove(&inf.client.id);
            }
        });
    };
    tokio::spawn(async move {
        let mut sub = client.subscribe("sys.online".to_string()).await.unwrap();
        while let Some(msg) = sub.next().await {
            // // let s = String::from_utf8(msg.payload.to_vec()).unwrap();
            // let inf: SysInfo = serde_json::from_slice(&msg.payload.to_vec()).unwrap();
            // info!("remove {} {}", inf.client.id, inf.client.name);
            // let mut db = db.lock().unwrap();
            // db.remove(&inf.client.id);
            if let Some(t) = msg.reply {
                let d = {
                    let tmp = db.lock().unwrap();
                    let payload: Vec<ClientInfo> = tmp.iter().map(|(_, c)| c.clone()).collect();
                    serde_json::to_string(&payload).unwrap()
                };
                match client.publish(t, d.into()).await {
                    Ok(_) => {}
                    Err(e) => {
                        warn!("{}", e);
                    }
                };
            }
        }
    });
}

pub fn start_demo() {
    tokio::spawn(async move {
        let mut interval = tokio::time::interval(Duration::from_secs(5));
        interval.tick().await;
        let start = Instant::now();
        println!("time:{:?}", start);
        loop {
            interval.tick().await;
            println!("time:{:?}", start.elapsed());
        }
    });
}
