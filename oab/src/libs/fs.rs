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

use http::Response;
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

pub async fn dav_handler(
    req: DavRequest,
    davhandler: web::Data<DavHandler>,
    stat: web::Data<AppState>,
) -> DavResponse {
    let root = stat.fs_root.clone();
    match handle_file(req, stat).await {
        Ok((p, req)) => {
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
            Response::builder()
                .status(401)
                .header("WWW-Authenticate", "Basic realm=\"file\"")
                .body(Body::from("please auth".to_string()))
                .unwrap()
                .into()
        }
    }
}

async fn handle_file(req: DavRequest, stat: web::Data<AppState>) -> Result<(String, DavRequest)> {
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
        Some(t) => match models::Token::from(t.to_str().unwrap_or("")) {
            Ok(t) => {
                if t.is_valid() {
                    if app_id != "" {
                        // 只有秘钥才能访问app数据
                        if t.can_read("app", app_id) {
                            return Ok((format!("app/{}/", app_id), req));
                        }
                    } else {
                        return Ok((format!("user/{}/", t.id), req));
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
                    return Ok((format!("user/{}/", u.id), req));
                }
                None => {}
            }
        }
        None => {}
    }
    Err(Error::NotAuthed)
}
