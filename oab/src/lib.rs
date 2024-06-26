//
// lib.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-05-29 18:48
// Distributed under terms of the MIT license.
//

pub mod api;
mod cfg;
mod cli;
pub mod fs;
pub mod libs;
pub mod models;
mod result;
pub use cfg::{init_log, AppState, ApplicationConfig, InfoOpt};
pub use cli::{AppCli, Clis};
pub use result::{Error, Result};
