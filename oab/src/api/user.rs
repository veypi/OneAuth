//
// user.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//

use std::fmt::Debug;

use crate::{models, Error, Result, CONFIG};
use actix_web::{delete, get, head, http, post, web, HttpResponse, Responder};
use base64;
use proc::access_read;
use rand::Rng;
use serde::{Deserialize, Serialize};
use tracing::info;

#[get("/user/{id}")]
#[access_read("user", id = "&id.clone()")]
pub async fn get(id: web::Path<String>) -> Result<impl Responder> {
    let n = id.into_inner();
    if !n.is_empty() {
        let s = sqlx::query!(
            "select id,updated,created,username,nickname,email,icon,status, used, space from user where id = ?",n
        ).map(|row| models::User {
            id: row.id,
        created: row.created,
        updated: row.updated,
        username: row.username,
        nickname: row.nickname,
        email: row.email,
        status: row.status,
        used: row.used,
        space: row.space.unwrap_or(0),
        icon: row.icon,
            ..Default::default()
        })
        .fetch_one(CONFIG.db())
        .await?;
        Ok(web::Json(s))
    } else {
        Err(Error::Missing("id".to_string()))
    }
}

#[get("/user/")]
#[access_read("user")]
pub async fn list() -> Result<impl Responder> {
    let result = sqlx::query!(
        "select id,updated,created,username,nickname,email,icon,status, used, space from user",
    )
    .map(|row| models::User {
        id: row.id,
        created: row.created,
        updated: row.updated,
        username: row.username,
        nickname: row.nickname,
        email: row.email,
        status: row.status,
        used: row.used,
        space: row.space.unwrap_or(0),
        icon: row.icon,
        ..Default::default()
    })
    .fetch_all(CONFIG.db())
    .await?;
    Ok(web::Json(result))
}

#[derive(Debug, Deserialize, Serialize)]
pub struct LoginOpt {
    typ: Option<String>,
    password: String,
}

#[head("/user/{id}")]
pub async fn login(q: web::Query<LoginOpt>, id: web::Path<String>) -> Result<HttpResponse> {
    let id = id.into_inner();
    let q = q.into_inner();
    let typ = match q.typ {
        Some(t) => match t.as_str() {
            "phone" => "phone",
            "email" => "email",
            _ => "username",
        },
        _ => "username",
    };
    let p = match base64::decode(q.password.as_bytes()) {
        Err(_) => return Err(Error::ArgInvalid("password".to_string())),
        Ok(p) => p,
    };
    let p = match std::str::from_utf8(&p) {
        Ok(p) => p,
        Err(_) => return Err(Error::ArgInvalid("password".to_string())),
    };
    let sql = format!("select * from user where {} = ?", typ);
    let u = sqlx::query_as::<_, models::User>(&sql)
        .bind(id)
        .fetch_optional(CONFIG.db())
        .await?;
    let u = match u {
        Some(u) => u,
        None => return Err(Error::NotFound("user".to_string())),
    };

    u.check_pass(p)?;
    let au = sqlx::query_as::<_, models::AppUser>(
        "select * from app_user where app_id = ? and user_id = ?",
    )
    .bind(&CONFIG.uuid)
    .bind(&u.id)
    .fetch_optional(CONFIG.db())
    .await?;
    let i: i64 = match au {
        Some(au) => match au.status {
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
            let app = sqlx::query_as::<_, models::App>("select * from app where id = ?")
                .bind(CONFIG.uuid.clone())
                .fetch_one(CONFIG.db())
                .await?;
            info!("{:#?}", u);
            let s = match app.join_method {
                models::app::AppJoin::Disabled => {
                    return Err(Error::BusinessException(
                        "this app diabled login".to_string(),
                    ))
                }
                models::app::AppJoin::Auto => models::AUStatus::OK,
                models::app::AppJoin::Applying => models::AUStatus::Applying,
            };
            sqlx::query(
                r#"
insert into app_user (app_id,user_id,status)
values ( ?, ?, ? )
        "#,
            )
            .bind(&app.id)
            .bind(&u.id)
            .bind(&s)
            .execute(CONFIG.db())
            .await?;
            match s {
                models::AUStatus::OK => 0,
                _ => 1,
            }
        }
    };
    if i == 0 {
        let result = sqlx::query_as::<_, models::AccessCore>(
            "select access.name,access.rid,access.level from access, user_role, role WHERE user_role.user_id = ? && access.role_id=user_role.role_id && role.id=user_role.role_id && role.app_id = ?",
        )
        .bind(&u.id)
        .bind(CONFIG.uuid.clone())
        .fetch_all(CONFIG.db())
        .await?;
        Ok(HttpResponse::build(http::StatusCode::OK)
            .insert_header(("auth_token", u.token(result).to_string()?))
            .body("".to_string()))
    } else {
        Ok(HttpResponse::build(http::StatusCode::OK)
            .insert_header(("data", "applying"))
            .body("".to_string()))
    }
}

#[derive(Debug, Deserialize, Serialize)]
pub struct RegisterOpt {
    username: String,
    password: String,
}

#[post("/user/")]
pub async fn register(q: web::Json<RegisterOpt>) -> Result<String> {
    let q = q.into_inner();
    // let mut tx = dbtx().await;
    println!("{:#?}", q);
    let u: Option<models::User> =
        sqlx::query_as::<_, models::User>("select * from user where username = ?")
            .bind(q.username.clone())
            .fetch_optional(CONFIG.db())
            .await?;
    let u: models::User = match u {
        Some(_) => return Err(Error::ArgDuplicated(format!("username: {}", q.username))),
        None => {
            let mut u = models::User::default();
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
    let oa: models::App = sqlx::query_as::<_, models::App>("select * from app where id = ?")
        .bind(CONFIG.uuid.clone())
        .fetch_one(CONFIG.db())
        .await?;

    let mut au = models::AppUser::new();
    au.app_id = oa.id;
    au.user_id = u.id.clone();
    match oa.join_method {
        models::app::AppJoin::Disabled => return Err(Error::AppDisabledRegister),
        models::app::AppJoin::Auto => {
            au.status = models::app::AUStatus::OK;
        }
        models::app::AppJoin::Applying => au.status = models::app::AUStatus::Applying,
    }
    let mut c = CONFIG.db().begin().await?;
    // 创建用户
    sqlx::query!(
        r#"
insert into user (id,username,real_code,check_code,icon)
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
        match au.status {
            models::app::AUStatus::OK => {
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
