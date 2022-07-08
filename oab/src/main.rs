//
// main.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-07 23:51
// Distributed under terms of the Apache license.
//

use actix_web::{middleware, web, App, HttpServer};
use oab::{api, init_log, models, Clis, Result, CLI, CONFIG};
use tracing::info;

#[tokio::main]
async fn main() -> Result<()> {
    init_log();
    if let Some(c) = &CLI.command {
        match c {
            Clis::Init => {
                CONFIG.connect().await;
                models::init().await;
                return Ok(());
            }
            _ => {}
        };
    };
    CONFIG.connect().await;
    web().await?;
    Ok(())
}
async fn web() -> Result<()> {
    std::env::set_var("RUST_LOG", "info");
    std::env::set_var("RUST_BACKTRACE", "1");
    let serv = HttpServer::new(move || {
        let logger = middleware::Logger::default();
        let app = App::new();
        app.wrap(logger)
            .wrap(middleware::Compress::default())
            .service(web::scope("api").configure(api::routes))
    });
    info!("listen to {}", CONFIG.server_url);
    serv.bind(CONFIG.server_url.clone())?.run().await?;
    Ok(())
}
