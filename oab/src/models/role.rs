//
// role.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 02:42
// Distributed under terms of the Apache license.
//

use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};

#[derive(Debug, Default, Serialize, Deserialize)]
pub struct Role {
    pub id: String,
    pub created: Option<NaiveDateTime>,
    pub updated: Option<NaiveDateTime>,
    pub delete_flag: bool,

    pub app_id: String,
    pub name: Option<String>,
    pub des: Option<String>,
    pub user_count: usize,
}

#[derive(Debug, Default, Serialize, Deserialize)]
pub struct Resource {
    pub created: Option<NaiveDateTime>,
    pub updated: Option<NaiveDateTime>,
    pub delete_flag: bool,

    pub app_id: String,
    pub name: String,
    pub des: Option<String>,
}

#[derive(Debug, Default, Serialize, Deserialize)]
pub struct Access {
    pub created: Option<NaiveDateTime>,
    pub updated: Option<NaiveDateTime>,
    pub delete_flag: bool,

    pub app_id: String,
    pub name: String,
    pub role_id: Option<String>,
    pub user_id: Option<String>,
    pub rid: Option<String>,
    pub level: usize,
}
