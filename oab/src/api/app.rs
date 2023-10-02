//
// app.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//
//
use actix_web::{delete, get, post, web, Responder};
use proc::{access_create, access_read};
use sea_orm::{ActiveModelTrait, EntityTrait, TransactionTrait};
use serde::{Deserialize, Serialize};
use tracing::info;

use crate::{
    models::{self, access, app, app_user, rand_str, AUStatus, AccessLevel, Token},
    AppState, Error, Result,
};
use chrono::NaiveDateTime;

#[get("/app/{id}")]
#[access_read("app")]
pub async fn get(id: web::Path<String>, stat: web::Data<AppState>) -> Result<impl Responder> {
    let n = id.into_inner();
    if !n.is_empty() {
        let s = app::Entity::find_by_id(n).one(stat.db()).await?;
        Ok(web::Json(s))
    } else {
        Err(Error::Missing("id".to_string()))
    }
}

#[derive(Debug, Serialize, Deserialize, sqlx::FromRow)]
pub struct App {
    pub id: String,
    pub created: Option<NaiveDateTime>,
    pub updated: Option<NaiveDateTime>,

    pub name: Option<String>,
    pub des: Option<String>,
    pub icon: Option<String>,
    pub user_count: i64,

    pub hide: bool,
    pub join_method: models::AppJoin,
    pub role_id: Option<String>,
    pub redirect: Option<String>,

    pub status: i64,
    pub u_status: i64,
}

#[get("/app/")]
#[access_read("app")]
pub async fn list(stat: web::Data<AppState>) -> Result<impl Responder> {
    let res = app::Entity::find().all(stat.db()).await?;
    // let result = sqlx::query_as::<_,app::Model>(
    //     "select app.*,app_userstatus as status from app left join  app_user on app_user.user_id = ? && app_user.app_id = app.id",
    //     ).bind(_auth_token.id)
    //     .fetch_all(stat.sqlx())
    //     .await?;

    Ok(web::Json(res))
}

#[derive(Debug, Deserialize, Serialize)]
pub struct CreateOpt {
    name: String,
    icon: Option<String>,
    // enable_register: Option<String>,
    // des: Option<String>,
    // host: Option<String>,
    // redirect: Option<String>,
}

#[post("/app/")]
#[access_create("app")]
pub async fn create(
    stat: web::Data<AppState>,
    data: web::Json<CreateOpt>,
    t: web::ReqData<Token>,
) -> Result<impl Responder> {
    let data = data.into_inner();
    let id = uuid::Uuid::new_v4().to_string().replace("-", "");
    let t = t.into_inner();
    info!("{} create app {}", t.id, t.nickname);
    let obj = app::ActiveModel {
        name: sea_orm::ActiveValue::Set(data.name),
        icon: sea_orm::ActiveValue::Set(data.icon),
        id: sea_orm::ActiveValue::Set(id),
        key: sea_orm::ActiveValue::Set(rand_str(32)),
        ..Default::default()
    };
    let db = stat.db().begin().await?;
    let obj: app::Model = obj.insert(&db).await?;
    let ac = access::ActiveModel {
        app_id: sea_orm::ActiveValue::Set(stat.uuid.clone()),
        name: sea_orm::ActiveValue::Set("app".to_string()),
        rid: sea_orm::ActiveValue::Set(Some(obj.id.clone())),
        user_id: sea_orm::ActiveValue::Set(Some(t.id.clone())),
        level: sea_orm::ActiveValue::Set(AccessLevel::ALL as i32),
        ..Default::default()
    };
    let ac: access::Model = ac.insert(&db).await?;
    let au = app_user::ActiveModel {
        app_id: sea_orm::ActiveValue::Set(obj.id.clone()),
        user_id: sea_orm::ActiveValue::Set(t.id.clone()),
        status: sea_orm::ActiveValue::Set(AUStatus::OK as i32),
        ..Default::default()
    };
    let au: app_user::Model = au.insert(&db).await?;
    db.commit().await?;
    Ok(web::Json(obj))
}
#[delete("/app/{id}")]
pub async fn del(_id: web::Path<String>) -> Result<impl Responder> {
    Ok("")
}
