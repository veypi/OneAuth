//
// appuser.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-09-30 23:11
// Distributed under terms of the MIT license.
//

use actix_web::{get, patch, post, web, Responder};
use proc::{access_delete, access_read, crud_update};
use sea_orm::{
    ActiveModelTrait, ColumnTrait, EntityTrait, LoaderTrait, QueryFilter, TransactionTrait,
};
use serde::{Deserialize, Serialize};

use crate::{
    libs,
    models::{self, app, app_user},
    AppState, Error, Result,
};

#[derive(Debug, Deserialize, Serialize)]
pub struct GetOptions {
    app: Option<bool>,
    user: Option<bool>,
}

#[get("/app/{aid}/user/{uid}")]
#[access_read("app")]
pub async fn get(
    params: web::Path<(String, String)>,
    stat: web::Data<AppState>,
    query: web::Query<GetOptions>,
) -> Result<impl Responder> {
    let (mut aid, mut uid) = params.into_inner();
    let mut q = app_user::Entity::find();
    let query = query.into_inner();
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
        let aus = q.all(stat.db()).await?;
        let mut res: Vec<models::UnionAppUser> = aus
            .iter()
            .map(|f| models::UnionAppUser {
                app: None,
                user: None,
                app_id: f.app_id.clone(),
                user_id: f.user_id.clone(),
                status: f.status,
                created: f.created,
                updated: f.updated,
            })
            .collect();
        if Some(true) == query.app {
            aus.load_one(app::Entity, stat.db())
                .await?
                .into_iter()
                .zip(res.iter_mut())
                .for_each(|(a, b)| {
                    b.app = a;
                });
        }
        if Some(true) == query.user {
            aus.load_one(models::user::Entity, stat.db())
                .await?
                .into_iter()
                .zip(res.iter_mut())
                .for_each(|(a, b)| {
                    b.user = a;
                });
        }
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

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct UpdateOpt {
    pub status: Option<i32>,
}

#[patch("/app/{aid}/user/{uid}")]
#[access_delete("app")]
#[crud_update(app_user, AppId = "_id", UserId = "_id", status)]
pub async fn update(
    id: web::Path<(String, String)>,
    data: web::Json<UpdateOpt>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    Ok("")
}
