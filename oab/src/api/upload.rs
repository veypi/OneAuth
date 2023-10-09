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
    form.files.into_iter().for_each(|v| {
        let fname = v.file_name.unwrap_or("unknown".to_string());
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
        info!("saving to {temp_file}");
        match v.file.persist(temp_file) {
            Ok(p) => {
                info!("{:#?}", p);
                res.push(format!("/media/{}/{}.{}", dir, t.id, fname))
            }
            Err(e) => {
                warn!("{}", e);
            }
        };
    });

    Ok(web::Json(res))
}

// #[actix_web::main]
// async fn main() -> std::io::Result<()> {
//     HttpServer::new(|| {
//         App::new()
//             .wrap(middleware::Logger::default())
//             .app_data(TempFileConfig::default().directory("./tmp"))
//             .service(
//                 web::resource("/")
//                     .route(web::get().to(index))
//                     .route(web::post().to(save_files)),
//             )
//     })
//     .bind(("127.0.0.1", 8080))?
//     .workers(2)
//     .run()
//     .await
// }

// /// Example of the old manual way of processing multipart forms.
// #[allow(unused)]
// async fn save_file_manual(mut payload: Multipart) -> Result<HttpResponse, Error> {
//     // iterate over multipart stream
//     while let Some(mut field) = payload.try_next().await? {
//         // A multipart/form-data stream has to contain `content_disposition`
//         let content_disposition = field.content_disposition();

//         let filename = content_disposition
//             .get_filename()
//             .map_or_else(|| Uuid::new_v4().to_string(), sanitize_filename::sanitize);
//         let filepath = format!("./tmp/{filename}");

//         // File::create is blocking operation, use threadpool
//         let mut f = web::block(|| std::fs::File::create(filepath)).await??;

//         // Field in turn is stream of *Bytes* object
//         while let Some(chunk) = field.try_next().await? {
//             // filesystem operations are blocking, we have to use threadpool
//             f = web::block(move || f.write_all(&chunk).map(|_| f)).await??;
//         }
//     }

//     Ok(HttpResponse::Ok().into())
// }
