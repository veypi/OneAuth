//
// user.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-06-02 23:03
// Distributed under terms of the MIT license.
//
//

use rbatis::{crud_table, DateTimeNative};

#[crud_table]
#[derive(Debug, Clone)]
pub struct User {
    pub id: String,
    pub created: Option<DateTimeNative>,
    pub updated: Option<DateTimeNative>,
    pub delete_flag: bool,

    pub username: Option<String>,
    pub nickname: Option<String>,
    pub email: Option<String>,
    pub phone: Option<String>,
    pub icon: Option<String>,
    pub status: usize,
    pub used: usize,
    pub space: usize,
}

impl Default for User {
    fn default() -> Self {
        Self {
            id: rbatis::plugin::object_id::ObjectId::new().to_string(),
            created: None,
            updated: None,
            delete_flag: false,

            username: None,
            nickname: None,
            email: None,
            phone: None,
            icon: None,
            status: 0,
            used: 0,
            space: 300,
        }
    }
}
