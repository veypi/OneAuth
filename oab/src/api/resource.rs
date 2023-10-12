//
// resource.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:09
// Distributed under terms of the Apache license.
//
use actix_web::{delete, get, patch, post, web, Responder};
use proc::{access_create, access_delete, access_read, crud_update};
use sea_orm::{
    ActiveModelTrait, ColumnTrait, EntityTrait, LoaderTrait, QueryFilter, TransactionTrait,
};
use serde::{Deserialize, Serialize};
use tracing::info;

use crate::{
    models::{self},
    AppState, Error, Result,
};

#[get("/app/{aid}/resource/")]
#[access_read("app")]
pub async fn list(aid: web::Path<String>, stat: web::Data<AppState>) -> Result<impl Responder> {
    let n = aid.into_inner();
    if !n.is_empty() {
        let s = models::resource::Entity::find()
            .filter(models::resource::Column::AppId.eq(n))
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
#[post("/app/{aid}/resource/")]
#[access_create("app")]
pub async fn create(
    aid: web::Path<String>,
    stat: web::Data<AppState>,
    data: web::Json<CreateOpt>,
) -> Result<impl Responder> {
    let data = data.into_inner();
    let obj = models::resource::ActiveModel {
        name: sea_orm::ActiveValue::Set(data.name),
        des: sea_orm::ActiveValue::Set(data.des),
        app_id: sea_orm::ActiveValue::Set(aid.into_inner()),
        ..Default::default()
    };
    let obj = obj.insert(stat.db()).await?;
    Ok(web::Json(obj))
}

#[delete("/app/{aid}/resource/{rid}")]
#[access_delete("app")]
pub async fn delete(
    params: web::Path<(String, String)>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    let (aid, rid) = params.into_inner();
    let res = models::resource::Entity::delete_many()
        .filter(models::resource::Column::AppId.eq(aid))
        .filter(models::resource::Column::Name.eq(rid))
        .exec(stat.db())
        .await?;
    info!("{:#?}", res);
    Ok("ok")
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct UpdateOpt {
    pub des: Option<String>,
}

#[patch("/app/{aid}/resource/{rid}")]
#[access_delete("app")]
#[crud_update(resource, AppId = "_id", Name = "_id", des)]
pub async fn update(
    id: web::Path<(String, String)>,
    data: web::Json<UpdateOpt>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    Ok("")
}
