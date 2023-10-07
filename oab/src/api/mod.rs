//
// mod.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-06-24 16:26
// Distributed under terms of the Apache license.
//
//

mod access;
mod app;
mod appuser;
mod resource;
mod role;
mod upload;
mod user;
use actix_web::web;

pub fn routes(cfg: &mut web::ServiceConfig) {
    cfg.service(upload::save_files);
    cfg.service(user::get)
        .service(user::list)
        .service(user::register)
        .service(user::login)
        .service(user::delete);
    cfg.service(app::get)
        .service(app::list)
        .service(app::create)
        .service(app::update)
        .service(app::del);
    // cfg.route("/acc", web::get().to(access::UpdateOpt::update));

    cfg.service(appuser::get).service(appuser::add);
}
