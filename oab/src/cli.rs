//
// cli.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-07-08 18:39
// Distributed under terms of the Apache license.
//

use clap::{Args, Parser, Subcommand};

use lazy_static::lazy_static;

lazy_static! {
    pub static ref CLI: cli = cli::new();
}

#[derive(Debug, Parser)]
#[clap(name = "oab")]
#[clap(about = "oab", long_about = None)]
struct cli {
    #[clap(subcommand)]
    command: Commands,
}

#[derive(Debug, Subcommand)]
enum Commands {
    /// Clones repos
    #[clap(arg_required_else_help = true)]
    Clone {
        /// The remote to clone
        #[clap(value_parser)]
        remote: String,
    },
    /// pushes things
    #[clap(arg_required_else_help = true)]
    Push {
        /// The remote to target
        #[clap(value_parser)]
        remote: String,
    },
    /// adds things
    #[clap(arg_required_else_help = true)]
    Add {
        /// Stuff to add
        #[clap(required = true, value_parser)]
        path: Vec<PathBuf>,
    },
    Stash(Stash),
    #[clap(external_subcommand)]
    External(Vec<OsString>),
}

impl cli {
    fn new() -> Self {
        cli::parse()
    }
}
