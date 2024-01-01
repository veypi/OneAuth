//
// lib.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-05-29 18:48
// Distributed under terms of the MIT license.
//

pub mod api;
mod cfg;
pub mod fs;
pub mod libs;
pub mod models;
mod result;
pub use cfg::{init_log, AppCli, AppState, ApplicationConfig, Clis, InfoOpt};
pub use result::{Error, Result};
