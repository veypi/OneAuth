//
// access.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//
//
//
use actix_web::{web, Responder};
use proc::crud_update;
use sea_orm::{ActiveModelTrait, ColumnTrait, EntityTrait, QueryFilter, TransactionTrait};
use serde::{Deserialize, Serialize};

use crate::{models::app, AppState, Error, Result};

#[derive(Debug, Deserialize, Serialize, Clone)]
pub struct UpdateOpt {
    pub name: Option<String>,
    pub icon: Option<String>,
    pub des: Option<String>,
    pub join_method: Option<i32>,
    pub role_id: Option<String>,
    pub redirect: Option<String>,
    pub status: Option<i32>,
}
impl UpdateOpt {
    // #[crud_update(
    //     app,
    //     id = "Id",
    //     name,
    //     icon,
    //     des,
    //     join_method,
    //     role_id,
    //     redirect,
    //     status
    // )]
    pub async fn update(
        id: web::Path<String>,
        stat: web::Data<AppState>,
        data: web::Json<UpdateOpt>,
    ) -> Result<impl Responder> {
        let data = data.into_inner();
        let id = id.into_inner();
        let obj = app::Entity::find_by_id(&id).one(stat.db()).await?;
        let mut obj: app::ActiveModel = match obj {
            Some(o) => o.into(),
            None => return Err(Error::NotFound(id)),
        };
        if let Some(name) = data.name {
            obj.name = sea_orm::Set(name)
        };
        let obj = obj.update(stat.db()).await?;
        Ok(web::Json(obj))
    }
}

pub async fn update(
    id: web::Path<String>,
    stat: web::Data<AppState>,
    data: web::Json<UpdateOpt>,
) -> Result<impl Responder> {
    let _id = &id.clone();
    let _data = data.clone();
    let _db = &stat.db().clone();
    let f = || async move {
        let data = data.into_inner();
        let id = id.into_inner();
        let obj = app::Entity::find_by_id(&id).one(stat.db()).await?;
        let mut obj: app::ActiveModel = match obj {
            Some(o) => o.into(),
            None => return Err(Error::NotFound(id)),
        };
        if let Some(name) = data.name {
            obj.name = sea_orm::Set(name)
        };
        let obj = obj.update(stat.db()).await?;
        Ok(web::Json(obj))
    };
    let res = f().await;
    match res {
        Err(e) => Err(e),
        Ok(res) => {
            let obj = crate::models::app::Entity::find_by_id(_id).one(_db).await?;
            let mut obj: crate::models::app::ActiveModel = match obj {
                Some(o) => o.into(),
                None => return Err(Error::NotFound(_id.to_owned())),
            };
            if let Some(name) = _data.name {
                obj.name = sea_orm::Set(name.into())
            };
            if let Some(icon) = _data.icon {
                obj.icon = sea_orm::Set(icon.into())
            };
            if let Some(des) = _data.des {
                obj.des = sea_orm::Set(des.into())
            };
            if let Some(join_method) = _data.join_method {
                obj.join_method = sea_orm::Set(join_method.into())
            };
            if let Some(role_id) = _data.role_id {
                obj.role_id = sea_orm::Set(role_id.into())
            };
            if let Some(redirect) = _data.redirect {
                obj.redirect = sea_orm::Set(redirect.into())
            };
            if let Some(status) = _data.status {
                obj.status = sea_orm::Set(status.into())
            };
            let obj = obj.update(_db).await?;
            Ok(actix_web::web::Json(obj))
        }
    }
}
