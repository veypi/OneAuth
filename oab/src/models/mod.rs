//
// mod.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-06-02 23:04
// Distributed under terms of the MIT license.
//

pub mod app;
mod role;
mod user;

use tracing::info;

pub use app::{AUStatus, App, AppJoin, AppUser};
pub use role::{Access, Resource, Role};
pub use user::{AccessCore, AccessLevel, Token, User};

use crate::CONFIG;

pub async fn init() {
    info!("init database");
    sqlx::migrate!().run(CONFIG.db()).await.unwrap();
}
