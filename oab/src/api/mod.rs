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
mod token;
mod upload;
mod user;
use actix_web::web;

pub fn routes(cfg: &mut web::ServiceConfig) {
    cfg.service(upload::save_files);
    cfg.service(user::get)
        .service(user::list)
        .service(user::register)
        .service(user::login)
        .service(user::update)
        .service(user::delete);
    cfg.service(app::get)
        .service(app::get_key)
        .service(app::list)
        .service(app::create)
        .service(app::update)
        .service(app::del);
    // cfg.route("/acc", web::get().to(access::UpdateOpt::update));
    cfg.service(token::get);

    cfg.service(appuser::get)
        .service(appuser::add)
        .service(appuser::update);

    cfg.service(access::list)
        .service(access::creat)
        .service(access::update)
        .service(access::delete);
    cfg.service(resource::list)
        .service(resource::create)
        .service(resource::update)
        .service(resource::delete);
    cfg.service(role::list)
        .service(role::create)
        .service(role::update)
        .service(role::delete)
        .service(role::add)
        .service(role::drop);
}
