//
// appuser.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-09-30 23:11
// Distributed under terms of the MIT license.
//

use actix_web::{delete, get, post, web, Responder};
use proc::access_read;
use serde::{Deserialize, Serialize};
use tracing::info;

use crate::{models, Error, Result, CONFIG};

#[get("/app/{aid}/user/{uid}")]
#[access_read("app")]
pub async fn get(params: web::Path<(String, String)>) -> Result<impl Responder> {
    let (mut aid, mut uid) = params.into_inner();
    if uid == "-" {
        uid = "".to_string();
    }
    if aid == "-" {
        aid = "".to_string();
    }
    let sql = format!("select * from app_user where");
    info!("111|{}|{}|", aid, uid);
    if uid.is_empty() && aid.is_empty() {
        Err(Error::Missing("uid or aid".to_string()))
    } else {
        let s = sqlx::query_as::<_, models::AppUser>(
            "select * from app_user where app_id = ? and user_id = ?",
        )
        .bind(aid)
        .bind(uid)
        .fetch_all(CONFIG.db())
        .await?;
        Ok(web::Json(s))
    }
}
