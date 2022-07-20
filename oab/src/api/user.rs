//
// user.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//

use crate::{dbtx, models, Error, Result, CONFIG, DB};
use actix_web::{delete, get, head, post, web, Responder};
use base64;
use rbatis::crud::CRUD;
use serde::{Deserialize, Serialize};
use tracing::info;

#[get("/user/{id}")]
pub async fn get(id: web::Path<String>) -> Result<models::User> {
    let n = id.into_inner();
    if !n.is_empty() {
        let u: Option<models::User> = DB.fetch_by_column("id", &n).await?;
        match u {
            Some(u) => {
                info!("{:#?}", u.token());
                return Ok(u);
            }
            None => Err(Error::NotFound(format!("user {}", n))),
        }
    } else {
        Err(Error::Missing("id".to_string()))
    }
}

#[get("/user/")]
pub async fn list() -> impl Responder {
    let result: Vec<models::User> = DB.fetch_list().await.unwrap();
    web::Json(result)
}

#[derive(Debug, Deserialize, Serialize)]
pub struct LoginOpt {
    typ: Option<String>,
    password: String,
}

#[head("/user/{id}")]
pub async fn login(q: web::Query<LoginOpt>, id: web::Path<String>) -> impl Responder {
    let id = id.into_inner();
    let q = q.into_inner();
    info!("{} try to login{:#?}", id, q);
    let mut w = DB.new_wrapper();
    match q.typ {
        _ => w = w.eq("username", id),
    }
    let u: Option<models::User> = DB.fetch_by_wrapper(w).await.unwrap();
    info!("{:#?}", u);

    ""
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
    let u: Option<models::User> = DB.fetch_by_column("username", &q.username).await.unwrap();
    let u: models::User = match u {
        Some(_) => return Err(Error::ArgDuplicated(format!("username: {}", q.username))),
        None => {
            let mut u = models::User::default();
            u.username = q.username.clone();
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
            u
        }
    };
    let oa: Option<models::App> = DB.fetch_by_column("id", CONFIG.uuid.clone()).await?;
    let oa = oa.unwrap();
    let mut au = models::AppUser::new();
    au.app_id = oa.id;
    au.user_id = u.id.clone();
    match oa.join_method {
        models::app::AppJoin::Disabled => return Err(Error::AppDisabledRegister),
        models::app::AppJoin::Auto => au.status = models::app::AUStatus::OK,
        models::app::AppJoin::Applying => au.status = models::app::AUStatus::Applying,
    }
    DB.save(&u, &[]).await?;
    DB.save(&au, &[]).await?;
    Ok("ok".to_string())
}

#[delete("/user/")]
pub async fn delete() -> impl Responder {
    ""
}
