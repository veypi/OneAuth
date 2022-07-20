//
// lib.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-05-29 18:48
// Distributed under terms of the MIT license.
//

pub mod api;
mod cfg;
pub mod libs;
pub mod models;
mod result;
pub use cfg::{dbtx, init_log, ApplicationConfig, Clis, CLI, CONFIG, DB};
pub use result::{Error, Result};
