//
// app.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 03:10
// Distributed under terms of the Apache license.
//
//
use actix_web::{delete, get, post, web, Responder};

use crate::{models, Error, Result, CONFIG};

#[get("/app/{id}")]
pub async fn get(id: web::Path<String>) -> Result<impl Responder> {
    let n = id.into_inner();
    if !n.is_empty() {
        let s = sqlx::query_as::<_, models::App>("select * from app where id = ?")
            .bind(n)
            .fetch_one(CONFIG.db())
            .await?;
        Ok(web::Json(s))
    } else {
        Err(Error::Missing("id".to_string()))
    }
}

#[get("/app/")]
pub async fn list() -> Result<impl Responder> {
    let result = sqlx::query_as::<_, models::App>("select * from app")
        .fetch_all(CONFIG.db())
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
