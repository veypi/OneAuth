//
// mod.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-06-24 16:26
// Distributed under terms of the Apache license.
//
//

mod access;
mod app;
mod resource;
mod role;
mod user;
use crate::{Error, Result};
use actix_web::{get, web};

#[get("/hello/{name}")]
async fn greet(name: web::Path<u32>) -> Result<String> {
    let n = name.into_inner();
    if n > 0 {
        Ok(format!("Hello {n}!"))
    } else {
        Err(Error::Unknown)
    }
}

pub fn routes(cfg: &mut web::ServiceConfig) {
    cfg.service(user::get)
        .service(user::list)
        .service(user::register)
        .service(user::login)
        .service(user::delete);
    cfg.service(app::get)
        .service(app::list)
        .service(app::create)
        .service(app::del);
    cfg.service(greet);
}
