//
// user.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//

use std::fmt::Debug;

use crate::{
    libs,
    models::{self, app_user, user, AUStatus, UserPlugin},
    AppState, Error, Result,
};
use actix_web::{delete, get, head, http, patch, post, web, HttpResponse, Responder};
use base64;
use chrono::Local;
use proc::{access_read, access_update, crud_update};
use rand::Rng;
use sea_orm::{ActiveModelTrait, ColumnTrait, EntityTrait, QueryFilter, TransactionTrait};
use serde::{Deserialize, Serialize};
use tracing::info;

#[get("/user/{id}")]
#[access_read("user", id = "&id.clone()")]
pub async fn get(id: web::Path<String>, stat: web::Data<AppState>) -> Result<impl Responder> {
    let n = id.into_inner();
    let db = stat.db();
    if !n.is_empty() {
        let d: Option<models::entity::user::Model> =
            models::entity::user::Entity::find_by_id(n).one(db).await?;
        Ok(web::Json(d))
    } else {
        Err(Error::Missing("id".to_string()))
    }
}

#[derive(Debug, Deserialize, Serialize)]
pub struct ListOptions {
    name: Option<String>,
    role_id: Option<String>,
    app_id: Option<String>,
}

#[get("/user/")]
#[access_read("user")]
pub async fn list(
    stat: web::Data<AppState>,
    query: web::Query<ListOptions>,
) -> Result<impl Responder> {
    let query = query.into_inner();
    let res = if let Some(v) = query.name {
        user::Entity::find()
            .filter(user::Column::Username.contains(v))
            .all(stat.db())
            .await?
    } else if let Some(v) = query.role_id {
        models::user_role::Entity::find()
            .filter(models::user_role::Column::RoleId.eq(v))
            .find_also_related(user::Entity)
            .all(stat.db())
            .await?
            .into_iter()
            .filter_map(|(_, u)| return u)
            .collect()
    } else if let Some(v) = query.app_id {
        models::app_user::Entity::find()
            .filter(models::app_user::Column::AppId.eq(v))
            .find_also_related(user::Entity)
            .all(stat.db())
            .await?
            .into_iter()
            .filter_map(|(_, u)| return u)
            .collect()
    } else {
        user::Entity::find().all(stat.db()).await?
    };
    Ok(web::Json(res))
}

#[derive(Debug, Deserialize, Serialize)]
pub struct LoginOpt {
    typ: Option<String>,
    password: String,
}

#[head("/user/{id}")]
pub async fn login(
    q: web::Query<LoginOpt>,
    id: web::Path<String>,
    stat: web::Data<AppState>,
) -> Result<HttpResponse> {
    let id = id.into_inner();
    let q = q.into_inner();
    let filter = match q.typ {
        Some(t) => match t.as_str() {
            "phone" => user::Column::Phone.eq(id),
            "email" => user::Column::Email.eq(id),
            _ => user::Column::Username.eq(id),
        },
        _ => user::Column::Username.eq(id),
    };
    let p = match base64::decode(q.password.as_bytes()) {
        Err(_) => return Err(Error::ArgInvalid("password".to_string())),
        Ok(p) => p,
    };
    let p = match std::str::from_utf8(&p) {
        Ok(p) => p,
        Err(_) => return Err(Error::ArgInvalid("password".to_string())),
    };
    let u: Option<user::Model> = models::user::Entity::find()
        .filter(filter)
        .one(stat.db())
        .await?;
    let u = match u {
        Some(u) => u,
        None => return Err(Error::NotFound("user".to_string())),
    };

    u.check_pass(p)?;
    let au: Option<app_user::Model> = app_user::Entity::find()
        .filter(app_user::Column::AppId.eq(&stat.uuid))
        .filter(app_user::Column::UserId.eq(&u.id))
        .one(stat.db())
        .await?;
    let au = match au {
        Some(au) => au,
        None => {
            // 未绑定应用时进行绑定操作
            let aid = stat.uuid.clone();
            let db = stat.db().begin().await?;
            let s = libs::user::connect_to_app(u.id.clone(), aid, &db, None).await?;
            db.commit().await?;
            s
        }
    };
    if au.status == AUStatus::OK as i32 {
        let result = sqlx::query_as::<_, models::AccessCore>(
            "select access.name, access.rid, access.level from access, user_role, role WHERE user_role.user_id = ? && access.role_id=user_role.role_id && role.id=user_role.role_id && role.app_id = ?",
        )
        .bind(&u.id)
        .bind(stat.uuid.clone())
        .fetch_all(stat.sqlx())
        .await?;
        Ok(HttpResponse::build(http::StatusCode::OK)
            .insert_header(("auth_token", u.token(result).to_string(&stat.key)?))
            .body("".to_string()))
    } else {
        Ok(HttpResponse::build(http::StatusCode::FORBIDDEN)
            .insert_header(("error", au.status))
            .body("".to_string()))
    }
}

#[derive(Debug, Deserialize, Serialize)]
pub struct RegisterOpt {
    username: String,
    password: String,
}

#[post("/user/")]
pub async fn register(
    q: web::Json<RegisterOpt>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    let q = q.into_inner();
    // let mut tx = dbtx().await;
    info!("{:#?}", q);
    let u: Option<models::user::Model> = user::Entity::find()
        .filter(user::Column::Username.eq(&q.username))
        .one(stat.db())
        .await?;
    // 初始化用户信息
    let u: models::user::ActiveModel = match u {
        Some(_) => return Err(Error::ArgDuplicated(format!("username: {}", q.username))),
        None => {
            let p = match base64::decode(q.password.as_bytes()) {
                Err(_) => return Err(Error::ArgInvalid("password".to_string())),
                Ok(p) => p,
            };
            let p = match std::str::from_utf8(&p) {
                Ok(p) => p,
                Err(_) => return Err(Error::ArgInvalid("password".to_string())),
            };
            let mut u = models::user::Model::default();
            u.username = q.username.clone();
            u.id = uuid::Uuid::new_v4().to_string().replace("-", "");
            u.update_pass(&p)?;
            let mut rng = rand::thread_rng();
            let idx: i64 = rng.gen_range(1..221);
            u.icon = Some(format!("/media/icon/usr/{:04}.jpg", idx));
            u.space = 300;
            u.used = 0;
            info!("{}", u.created.to_string());
            u.created = Local::now().naive_utc();
            u.updated = Local::now().naive_utc();
            info!("{}", u.created.to_string());
            u.into()
        }
    };
    let db = stat.db().begin().await?;

    // 创建用户
    let u = u.insert(&db).await?;

    // 关联应用
    libs::user::connect_to_app(u.id.clone(), stat.uuid.clone(), &db, None).await?;

    db.commit().await?;

    Ok(web::Json(u))
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct UpdateOpt {
    pub username: Option<String>,
    pub icon: Option<String>,
    pub nickname: Option<String>,
    pub email: Option<String>,
    pub phone: Option<String>,
    pub test: serde_json::Value,
}

#[patch("/user/{id}")]
#[access_update("user")]
#[crud_update(user, filter = "Id", props = "username, icon, nickname, email, phone")]
pub async fn update(
    id: web::Path<String>,
    stat: web::Data<AppState>,
    data: web::Json<UpdateOpt>,
) -> Result<impl Responder> {
    info!("{:#?}", data.test);
    Ok("")
}

#[delete("/user/")]
pub async fn delete() -> impl Responder {
    ""
}
