use oab::{cfg, models};
use tracing::{info, warn};

#[tokio::main]
async fn main() -> std::io::Result<()> {
    let mut cf = cfg::ApplicationConfig::new();
    cfg::init_log();
    cf.log_dir = "".to_string();
    cf.connect().await;
    info!("{}", cf.server_url);
    warn!("{}", cf.db_url);
    models::init().await;
    println!("Hello, world!");
    return std::io::Result::Ok(());
}
