//
// user.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-06-02 23:03
// Distributed under terms of the MIT license.
//
//

use rbatis::crud_table;

#[crud_table]
#[derive(Debug, Clone)]
pub struct User {
    pub id: String,
    pub name: Option<String>,
}

impl Default for User {
    fn default() -> Self {
        User {
            id: rbatis::plugin::object_id::ObjectId::new().to_string(),
            name: None,
        }
    }
}
