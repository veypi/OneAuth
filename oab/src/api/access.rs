//
// access.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//
//
//
use actix_web::{delete, get, patch, post, web, Responder};
use proc::{access_create, access_delete, access_read, crud_update};
use sea_orm::{ActiveModelTrait, ColumnTrait, EntityTrait, QueryFilter};
use serde::{Deserialize, Serialize};
use tracing::info;

use crate::{
    models::{self},
    AppState, Error, Result,
};

#[derive(Debug, Deserialize, Serialize)]
pub struct Options {
    name: Option<String>,
    role_id: Option<String>,
    user_id: Option<String>,
    rid: Option<String>,
    level: Option<bool>,
}

#[get("/app/{aid}/access/")]
#[access_read("app")]
pub async fn list(
    aid: web::Path<String>,
    stat: web::Data<AppState>,
    query: web::Query<Options>,
) -> Result<impl Responder> {
    let aid = aid.into_inner();
    let mut q = models::access::Entity::find().filter(models::access::Column::AppId.eq(aid));
    let query = query.into_inner();
    if let Some(rid) = query.role_id {
        q = q.filter(models::access::Column::RoleId.eq(rid));
    };
    if let Some(v) = query.name {
        q = q.filter(models::access::Column::Name.eq(v));
    };
    if let Some(v) = query.user_id {
        q = q.filter(models::access::Column::UserId.eq(v));
    };
    let aus = q.all(stat.db()).await?;
    Ok(web::Json(aus))
}

#[derive(Debug, Deserialize, Serialize)]
pub struct CreateOptions {
    name: String,
    level: i32,
    role_id: Option<String>,
    user_id: Option<String>,
    rid: Option<String>,
}
#[post("/app/{aid}/access/")]
#[access_create("app")]
pub async fn creat(
    aid: web::Path<String>,
    stat: web::Data<AppState>,
    query: web::Json<CreateOptions>,
) -> Result<impl Responder> {
    let aid = aid.into_inner();
    let query = query.into_inner();
    let obj = models::access::ActiveModel {
        app_id: sea_orm::ActiveValue::Set(aid),
        name: sea_orm::ActiveValue::Set(query.name),
        level: sea_orm::ActiveValue::Set(query.level),
        role_id: sea_orm::ActiveValue::Set(query.role_id),
        user_id: sea_orm::ActiveValue::Set(query.user_id),
        rid: sea_orm::ActiveValue::Set(query.rid),
        ..Default::default()
    };
    let obj: models::access::Model = obj.insert(stat.db()).await?;
    Ok(web::Json(obj))
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct UpdateOpt {
    pub level: Option<i32>,
    pub rid: Option<String>,
}

#[patch("/app/{aid}/access/{id}")]
#[access_delete("app")]
#[crud_update(access, filter = "AppId, Id", props = "level, rid")]
pub async fn update(
    id: web::Path<[String; 2]>,
    data: web::Json<UpdateOpt>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    Ok("")
}

#[delete("/app/{aid}/access/{id}")]
#[access_delete("app")]
pub async fn delete(
    params: web::Path<(String, String)>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    let (aid, rid) = params.into_inner();
    let res = models::access::Entity::delete_many()
        .filter(models::access::Column::AppId.eq(aid))
        .filter(models::access::Column::Id.eq(rid))
        .exec(stat.db())
        .await?;
    info!("{:#?}", res);
    Ok("ok")
}

// mod test {
//     use crate::{
//         models::{self},
//         AppState, Error, Result,
//     };
//     use actix_web::{delete, get, patch, post, web, Responder};
//     use proc::crud_test;
//     use proc::{access_create, access_delete, access_read, crud_update};
//     use sea_orm::{ActiveModelTrait, ColumnTrait, EntityTrait, QueryFilter};
//     use serde::{Deserialize, Serialize};
//     use tracing::info;
//     #[derive(Debug, Clone, Deserialize, Serialize)]
//     pub struct UpdateOpt {
//         pub level: Option<i32>,
//         pub rid: Option<String>,
//     }
//     #[derive(Debug, Clone, Deserialize, Serialize)]
//     pub struct IDOpt {
//         pub app_id: Option<String>,
//         pub id: Option<String>,
//     }
//     #[crud_test(access, filter = "AppId, Id", props = "level, rid")]
//     pub async fn update(
//         id: web::Path<[String; 2]>,
//         data: web::Json<UpdateOpt>,
//         stat: web::Data<AppState>,
//     ) -> Result<impl Responder> {
//         Ok("")
//     }
// }
