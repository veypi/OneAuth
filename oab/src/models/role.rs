//
// role.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 02:42
// Distributed under terms of the Apache license.
//
use rbatis::{crud_table, DateTimeNative};

#[crud_table]
#[derive(Debug, Clone, Default)]
pub struct Role {
    pub id: String,
    pub created: Option<DateTimeNative>,
    pub updated: Option<DateTimeNative>,
    pub delete_flag: bool,

    pub app_id: String,
    pub name: Option<String>,
    pub des: Option<String>,
    pub user_count: usize,
}

#[crud_table]
#[derive(Debug, Clone, Default)]
pub struct Resource {
    pub created: Option<DateTimeNative>,
    pub updated: Option<DateTimeNative>,
    pub delete_flag: bool,

    pub app_id: String,
    pub name: String,
    pub des: Option<String>,
}

#[crud_table]
#[derive(Debug, Clone, Default)]
pub struct Access {
    pub created: Option<DateTimeNative>,
    pub updated: Option<DateTimeNative>,
    pub delete_flag: bool,

    pub app_id: String,
    pub name: String,
    pub role_id: Option<String>,
    pub user_id: Option<String>,
    pub rid: Option<String>,
    pub level: usize,
}
