//
// user.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//

use std::fmt::Debug;

use crate::{
    models::{self, access, app, app_user, user, UserPlugin},
    AppState, Error, Result,
};
use actix_web::{delete, get, head, http, post, web, HttpResponse, Responder};
use base64;
use proc::access_read;
use rand::Rng;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter};
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

#[get("/user/")]
#[access_read("user")]
pub async fn list(stat: web::Data<AppState>) -> Result<impl Responder> {
    let res: Vec<user::Model> = user::Entity::find().all(stat.db()).await?;
    // let result = sqlx::query!(
    //     "select id,updated,created,username,nickname,email,icon,status, used, space from user",
    // )
    // .map(|row| models::user::Model {
    //     id: row.id,
    //     created: row.created,
    //     updated: row.updated,
    //     username: row.username,
    //     nickname: row.nickname,
    //     email: row.email,
    //     status: row.status,
    //     used: row.used,
    //     space: row.space,
    //     icon: row.icon,
    //     ..Default::default()
    // })
    // .fetch_all(CONFIG.sqlx())
    // .await?;
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
    let db = stat.db();
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
    let u: Option<user::Model> = models::user::Entity::find().filter(filter).one(db).await?;
    let u = match u {
        Some(u) => u,
        None => return Err(Error::NotFound("user".to_string())),
    };

    u.check_pass(p)?;
    let au: Option<app_user::Model> = app_user::Entity::find()
        .filter(app_user::Column::AppId.eq(&stat.uuid))
        .filter(app_user::Column::UserId.eq(&u.id))
        .one(db)
        .await?;
    // let au = sqlx::query_as::<_, models::AppUser>(
    //     "select * from app_user where app_id = ? and user_id = ?",
    // )
    // .bind(&CONFIG.uuid)
    // .bind(&u.id)
    // .fetch_optional(CONFIG.sqlx())
    // .await?;
    let i: i64 = match au {
        Some(au) => match au.status.into() {
            models::AUStatus::OK => 0,
            models::AUStatus::Deny => {
                return Err(Error::BusinessException("apply denied".to_string()))
            }
            models::AUStatus::Disabled => {
                return Err(Error::BusinessException("login disabled".to_string()))
            }
            models::AUStatus::Applying => {
                return Err(Error::BusinessException("applying".to_string()))
            }
        },
        None => {
            let app_obj: app::Model = app::Entity::find_by_id(stat.uuid.clone())
                .one(db)
                .await?
                .unwrap();
            let s = match app_obj.join_method.into() {
                models::AppJoin::Disabled => {
                    return Err(Error::BusinessException(
                        "this app diabled login".to_string(),
                    ))
                }
                models::AppJoin::Auto => models::AUStatus::OK,
                models::AppJoin::Applying => models::AUStatus::Applying,
            };
            sqlx::query(
                r#"
insert into app_user (app_id,user_id,status)
values ( ?, ?, ? )
        "#,
            )
            .bind(&app_obj.id)
            .bind(&u.id)
            .bind(&s)
            .execute(stat.sqlx())
            .await?;
            match s {
                models::AUStatus::OK => 0,
                _ => 1,
            }
        }
    };
    if i == 0 {
        // let result: Vec<models::access::Model> = access::Entity::find().all(db).await?;
        let result = sqlx::query_as::<_, access::Model>(
            "select access.name,access.rid,access.level from access, user_role, role WHERE user_role.user_id = ? && access.role_id=user_role.role_id && role.id=user_role.role_id && role.app_id = ?",
        )
        .bind(&u.id)
        .bind(stat.uuid.clone())
        .fetch_all(stat.sqlx())
        .await?;
        Ok(HttpResponse::build(http::StatusCode::OK)
            .insert_header(("auth_token", u.token(result).to_string()?))
            .body("".to_string()))
    } else {
        Ok(HttpResponse::build(http::StatusCode::OK)
            .insert_header(("stat", "applying"))
            .body("".to_string()))
    }
}

#[derive(Debug, Deserialize, Serialize)]
pub struct RegisterOpt {
    username: String,
    password: String,
}

#[post("/user/")]
pub async fn register(q: web::Json<RegisterOpt>, stat: web::Data<AppState>) -> Result<String> {
    let q = q.into_inner();
    // let mut tx = dbtx().await;
    info!("{:#?}", q);
    let u: Option<models::user::Model> =
        sqlx::query_as::<_, models::user::Model>("select * from user where username = ?")
            .bind(q.username.clone())
            .fetch_optional(stat.sqlx())
            .await?;
    let u: models::user::Model = match u {
        Some(_) => return Err(Error::ArgDuplicated(format!("username: {}", q.username))),
        None => {
            let mut u = models::user::Model::default();
            u.username = q.username.clone();
            u.id = uuid::Uuid::new_v4().to_string().replace("-", "");
            let p = match base64::decode(q.password.as_bytes()) {
                Err(_) => return Err(Error::ArgInvalid("password".to_string())),
                Ok(p) => p,
            };
            let p = match std::str::from_utf8(&p) {
                Ok(p) => p,
                Err(_) => return Err(Error::ArgInvalid("password".to_string())),
            };
            info!("{}", p);
            u.update_pass(&p)?;
            let mut rng = rand::thread_rng();
            let idx: i64 = rng.gen_range(1..221);
            u.icon = Some(format!("/media/icon/usr/{:04}.jpg", idx));
            u
        }
    };
    let oa: app::Model = sqlx::query_as::<_, app::Model>("select * from app where id = ?")
        .bind(stat.uuid.clone())
        .fetch_one(stat.sqlx())
        .await?;

    let mut au = app_user::Model::default();
    au.app_id = oa.id;
    au.user_id = u.id.clone();
    match oa.join_method.into() {
        models::AppJoin::Disabled => return Err(Error::AppDisabledRegister),
        models::AppJoin::Auto => {
            au.status = models::AUStatus::OK as i32;
        }
        models::AppJoin::Applying => au.status = models::AUStatus::Applying as i32,
    }
    let mut c = stat.sqlx().begin().await?;
    // 创建用户
    sqlx::query!(
        r#"
insert into user (id,username,_real_code,_check_code,icon)
values ( ?, ?, ?, ?, ?)
        "#,
        u.id,
        u.username,
        u.real_code,
        u.check_code,
        u.icon,
    )
    .execute(&mut c)
    .await?;

    // 关联应用
    sqlx::query!(
        r#"
        insert into app_user ( app_id, user_id, status)
        values (?, ?, ? )
        "#,
        au.app_id,
        au.user_id,
        au.status,
    )
    .execute(&mut c)
    .await?;
    if oa.role_id.is_some() {
        match au.status.into() {
            models::AUStatus::OK => {
                sqlx::query!(
                    r#"
        insert into user_role (user_id, role_id)
        values (?, ?)
        "#,
                    au.user_id,
                    oa.role_id.unwrap(),
                )
                .execute(&mut c)
                .await?;
            }
            _ => {}
        }
    }
    c.commit().await?;
    Ok("ok".to_string())
}

#[delete("/user/")]
pub async fn delete() -> impl Responder {
    ""
}
