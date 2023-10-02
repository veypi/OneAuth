//
// app.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-09 00:18
// Distributed under terms of the Apache license.
//

use serde_repr::*;

#[derive(Debug, Serialize_repr, Deserialize_repr, Clone, sqlx::Type)]
#[repr(i32)]
pub enum AppJoin {
    Auto = 0,
    Disabled = 1,
    Applying = 2,
}
impl From<i32> for AppJoin {
    fn from(v: i32) -> Self {
        match v {
            x if x == AppJoin::Auto as i32 => AppJoin::Auto,
            x if x == AppJoin::Disabled as i32 => AppJoin::Disabled,
            x if x == AppJoin::Applying as i32 => AppJoin::Applying,
            _ => AppJoin::Auto,
        }
    }
}

#[derive(Debug, Deserialize_repr, Serialize_repr, Clone, sqlx::Type)]
#[repr(i32)]
pub enum AUStatus {
    OK = 0,
    Disabled = 1,
    Applying = 2,
    Deny = 3,
}

impl From<i32> for AUStatus {
    fn from(v: i32) -> Self {
        match v {
            x if x == AUStatus::OK as i32 => AUStatus::OK,
            x if x == AUStatus::Disabled as i32 => AUStatus::Disabled,
            x if x == AUStatus::Applying as i32 => AUStatus::Applying,
            x if x == AUStatus::Deny as i32 => AUStatus::Deny,
            _ => AUStatus::OK,
        }
    }
}
