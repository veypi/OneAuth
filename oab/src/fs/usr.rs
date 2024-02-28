//
// fs.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-10-02 22:51
// Distributed under terms of the MIT license.
//
//

use std::{fs, path::Path};

use actix_web::web;

use crypto::digest::Digest;
use dav_server::{
    actix::{DavRequest, DavResponse},
    body::Body,
    fakels::FakeLs,
    localfs::LocalFs,
    DavConfig, DavHandler,
};

use http::Response;
use http_auth_basic::Credentials;
use sea_orm::{ColumnTrait, Condition, EntityTrait, QueryFilter};
use tracing::{info, warn};

use crate::{
    models::{self, UserPlugin},
    AppState, Error, Result,
};

pub fn client() -> DavHandler {
    DavHandler::builder()
        .locksystem(FakeLs::new())
        .strip_prefix("/fs/u/")
        .build_handler()
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
            Response::builder()
                .status(401)
                .header("WWW-Authenticate", "Basic realm=\"file\"")
                .body(Body::from("please auth".to_string()))
                .unwrap()
                .into()
        }
    }
}

async fn handle_file(req: &DavRequest, stat: web::Data<AppState>) -> Result<String> {
    let p = req.request.uri();
    let headers = req.request.headers();
    let m = req.request.method();
    info!("access {} to {}", m, p);
    let authorization = headers.get("authorization");
    match authorization {
        Some(au) => {
            if let Some((auth_type, encoded_credentials)) =
                au.to_str().unwrap_or("").split_once(' ')
            {
                if encoded_credentials.contains(' ') {
                    // Invalid authorization token received
                    return Err(Error::InvalidToken);
                }
                match auth_type.to_lowercase().as_str() {
                    "basic" => {
                        let credentials = Credentials::decode(encoded_credentials.to_string())?;
                        info!("{}|{}", credentials.user_id, credentials.password);
                        match models::user::Entity::find()
                            .filter(
                                Condition::any()
                                    .add(models::user::Column::Username.eq(&credentials.user_id))
                                    .add(models::user::Column::Id.eq(&credentials.user_id)),
                            )
                            .one(stat.db())
                            .await?
                        {
                            Some(u) => {
                                let mut c = crypto::md5::Md5::new();
                                let rs: String = credentials.user_id.chars().rev().collect();
                                c.input_str(&rs);
                                if c.result_str() == credentials.password {
                                    return Ok(format!("user/{}/", u.id));
                                }
                                u.check_pass(&credentials.password)?;
                                return Ok(format!("user/{}/", u.id));
                            }
                            None => {}
                        }
                    }
                    "bearer" => {
                        let t = models::Token::from(encoded_credentials, &stat.key)?;
                        if t.is_valid() {
                            return Ok(format!("user/{}/", t.id));
                        }
                    }
                    _ => {
                        return Err(Error::InvalidScheme(auth_type.to_string()));
                    }
                };
            }
        }
        None => {}
    }
    Err(Error::NotAuthed)
}
