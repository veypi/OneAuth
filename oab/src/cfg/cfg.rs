//
// cfg.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// Distributed under terms of the MIT license.
//
//
//
//
//

use std::borrow::Borrow;

use lazy_static::lazy_static;
use rbatis::rbatis::Rbatis;
lazy_static! {
  // Rbatis是线程、协程安全的，运行时的方法是Send+Sync，无需担心线程竞争
  pub static ref DB:Rbatis=Rbatis::new();
}

#[derive(Debug, PartialEq, serde::Serialize, serde::Deserialize)]
pub struct ApplicationConfig {
    pub debug: bool,
    pub server_url: String,
    pub db_url: String,
    pub db_user: String,
    pub db_pass: String,
    pub db_name: String,
    pub log_dir: String,
    /// "100MB" 日志分割尺寸-单位KB,MB,GB
    pub log_temp_size: String,
    pub log_pack_compress: String,
    pub log_level: String,
    pub jwt_secret: String,
}

impl ApplicationConfig {
    pub fn new() -> Self {
        ApplicationConfig {
            debug: true,
            server_url: "127.0.0.1:4000".to_string(),
            db_url: "127.0.0.1:3306".to_string(),
            db_user: "root".to_string(),
            db_name: "test".to_string(),
            db_pass: "123456".to_string(),
            log_dir: "".to_string(),
            log_temp_size: "".to_string(),
            log_pack_compress: "".to_string(),
            log_level: "".to_string(),
            jwt_secret: "".to_string(),
        }
    }
    pub fn db(&self) -> &DB {
        DB.borrow()
    }
    pub async fn connect(&self) {
        let url = format!(
            "mysql://{}:{}@{}/{}",
            self.db_user, self.db_pass, self.db_url, self.db_name
        );
        DB.link(&url).await.unwrap();
    }
}

///默认配置
impl Default for ApplicationConfig {
    fn default() -> Self {
        let yml_data = include_str!("./mod.rs");
        //读取配置
        let result: ApplicationConfig =
            serde_yaml::from_str(yml_data).expect("load config file fail");
        if result.debug {
            println!("load config:{:?}", result);
            println!("///////////////////// Start On Debug Mode ////////////////////////////");
        } else {
            println!("release_mode is enable!")
        }
        result
    }
}

struct FormatTime;
impl tracing_subscriber::fmt::time::FormatTime for FormatTime {
    fn format_time(&self, w: &mut tracing_subscriber::fmt::format::Writer<'_>) -> std::fmt::Result {
        let d =
            time::OffsetDateTime::now_utc().to_offset(time::UtcOffset::from_hms(8, 0, 0).unwrap());
        w.write_str(&format!(
            "{} {}:{}:{}",
            d.date(),
            d.hour(),
            d.minute(),
            d.second()
        ))
    }
}

pub fn init_log() {
    tracing_subscriber::fmt().with_timer(FormatTime {}).init();
}
