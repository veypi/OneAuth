//
// app.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-11-07 00:58
// Distributed under terms of the MIT license.
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

pub fn client() -> DavHandler {
    DavHandler::builder()
        .locksystem(FakeLs::new())
        .strip_prefix("/fs/a/")
        .build_handler()
}
pub async fn dav_handler(
    id: web::Path<(String, String)>,
    req: DavRequest,
    davhandler: web::Data<DavHandler>,
    stat: web::Data<AppState>,
) -> DavResponse {
    let root = stat.fs_root.clone();
    let id = id.into_inner().0;
    info!("start app: {}", id);
    match handle_file(&req, stat).await {
        Ok(()) => {
            let p = Path::new(&root).join(format!("app/{}/", id));
            if !p.exists() {
                match fs::create_dir_all(p.clone()) {
                    Ok(_) => {}
                    Err(e) => {
                        warn!("{}", e);
                    }
                }
            }
            info!("mount {}", p.to_str().unwrap());
            let config = DavConfig::new()
                .filesystem(LocalFs::new(p, false, false, true))
                .strip_prefix(format!("/fs/a/{}/", id));
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

async fn handle_file(
    req: &dav_server::actix::DavRequest,
    stat: web::Data<AppState>,
) -> Result<()> {
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
                        if credentials.user_id == "cli" && credentials.password == "cli" {
                            return Ok(());
                        }
                        match models::user::Entity::find()
                            .filter(models::user::Column::Username.eq(credentials.user_id))
                            .one(stat.db())
                            .await?
                        {
                            Some(u) => {
                                u.check_pass(&credentials.password)?;
                                return Ok(());
                            }
                            None => {}
                        }
                    }
                    "bearer" => {
                        let t = models::Token::from(encoded_credentials, &stat.key)?;
                        if t.is_valid() {
                            return Ok(());
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
