//
// app.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//
//
use actix_web::{delete, get, patch, post, web, Responder};
use proc::{access_create, access_delete, access_read, access_update, crud_update};
use sea_orm::{ActiveModelTrait, ColumnTrait, EntityTrait, QueryFilter, TransactionTrait};
use serde::{Deserialize, Serialize};
use tracing::info;

use crate::{
    libs,
    models::{access, app, rand_str, AccessLevel, Token},
    AppState, Error, Result,
};

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

#[get("/app/")]
#[access_read("app")]
pub async fn list(stat: web::Data<AppState>) -> Result<impl Responder> {
    let res = app::Entity::find().all(stat.db()).await?;
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
    ac.insert(&db).await?;
    libs::user::connect_to_app(t.id.clone(), obj.id.clone(), &db, Some(obj.clone())).await?;
    db.commit().await?;
    Ok(web::Json(obj))
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct UpdateOpt {
    pub name: Option<String>,
    pub icon: Option<String>,
    pub des: Option<String>,
    pub join_method: Option<i32>,
    pub role_id: Option<String>,
    pub redirect: Option<String>,
    pub host: Option<String>,
    pub status: Option<i32>,
}

#[patch("/app/{id}")]
#[access_update("app")]
#[crud_update(
    app,
    filter = "Id",
    props = "name,icon,des,join_method,role_id,redirect,host,status"
)]
pub async fn update(
    id: web::Path<String>,
    stat: web::Data<AppState>,
    data: web::Json<UpdateOpt>,
) -> Result<impl Responder> {
    Ok("")
}

#[delete("/app/{id}")]
#[access_delete("app")]
pub async fn del(_id: web::Path<String>) -> Result<impl Responder> {
    Ok("")
}
