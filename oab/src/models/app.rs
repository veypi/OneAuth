//
// app.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 00:18
// Distributed under terms of the Apache license.
//

use rbatis::{crud_table, DateTimeNative};

#[crud_table]
#[derive(Debug, Clone)]
pub struct App {
    pub id: String,
    pub created: Option<DateTimeNative>,
    pub updated: Option<DateTimeNative>,
    pub delete_flag: bool,

    pub name: Option<String>,
    pub des: Option<String>,
    pub icon: Option<String>,
    pub user_count: usize,

    pub hide: bool,
    pub register: bool,
    pub role_id: Option<String>,
    pub redirect: Option<String>,

    pub status: usize,
}

impl App {
    pub fn new() -> Self {
        Self {
            id: rbatis::plugin::object_id::ObjectId::new().to_string(),
            created: None,
            updated: None,
            delete_flag: false,

            name: None,
            des: None,
            icon: None,
            user_count: 0,
            hide: false,
            register: false,
            role_id: None,
            redirect: None,
            status: 0,
        }
    }
}

#[derive(Debug, serde::Serialize, serde::Deserialize, Clone)]
pub enum AUStatus {
    OK,
    Disabled,
    Applying,
    Deny,
}

#[crud_table]
#[derive(Debug, serde::Serialize, serde::Deserialize, Clone)]
pub struct AppUser {
    pub created: Option<DateTimeNative>,
    pub updated: Option<DateTimeNative>,
    pub app_id: String,
    pub user_id: String,
    pub status: AUStatus,
}

impl AppUser {
    pub fn new() -> Self {
        Self {
            created: None,
            updated: None,
            app_id: "".to_string(),
            user_id: "".to_string(),
            status: AUStatus::OK,
        }
    }
}
