//
// mod.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-06-02 23:04
// Distributed under terms of the MIT license.
//

mod app_plugin;
pub mod entity;
mod role;
mod user_plugin;

use tracing::info;

pub use app_plugin::{AUStatus, AppJoin};
pub use entity::{access, app, app_user, user};
pub use user_plugin::{AccessLevel, Token, UserPlugin};

use crate::AppState;

pub async fn init(data: AppState) {
    info!("init database");
    sqlx::migrate!().run(data.sqlx()).await.unwrap();
}
