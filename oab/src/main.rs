//
// main.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-07 23:51
// Distributed under terms of the Apache license.
//
use actix_files as fs;
use actix_web::{
    dev,
    http::StatusCode,
    middleware::{self, ErrorHandlerResponse, ErrorHandlers},
    web::{self},
    App, HttpServer,
};

use oab::{api, init_log, libs, models, AppState, Clis, Result, CLI, CONFIG};
use tracing::{error, info, warn};

#[tokio::main]
async fn main() -> Result<()> {
    init_log();
    if let Some(c) = &CLI.command {
        match c {
            Clis::Init => {
                models::init().await;
                return Ok(());
            }
            _ => {}
        };
    };
    web().await?;
    Ok(())
}
async fn web() -> Result<()> {
    let db = CONFIG.connect().await?;
    let data = AppState { db };
    std::env::set_var("RUST_LOG", "info");
    std::env::set_var("RUST_BACKTRACE", "1");
    let serv = HttpServer::new(move || {
        let logger = middleware::Logger::default();
        let json_config = web::JsonConfig::default()
            .limit(4096)
            .error_handler(|err, _req| {
                // create custom error response
                // oab::Error::InternalServerError
                warn!("{:#?}", err);
                actix_web::error::InternalError::from_response(
                    err,
                    actix_web::HttpResponse::Conflict().finish(),
                )
                .into()
            });
        let app = App::new();
        app.wrap(logger)
            .wrap(middleware::Compress::default())
            .service(fs::Files::new("/media", CONFIG.media_path.clone()).show_files_listing())
            .service(
                web::scope("api")
                    .app_data(web::Data::new(data.clone()))
                    .wrap(
                        ErrorHandlers::new()
                            .handler(StatusCode::INTERNAL_SERVER_ERROR, add_error_header),
                    )
                    .wrap(libs::auth::Auth)
                    .app_data(json_config)
                    .configure(api::routes),
            )
    });
    info!("listen to {}", CONFIG.server_url);
    serv.bind(CONFIG.server_url.clone())?.run().await?;
    Ok(())
}

fn add_error_header<B>(
    res: dev::ServiceResponse<B>,
) -> std::result::Result<ErrorHandlerResponse<B>, actix_web::Error> {
    error!("{}", res.response().error().unwrap());

    Ok(ErrorHandlerResponse::Response(res.map_into_left_body()))
}
