//
// cli.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-08 18:39
// Distributed under terms of the Apache license.
//

use crate::cfg::AppState;
use crate::Result;
use clap::{Args, Parser, Subcommand};
use tracing::info;

#[derive(Debug, Parser)]
#[clap(name = "oab")]
#[clap(about = "oab", long_about = None)]
pub struct AppCli {
    #[clap(short = 'c', value_name = "cfg",default_value_t = String::from("~/.config/oa/oab.yml"), value_hint = clap::ValueHint::DirPath)]
    pub cfg: String,
    #[clap(short = 'n', value_name = "name",default_value_t = String::from("v.oab"))]
    pub name: String,
    #[clap(subcommand)]
    pub command: Option<Clis>,
}

#[derive(Debug, Subcommand)]
pub enum Clis {
    Init,
    Install,
    Uninstall,
    Start,
    Stop,
    Web,
    Dump,
    Cfg(CfgOpt),
}

#[derive(Debug, Args)]
#[clap(args_conflicts_with_subcommands = true)]
pub struct CfgOpt {
    command: Option<String>,
}

impl AppCli {
    pub fn new() -> Self {
        AppCli::parse()
    }
    pub fn handle_service(&self, data: AppState) -> Result<bool> {
        let label: service_manager::ServiceLabel = self.name.parse().unwrap();

        // Get generic service by detecting what is available on the platform
        let manager = <dyn service_manager::ServiceManager>::native()
            .expect("Failed to detect management platform");

        if let Some(c) = &self.command {
            match c {
                Clis::Install => {
                    let p = std::env::current_exe()?;
                    info!("deploy {} -c {}", p.to_str().unwrap(), self.cfg);
                    manager.install(service_manager::ServiceInstallCtx {
                        label: label.clone(),
                        program: p,
                        args: vec![format!("-c {}", self.cfg).into()],
                        contents: None, // Optional String for system-specific service content.
                    })?
                }
                Clis::Uninstall => manager.uninstall(service_manager::ServiceUninstallCtx {
                    label: label.clone(),
                })?,
                Clis::Start => manager.start(service_manager::ServiceStartCtx {
                    label: label.clone(),
                })?,
                Clis::Stop => manager.stop(service_manager::ServiceStopCtx {
                    label: label.clone(),
                })?,
                Clis::Dump => {
                    let res = serde_yaml::to_string(&data)?;
                    println!("{}", res);
                }
                _ => return Ok(false),
            }
            return Ok(true);
        };
        Ok(false)
    }
}
