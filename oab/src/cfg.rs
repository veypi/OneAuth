//
// cfg.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// Distributed under terms of the MIT license.
//
//
//
//
//

use std::{
    fs::File,
    io::{self, Read},
    time::Duration,
};

use sea_orm::{ConnectOptions, Database, DatabaseConnection};
use serde::{Deserialize, Serialize};
use sqlx::{mysql::MySqlPoolOptions, Pool};
use tracing::Level;

use crate::Result;

#[derive(Debug, serde::Serialize, serde::Deserialize, Clone)]
pub struct ApplicationConfig {
    #[serde(skip)]
    pub db: DatabaseConnection,
}

#[derive(Debug, Clone, Deserialize, Serialize, Default)]
pub struct InfoOpt {
    pub nats_url: String,
    pub ws_url: String,
    pub api_url: String,
    pub ts_url: String,
    pub token: Option<String>,
}

#[derive(Debug, Clone, serde::Serialize, serde::Deserialize)]
pub struct AppState {
    pub uuid: String,
    pub key: String,
    pub debug: bool,
    pub server_url: String,
    pub media_path: String,
    pub db_url: String,
    pub db_user: String,
    pub db_pass: String,
    pub db_name: String,
    pub log_dir: Option<String>,
    pub auto_task: bool,
    pub fs_root: String,
    pub ts_url: String,
    pub nats_url: String,
    pub nats_usr: [String; 2],
    pub nats_node: [String; 2],
    pub nats_sys: [String; 2],

    pub info: InfoOpt,

    pub log_level: Option<String>,
    pub user_init_space: i64,

    #[serde(skip)]
    pub _sqlx: Option<Pool<sqlx::MySql>>,
    #[serde(skip)]
    pub _db: Option<DatabaseConnection>,
}

impl AppState {
    pub fn new(cli_path: &str) -> Self {
        let mut res = Self::defaut();
        let mut f = match File::open(cli_path) {
            Ok(f) => f,
            Err(ref e) if e.kind() == io::ErrorKind::NotFound => {
                // res.connect_sqlx().unwrap();
                return res;
            }
            Err(e) => panic!("{}", e),
        };
        File::open(cli_path).unwrap();
        let mut yml_data = String::new();
        f.read_to_string(&mut yml_data).unwrap();
        //读取配置
        res = serde_yaml::from_str(&yml_data).expect("load config file fail");
        res
    }
    pub fn defaut() -> Self {
        Self {
            uuid: "FR9P5t8debxc11aFF".to_string(),
            key: "AMpjwQHwVjGsb1WC4WG6".to_string(),
            debug: true,
            server_url: "127.0.0.1:4001".to_string(),
            db_url: "localhost:3306".to_string(),
            db_user: "root".to_string(),
            db_pass: "123456".to_string(),
            db_name: "oneauth".to_string(),
            log_dir: None,
            auto_task: true,
            media_path: "/Users/veypi/test/media".to_string(),
            fs_root: "/Users/veypi/test/media".to_string(),
            log_level: None,
            _sqlx: None,
            _db: None,
            ts_url: "127.0.0.1:8428".to_string(),
            nats_url: "127.0.0.1:4222".to_string(),
            nats_usr: [
                String::from("UCXFAAVMCPTATZUZX6H24YF6FI3NKPQBPLM6BNN2EDFPNSUUEZPNFKEL"),
                String::from("SUACQNAAFKDKRBXS62J4JYZ7DWZS7UNUQI52BOFGGBUACHTDHRQP7I66GI"),
            ],
            nats_node: [
                String::from("UAU6HPAHVIQWODQ365HMSHGZPSXJHR35T6ACURR3STGXFZNWXFNG5EA6"),
                String::from("SUACZVC4UWLCKFA3DJFIYO5XYYGPJRQEKCBC773PKCD4TZS52GDU6JJ2JE"),
            ],
            nats_sys: [
                String::from("UCOKXBGDAXXQOR4XUPUJ4O22HZ2A3KQN3JLCCYM3ISSKHLBZJXXQ3NLF"),
                String::from("SUAEILQZDD2UT2ZNR6DCA44YCRKAZDYDOJRUPAUA7AOWFVGSSPFPCLXF24"),
            ],
            user_init_space: 300,
            info: InfoOpt {
                ws_url: "127.0.0.1:4221".to_string(),
                nats_url: "127.0.0.1:4222".to_string(),
                api_url: "127.0.0.1:4001".to_string(),
                ts_url: "127.0.0.1:8428".to_string(),
                token: None,
            },
        }
    }
    pub fn save(&self) {}

    pub fn db(&self) -> &DatabaseConnection {
        match &self._db {
            Some(d) => d,
            None => panic!("failed"),
        }
    }
    pub fn sqlx(&self) -> &sqlx::MySqlPool {
        match &self._sqlx {
            Some(d) => d,
            None => panic!("failed"),
        }
    }

    pub fn connect_sqlx(&mut self) -> Result<()> {
        let url = format!(
            "mysql://{}:{}@{}/{}",
            self.db_user, self.db_pass, self.db_url, self.db_name
        );

        let p = MySqlPoolOptions::new()
            .max_connections(5)
            .connect_lazy(&url)?;
        self._sqlx = Some(p);
        Ok(())
    }
    pub async fn connect(&mut self) -> Result<()> {
        let url = format!(
            "mysql://{}:{}@{}/{}",
            self.db_user, self.db_pass, self.db_url, self.db_name
        );
        let mut opt = ConnectOptions::new(url);
        opt.max_connections(100)
            .min_connections(1)
            .connect_timeout(Duration::from_secs(8))
            .acquire_timeout(Duration::from_secs(8))
            .idle_timeout(Duration::from_secs(8))
            .sqlx_logging(false)
            .max_lifetime(Duration::from_secs(8));

        self._db = Some(Database::connect(opt).await?);

        Ok(())
    }
}

struct FormatTime;
impl tracing_subscriber::fmt::time::FormatTime for FormatTime {
    fn format_time(&self, w: &mut tracing_subscriber::fmt::format::Writer<'_>) -> std::fmt::Result {
        let d = chrono::Local::now();
        w.write_str(&d.format("%Y-%m-%d %H:%M:%S").to_string())
    }
}

pub fn init_log(stat: &AppState) -> Option<tracing_appender::non_blocking::WorkerGuard> {
    let level = stat.log_level.clone().unwrap_or("info".to_string());
    let level = match level.as_str() {
        "trace" => Level::TRACE,
        "debug" => Level::DEBUG,
        "warn" => Level::WARN,
        "error" => Level::ERROR,
        "info" => Level::INFO,
        _ => Level::INFO,
    };
    if let Some(log_dir) = stat.log_dir.clone() {
        let file_appender = tracing_appender::rolling::hourly(log_dir, "oab.log");
        let (non_blocking, _guard) = tracing_appender::non_blocking(file_appender);
        tracing_subscriber::fmt()
            .with_writer(non_blocking)
            .with_line_number(true)
            .with_timer(FormatTime {})
            .with_max_level(level)
            .with_ansi(false)
            .init();
        Some(_guard)
    } else {
        tracing_subscriber::fmt()
            .with_line_number(true)
            .with_timer(FormatTime {})
            .with_max_level(level)
            .init();
        None
    }
}
