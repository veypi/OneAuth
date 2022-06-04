//
// mod.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-06-02 23:04
// Distributed under terms of the MIT license.
//

mod user;
use rbatis::crud::CRUD;
use tracing::info;
use user::User;

use crate::cfg;

pub async fn init() {
    let mut u = User::default();
    u.name = Some("asd".to_string());
    cfg::DB.save(&u, &[]).await.unwrap();
    info!("{:#?}", u);
}
