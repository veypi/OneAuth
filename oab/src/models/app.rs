//
// app.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 00:18
// Distributed under terms of the Apache license.
//

use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Clone, sqlx::Type)]
#[repr(i64)]
pub enum AppJoin {
    Auto = 0,
    Disabled = 1,
    Applying = 2,
}

#[derive(Debug, Serialize, Deserialize, sqlx::FromRow)]
pub struct App {
    pub id: String,
    pub created: Option<NaiveDateTime>,
    pub updated: Option<NaiveDateTime>,
    pub delete_flag: bool,

    pub name: Option<String>,
    pub des: Option<String>,
    pub icon: Option<String>,
    pub user_count: i64,

    pub hide: bool,
    pub join_method: AppJoin,
    pub role_id: Option<String>,
    pub redirect: Option<String>,

    pub status: i64,
}

impl App {
    pub fn new() -> Self {
        Self {
            id: "".to_string(),
            created: None,
            updated: None,
            delete_flag: false,

            name: None,
            des: None,
            icon: None,
            user_count: 0,
            hide: false,
            join_method: AppJoin::Auto,
            role_id: None,
            redirect: None,
            status: 0,
        }
    }
}

#[derive(Debug, Deserialize, Serialize, Clone, sqlx::Type)]
#[repr(i64)]
pub enum AUStatus {
    OK = 0,
    Disabled = 1,
    Applying = 2,
    Deny = 3,
}

#[derive(Debug, Serialize, Deserialize, sqlx::FromRow)]
pub struct AppUser {
    pub created: Option<NaiveDateTime>,
    pub updated: Option<NaiveDateTime>,
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
