//
// main.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-07 23:51
// Distributed under terms of the Apache license.
//

use actix_files;
use actix_web::{
    dev::{self},
    http::StatusCode,
    middleware::{self, ErrorHandlerResponse, ErrorHandlers},
    web::{self},
    App, HttpResponse, HttpServer, Responder,
};
use mime_guess::from_path;
use rust_embed::RustEmbed;

use oab::{api, init_log, libs, models, AppCli, AppState, Clis, Result};
use tracing::{error, info, warn};

#[tokio::main]
async fn main() -> Result<()> {
    let cli = AppCli::new();
    let mut data = AppState::new(&cli);
    if data.debug {
        std::env::set_var("RUST_LOG", "debug");
        std::env::set_var("RUST_BACKTRACE", "full");
    }
    let _log = init_log(&data);
    if cli.handle_service(data.clone())? {
        info!("2");
        return Ok(());
    }
    if let Some(c) = &cli.command {
        match c {
            Clis::Init => {
                data.connect_sqlx()?;
                models::init(data).await;
                return Ok(());
            }
            _ => {}
        };
    };
    data.connect().await?;
    data.connect_sqlx()?;
    web(data).await?;
    Ok(())
}

async fn web(data: AppState) -> Result<()> {
    // libs::task::start_nats_online(client.clone());
    libs::task::start_stats_info(data.ts_url.clone());
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
        let cors = actix_cors::Cors::default()
            .allow_any_method()
            .allow_any_header()
            .supports_credentials()
            .allowed_origin_fn(|_, _| {
                return true;
            });
        let app = App::new();
        app.wrap(logger)
            .wrap(middleware::Compress::default())
            .app_data(web::Data::new(data.clone()))
            .service(
                actix_files::Files::new("/media", data.media_path.clone()).show_files_listing(),
            )
            .service(
                web::scope("api")
                    .wrap(cors)
                    .wrap(
                        ErrorHandlers::new()
                            .handler(StatusCode::INTERNAL_SERVER_ERROR, add_error_header),
                    )
                    .wrap(libs::auth::Auth {
                        key: data.key.clone(),
                    })
                    .app_data(json_config)
                    .configure(api::routes),
            )
            .service(
                web::scope("fs")
                    .wrap(oab::fs::FsWrap {})
                    .configure(oab::fs::routes),
            )
            .service(index)
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

#[derive(RustEmbed)]
#[folder = "dist/"]
struct Asset;

#[actix_web::get("/{_:.*}")]
async fn index(p: web::Path<String>) -> impl Responder {
    info!("{}", p);

    let p = &p.into_inner();
    match Asset::get(p) {
        Some(content) => HttpResponse::Ok()
            .content_type(from_path(p).first_or_octet_stream().as_ref())
            .body(content.data.into_owned()),
        None => HttpResponse::Ok()
            .content_type(from_path("index.html").first_or_octet_stream().as_ref())
            .body(Asset::get("index.html").unwrap().data.into_owned()),
    }
}
