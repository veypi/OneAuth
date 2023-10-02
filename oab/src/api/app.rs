//
// app.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//
//
use actix_web::{delete, get, post, web, Responder};
use proc::access_read;
use serde::{Deserialize, Serialize};

use crate::{
    models::{self, app},
    Error, Result, CONFIG,
};
use chrono::NaiveDateTime;

#[get("/app/{id}")]
#[access_read("app")]
pub async fn get(id: web::Path<String>) -> Result<impl Responder> {
    let n = id.into_inner();
    if !n.is_empty() {
        let s = sqlx::query_as::<_, app::Model>("select * from app where id = ?")
            .bind(n)
            .fetch_one(CONFIG.sqlx())
            .await?;
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
pub async fn list() -> Result<impl Responder> {
    let result = sqlx::query_as::<_, App>(
        "select app.id,app.created, app.updated, app.icon, app.name, app.des, app.user_count, app.hide,app.join_method, app.role_id, app.redirect, app.status, app_user.status as u_status from app left join  app_user on app_user.user_id = ? && app_user.app_id = app.id",
        ).bind(_auth_token.id)
        .fetch_all(CONFIG.sqlx())
        .await?;

    Ok(web::Json(result))
}

#[post("/app/")]
pub async fn create() -> Result<impl Responder> {
    Ok("")
}
#[delete("/app/{id}")]
pub async fn del(id: web::Path<String>) -> Result<impl Responder> {
    Ok("")
}
