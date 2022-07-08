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
    borrow::Borrow,
    fs::File,
    io::{self, Read},
};

use clap::{Args, Parser, Subcommand};
use lazy_static::lazy_static;
use rbatis::rbatis::Rbatis;
lazy_static! {
  // Rbatis是线程、协程安全的，运行时的方法是Send+Sync，无需担心线程竞争
  pub static ref DB:Rbatis=Rbatis::new();
}

lazy_static! {
    pub static ref CLI: AppCli = AppCli::new();
}

lazy_static! {
    pub static ref CONFIG: ApplicationConfig = ApplicationConfig::new();
}

#[derive(Debug, Parser)]
#[clap(name = "oab")]
#[clap(about = "oab", long_about = None)]
pub struct AppCli {
    #[clap(short = 'c', value_name = "cfg",default_value_t = String::from("~/.config/oab/oab.yml"), value_hint = clap::ValueHint::DirPath)]
    cfg: String,
    #[clap(subcommand)]
    pub command: Option<Clis>,
}

#[derive(Debug, Subcommand)]
pub enum Clis {
    Init,
    Web,
    Stash(StashData),
}

#[derive(Debug, Args)]
#[clap(args_conflicts_with_subcommands = true)]
pub struct StashData {
    command: Option<String>,
}

impl AppCli {
    fn new() -> Self {
        AppCli::parse()
    }
}

#[derive(Debug, PartialEq, serde::Serialize, serde::Deserialize)]
pub struct ApplicationConfig {
    pub debug: bool,
    pub server_url: String,
    pub db_url: String,
    pub db_user: String,
    pub db_pass: String,
    pub db_name: String,
    pub log_dir: Option<String>,
    /// "100MB" 日志分割尺寸-单位KB,MB,GB
    pub log_temp_size: Option<String>,
    pub log_pack_compress: Option<String>,
    pub log_level: Option<String>,
    pub jwt_secret: Option<String>,
}

impl ApplicationConfig {
    pub fn new() -> Self {
        let mut f = match File::open(CLI.cfg.clone()) {
            Ok(f) => f,
            Err(ref e) if e.kind() == io::ErrorKind::NotFound => return Self::defaut(),
            Err(e) => panic!("{}", e),
        };
        File::open(CLI.cfg.clone()).unwrap();
        let mut yml_data = String::new();
        f.read_to_string(&mut yml_data).unwrap();
        //读取配置
        let result: ApplicationConfig =
            serde_yaml::from_str(&yml_data).expect("load config file fail");
        if result.debug {
            println!("load config:{:?}", result);
            println!("///////////////////// Start On Debug Mode ////////////////////////////");
        } else {
            println!("release_mode is enable!")
        }
        result
    }
    pub fn defaut() -> Self {
        Self {
            debug: true,
            server_url: "127.0.0.1:4001".to_string(),
            db_url: "127.0.0.1:3306".to_string(),
            db_user: "root".to_string(),
            db_pass: "123456".to_string(),
            db_name: "test".to_string(),
            log_dir: None,
            log_temp_size: None,
            log_pack_compress: None,
            log_level: None,
            jwt_secret: None,
        }
    }
    pub fn save(&self) {}
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
    tracing_subscriber::fmt()
        .with_line_number(true)
        .with_timer(FormatTime {})
        .init();
}
