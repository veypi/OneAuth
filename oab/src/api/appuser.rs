//
// appuser.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-09-30 23:11
// Distributed under terms of the MIT license.
//

use actix_web::{get, post, web, Responder};
use proc::access_read;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter, TransactionTrait};

use crate::{
    libs,
    models::{app, app_user},
    AppState, Error, Result,
};

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
        let s: Vec<(app_user::Model, Option<app::Model>)> =
            q.find_also_related(app::Entity).all(stat.db()).await?;
        let res: Vec<app::Model> = s
            .into_iter()
            .filter_map(|(l, a)| match a {
                Some(a) => Some(app::Model {
                    status: l.status,
                    ..a
                }),
                None => None,
            })
            .collect();
        // let s: Vec<app_user::Model> = q.all(stat.db()).await?;
        Ok(web::Json(res))
    }
}

#[post("/app/{aid}/user/{uid}")]
#[access_read("app")]
pub async fn add(
    params: web::Path<(String, String)>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    let (aid, uid) = params.into_inner();
    let db = stat.db().begin().await?;
    let res = libs::user::connect_to_app(uid, aid, &db, None).await?;
    db.commit().await?;
    Ok(web::Json(res))
}
