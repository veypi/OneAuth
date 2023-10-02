//
// appuser.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-09-30 23:11
// Distributed under terms of the MIT license.
//

use actix_web::{get, web, Responder};
use proc::access_read;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter};

use crate::{models::app_user, AppState, Error, Result};

#[get("/app/{aid}/user/{uid}")]
#[access_read("app")]
pub async fn get(
    params: web::Path<(String, String)>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    let (mut aid, mut uid) = params.into_inner();
    let mut q = app_user::Entity::find();
    if uid == "-" {
        uid = "".to_string();
    } else {
        q = q.filter(app_user::Column::UserId.eq(uid.clone()));
    }
    if aid == "-" {
        aid = "".to_string();
    } else {
        q = q.filter(app_user::Column::AppId.eq(aid.clone()));
    }
    if uid.is_empty() && aid.is_empty() {
        Err(Error::Missing("uid or aid".to_string()))
    } else {
        let s: Vec<app_user::Model> = q.all(stat.db()).await?;
        Ok(web::Json(s))
    }
}
