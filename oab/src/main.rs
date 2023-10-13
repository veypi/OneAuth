//
// main.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-07 23:51
// Distributed under terms of the Apache license.
//

use actix_files as fs;
use actix_web::{
    dev::{self, Service},
    http::StatusCode,
    middleware::{self, ErrorHandlerResponse, ErrorHandlers},
    web::{self},
    App, HttpResponse, HttpServer, Responder,
};
use futures_util::future::FutureExt;
use mime_guess::from_path;
use rust_embed::RustEmbed;

use http::{HeaderName, HeaderValue};
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
            Clis::Install | Clis::Uninstall | Clis::Start | Clis::Stop => {
                return CLI.handle_service();
            }
            _ => {}
        };
    };
    web(data).await?;
    Ok(())
}
async fn web(data: AppState) -> Result<()> {
    let url = data.server_url.clone();
    let dav = libs::fs::core();
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
        let cors = actix_cors::Cors::permissive();
        let app = App::new();
        app.wrap(logger)
            .wrap(middleware::Compress::default())
            .app_data(web::Data::new(data.clone()))
            .service(fs::Files::new("/media", data.media_path.clone()).show_files_listing())
            .service(
                web::scope("api")
                    .wrap(cors)
                    .wrap(
                        ErrorHandlers::new()
                            .handler(StatusCode::INTERNAL_SERVER_ERROR, add_error_header),
                    )
                    .wrap(libs::auth::Auth)
                    .app_data(json_config)
                    .configure(api::routes),
            )
            .service(
                web::scope("file")
                    .wrap_fn(|req, srv| {
                        let headers = &req.headers().clone();
                        let origin = match headers.get("Origin") {
                            Some(o) => o.to_str().unwrap().clone().to_string(),
                            None => "".to_string(),
                        };
                        srv.call(req).map(move |res| {
                            let res = match res {
                                Ok(mut expr) => {
                                    let headers = expr.headers_mut();
                                    headers.insert(
                                        HeaderName::try_from("Access-Control-Allow-Origin")
                                            .unwrap(),
                                        HeaderValue::from_str(&origin).unwrap(),
                                    );
                                    Ok(expr)
                                }
                                Err(e) => Err(e),
                            };
                            res
                        })
                    })
                    .app_data(web::Data::new(dav.clone()))
                    .service(web::resource("/{tail:.*}").to(libs::fs::dav_handler)),
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
