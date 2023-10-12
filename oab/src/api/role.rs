//
// role.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//
//
use actix_web::{delete, get, patch, post, web, Responder};
use proc::{access_create, access_delete, access_read, access_update, crud_update};
use sea_orm::{ActiveModelTrait, ColumnTrait, ConnectionTrait, EntityTrait, QueryFilter};
use serde::{Deserialize, Serialize};
use tracing::info;

use crate::{
    models::{self},
    AppState, Error, Result,
};

#[get("/app/{aid}/role/")]
#[access_read("app")]
pub async fn list(aid: web::Path<String>, stat: web::Data<AppState>) -> Result<impl Responder> {
    let n = aid.into_inner();
    if !n.is_empty() {
        let s = models::role::Entity::find()
            .filter(models::role::Column::AppId.eq(n))
            .all(stat.db())
            .await?;
        Ok(web::Json(s))
    } else {
        Err(Error::Missing("id".to_string()))
    }
}

#[derive(Debug, Deserialize, Serialize)]
pub struct CreateOpt {
    name: String,
    des: Option<String>,
}
#[post("/app/{aid}/role/")]
#[access_create("app")]
pub async fn create(
    aid: web::Path<String>,
    stat: web::Data<AppState>,
    data: web::Json<CreateOpt>,
) -> Result<impl Responder> {
    let data = data.into_inner();
    let id = uuid::Uuid::new_v4().to_string().replace("-", "");
    let obj = models::role::ActiveModel {
        name: sea_orm::ActiveValue::Set(data.name),
        id: sea_orm::ActiveValue::Set(id),
        des: sea_orm::ActiveValue::Set(data.des),
        app_id: sea_orm::ActiveValue::Set(aid.into_inner()),
        ..Default::default()
    };
    let obj = obj.insert(stat.db()).await?;
    Ok(web::Json(obj))
}

#[delete("/app/{aid}/role/{rid}")]
#[access_delete("app", id = "&id.clone().0")]
pub async fn delete(
    id: web::Path<(String, String)>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    let (aid, rid) = id.into_inner();
    let res = models::role::Entity::delete_many()
        .filter(models::role::Column::AppId.eq(aid))
        .filter(models::role::Column::Id.eq(rid))
        .exec(stat.db())
        .await?;
    info!("{:#?}", res);
    Ok("ok")
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct UpdateOpt {
    pub des: Option<String>,
}

#[patch("/app/{aid}/role/{rid}")]
#[access_update("app", id = "&id.clone()[0]")]
#[crud_update(role, filter = "AppId, Id", props = "des")]
pub async fn update(
    id: web::Path<[String; 2]>,
    data: web::Json<UpdateOpt>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    Ok("")
}

#[get("/app/{aid}/role/{rid}/user/{uid}")]
#[access_delete("app", id = "&id.clone().0")]
pub async fn add(
    id: web::Path<(String, String, String)>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    let (_, rid, uid) = id.into_inner();
    let s = models::user_role::ActiveModel {
        user_id: sea_orm::ActiveValue::Set(uid),
        role_id: sea_orm::ActiveValue::Set(rid.clone()),
        ..Default::default()
    };
    let s = s.insert(stat.db()).await?;
    let sql = format!(
        "update role set user_count = user_count + 1 where id = '{}'",
        rid,
    );
    stat.db()
        .execute(sea_orm::Statement::from_string(
            sea_orm::DatabaseBackend::MySql,
            sql,
        ))
        .await?;
    Ok(web::Json(s))
}

#[delete("/app/{aid}/role/{rid}/user/{uid}")]
#[access_delete("app", id = "&id.clone().0")]
pub async fn drop(
    id: web::Path<(String, String, String)>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    let (_, rid, uid) = id.into_inner();
    models::user_role::Entity::delete_many()
        .filter(models::user_role::Column::RoleId.eq(rid.clone()))
        .filter(models::user_role::Column::UserId.eq(uid))
        .exec(stat.db())
        .await?;
    let sql = format!(
        "update role set user_count = user_count - 1 where id = '{}'",
        rid,
    );
    stat.db()
        .execute(sea_orm::Statement::from_string(
            sea_orm::DatabaseBackend::MySql,
            sql,
        ))
        .await?;
    Ok("ok")
}
