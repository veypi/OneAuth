//
// fs.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-10-02 22:51
// Distributed under terms of the MIT license.
//
//

use std::{fs, path::Path};

use actix_web::web;

use dav_server::{
    actix::{DavRequest, DavResponse},
    body::Body,
    fakels::FakeLs,
    localfs::LocalFs,
    DavConfig, DavHandler,
};

use http::{header, Method, Response};
use http_auth_basic::Credentials;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter};
use tracing::{info, warn};

use crate::{
    models::{self, UserPlugin},
    AppState, Error, Result,
};

pub fn core() -> DavHandler {
    DavHandler::builder()
        .locksystem(FakeLs::new())
        .strip_prefix("/file/")
        .build_handler()
}
/// Try to parse header value as HTTP method.
fn header_value_try_into_method(hdr: &header::HeaderValue) -> Option<Method> {
    hdr.to_str()
        .ok()
        .and_then(|meth| Method::try_from(meth).ok())
}

fn is_request_preflight(req: &DavRequest) -> bool {
    // check request method is OPTIONS
    if req.request.method() != Method::OPTIONS {
        return false;
    }

    // check follow-up request method is present and valid
    if req
        .request
        .headers()
        .get(header::ACCESS_CONTROL_REQUEST_METHOD)
        .and_then(header_value_try_into_method)
        .is_none()
    {
        return false;
    }

    true
}

pub async fn dav_handler(
    req: DavRequest,
    davhandler: web::Data<DavHandler>,
    stat: web::Data<AppState>,
) -> DavResponse {
    let root = stat.fs_root.clone();
    match handle_file(&req, stat).await {
        Ok(p) => {
            let p = Path::new(&root).join(p);
            if !p.exists() {
                match fs::create_dir_all(p.clone()) {
                    Ok(_) => {}
                    Err(e) => {
                        warn!("{}", e);
                    }
                }
            }
            info!("mount {}", p.to_str().unwrap());
            let config = DavConfig::new().filesystem(LocalFs::new(p, false, false, true));
            davhandler.handle_with(config, req.request).await.into()
        }
        Err(e) => {
            warn!("handle file failed: {}", e);
            if is_request_preflight(&req) {
                let origin = match req.request.headers().get("Origin") {
                    Some(o) => o.to_str().unwrap(),
                    None => "",
                };
                let allowed_headers =
                    match req.request.headers().get("Access-Control-Request-Headers") {
                        Some(o) => o.to_str().unwrap(),
                        None => "",
                    };
                let allowed_method =
                    match req.request.headers().get("Access-Control-Request-Method") {
                        Some(o) => o.to_str().unwrap(),
                        None => "",
                    };
                Response::builder()
                    .status(200)
                    .header("WWW-Authenticate", "Basic realm=\"file\"")
                    .header("Access-Control-Allow-Origin", origin)
                    .header("Access-Control-Allow-Credentials", "true")
                    .header("Access-Control-Allow-Headers", allowed_headers)
                    .header("Access-Control-Allow-Methods", allowed_method)
                    .header(
                        "Access-Control-Expose-Headers",
                        "access-control-allow-origin, content-type",
                    )
                    .body(Body::from("please auth".to_string()))
                    .unwrap()
                    .into()
            } else {
                Response::builder()
                    .status(401)
                    .header("WWW-Authenticate", "Basic realm=\"file\"")
                    .body(Body::from("please auth".to_string()))
                    .unwrap()
                    .into()
            }
        }
    }
}

async fn handle_file(req: &DavRequest, stat: web::Data<AppState>) -> Result<String> {
    let p = req.request.uri();
    let headers = req.request.headers();
    let m = req.request.method();
    // handle_authorization(req.request.headers());
    info!("access {} to {}", m, p);
    let auth_token = headers.get("auth_token");
    let authorization = headers.get("authorization");
    let app_id = match headers.get("app_id") {
        Some(i) => i.to_str().unwrap_or(""),
        None => "",
    };
    match auth_token {
        Some(t) => match models::Token::from(t.to_str().unwrap_or(""), &stat.key) {
            Ok(t) => {
                if t.is_valid() {
                    if app_id != "" {
                        // 只有秘钥才能访问app数据
                        if t.can_read("app", app_id) {
                            return Ok(format!("app/{}/", app_id));
                        }
                    } else {
                        return Ok(format!("user/{}/", t.id));
                    }
                }
            }
            Err(_) => {}
        },
        None => {}
    }
    match authorization {
        Some(au) => {
            let credentials =
                Credentials::from_header(au.to_str().unwrap_or("").to_string()).unwrap();
            info!("{}|{}", credentials.user_id, credentials.password);
            match models::user::Entity::find()
                .filter(models::user::Column::Username.eq(credentials.user_id))
                .one(stat.db())
                .await?
            {
                Some(u) => {
                    u.check_pass(&credentials.password)?;
                    return Ok(format!("user/{}/", u.id));
                }
                None => {}
            }
        }
        None => {}
    }
    Err(Error::NotAuthed)
}
