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

use oab::{api, init_log, libs, models, AppState, Clis, Result, CLI};
use tracing::{error, info, warn};

#[tokio::main]
async fn main() -> Result<()> {
    std::env::set_var("RUST_LOG", "debug");
    std::env::set_var("RUST_BACKTRACE", "1");
    init_log();
    let mut data = AppState::new();
    data.connect().await?;
    data.connect_sqlx()?;
    if let Some(c) = &CLI.command {
        match c {
            Clis::Init => {
                models::init(data).await;
                return Ok(());
            }
            _ => {}
        };
    };
    web(data).await?;
    Ok(())
}
async fn web(data: AppState) -> Result<()> {
    let url = data.server_url.clone();
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
            .service(fs::Files::new("/media", data.media_path.clone()).show_files_listing())
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
    info!("listen to {}", url);
    serv.bind(url)?.run().await?;
    Ok(())
}

fn add_error_header<B>(
    res: dev::ServiceResponse<B>,
) -> std::result::Result<ErrorHandlerResponse<B>, actix_web::Error> {
    error!("{}", res.response().error().unwrap());

    Ok(ErrorHandlerResponse::Response(res.map_into_left_body()))
}
