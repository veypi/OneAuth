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

use clap::{Args, Parser, Subcommand};
use lazy_static::lazy_static;
use sea_orm::{ConnectOptions, Database, DatabaseConnection};
use sqlx::{mysql::MySqlPoolOptions, Pool};
use tracing::Level;

use crate::Result;

lazy_static! {
    pub static ref CLI: AppCli = AppCli::new();
}

// lazy_static! {
//     pub static ref CONFIG: ApplicationConfig = ApplicationConfig::new();
// }

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

#[derive(Debug, serde::Serialize, serde::Deserialize, Clone)]
pub struct ApplicationConfig {
    #[serde(skip)]
    pub db: DatabaseConnection,
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
    pub fs_root: String,
    /// "100MB" 日志分割尺寸-单位KB,MB,GB
    pub log_temp_size: Option<String>,
    pub log_pack_compress: Option<String>,
    pub log_level: Option<String>,
    pub jwt_secret: Option<String>,
    pub user_init_space: i64,

    #[serde(skip)]
    pub _sqlx: Option<Pool<sqlx::MySql>>,
    #[serde(skip)]
    pub _db: Option<DatabaseConnection>,
}

impl AppState {
    pub fn new() -> Self {
        let mut res = Self::defaut();
        let mut f = match File::open(CLI.cfg.clone()) {
            Ok(f) => f,
            Err(ref e) if e.kind() == io::ErrorKind::NotFound => {
                res.connect_sqlx().unwrap();
                return res;
            }
            Err(e) => panic!("{}", e),
        };
        File::open(CLI.cfg.clone()).unwrap();
        let mut yml_data = String::new();
        f.read_to_string(&mut yml_data).unwrap();
        //读取配置
        res = serde_yaml::from_str(&yml_data).expect("load config file fail");
        if res.debug {
            println!("load config:{:?}", res);
            println!("///////////////////// Start On Debug Mode ////////////////////////////");
        } else {
            println!("release_mode is enable!")
        }
        res.connect_sqlx().unwrap();
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
            log_temp_size: None,
            log_pack_compress: None,
            media_path: "/Users/veypi/test/media".to_string(),
            fs_root: "/Users/veypi/test/media".to_string(),
            log_level: None,
            jwt_secret: None,
            _sqlx: None,
            _db: None,
            user_init_space: 300,
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

pub fn init_log() {
    tracing_subscriber::fmt()
        .with_line_number(true)
        .with_timer(FormatTime {})
        .with_max_level(Level::INFO)
        // .with_target(false)
        // .with_file(true)
        .init();
}
