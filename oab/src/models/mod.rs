//
// mod.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-06-02 23:04
// Distributed under terms of the MIT license.
//

pub mod app;
mod role;
mod user;
use std::{fs::File, io::Read};

use tracing::info;

use crate::DB;
pub use app::{AUStatus, App, AppUser};
pub use role::{Access, Resource, Role};
pub use user::User;

pub async fn init() {
    info!("init database");
    let mut f = File::open("./sql/table.sql").unwrap();
    let mut sql = String::new();
    f.read_to_string(&mut sql).unwrap();
    DB.exec(&sql, vec![]).await.unwrap();
}

pub fn new_id() -> rbatis::object_id::ObjectId {
    rbatis::plugin::object_id::ObjectId::new()
}
