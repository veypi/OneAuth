//
// upload.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-10-03 21:50
// Distributed under terms of the MIT license.
//
//

use std::{fs, path::Path};

use actix_multipart::form::{tempfile::TempFile, MultipartForm};

use actix_web::{post, web, Responder};
use proc::access_read;
use tracing::{info, warn};

use crate::{models::Token, AppState, Error, Result};

#[derive(Debug, MultipartForm)]
struct UploadForm {
    files: Vec<TempFile>,
}

#[post("/upload/{dir:.*}")]
#[access_read("app")]
async fn save_files(
    MultipartForm(form): MultipartForm<UploadForm>,
    t: web::ReqData<Token>,
    dir: web::Path<String>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    let t = t.into_inner();
    let mut dir = dir.into_inner();
    if dir.is_empty() {
        dir = "tmp".to_string();
    }
    let mut res: Vec<String> = Vec::new();
    for v in form.files.into_iter() {
        let fname = v.file_name.clone().unwrap_or("unknown".to_string());
        let root = Path::new(&stat.media_path).join(dir.clone());
        if !root.exists() {
            match fs::create_dir_all(root.clone()) {
                Ok(_) => {}
                Err(e) => {
                    warn!("{}", e);
                }
            }
        }
        let temp_file = format!(
            "{}/{}.{}",
            root.to_str().unwrap_or(&stat.media_path),
            t.id,
            fname
        );
        info!("saving {:?} to {temp_file}", v.file.path().to_str());
        match fs::copy(v.file.path(), &temp_file) {
            Ok(p) => {
                info!("{:#?}", p);
                res.push(format!("/media/{}/{}.{}", dir, t.id, fname))
            }
            Err(e) => {
                warn!("{}", e);
            }
        };
    }

    Ok(web::Json(res))
}
