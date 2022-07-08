//
// mod.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-06-24 16:26
// Distributed under terms of the Apache license.
//
//
use crate::{Error, Result};
use actix_web::{get, web, Responder};

#[get("/hello/{name}")]
async fn greet(name: web::Path<u32>) -> Result<String> {
    let n = name.into_inner();
    if n > 0 {
        Ok(format!("Hello {n}!"))
    } else {
        Err(Error::Unknown)
    }
}

#[get("/topic/derive")]
async fn hello() -> impl Responder {
    "Hello World!"
}

pub fn routes(cfg: &mut web::ServiceConfig) {
    cfg.service(greet);
    cfg.service(hello);
}
