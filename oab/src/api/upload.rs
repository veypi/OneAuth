//
// upload.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-10-03 21:50
// Distributed under terms of the MIT license.
//
//

use actix_multipart::form::{tempfile::TempFile, MultipartForm};

use actix_web::{post, web, Responder};
use proc::access_read;
use tracing::{info, warn};

use crate::{models::Token, AppState, Error, Result};

#[derive(Debug, MultipartForm)]
struct UploadForm {
    files: Vec<TempFile>,
}

#[post("/upload/")]
#[access_read("app")]
async fn save_files(
    MultipartForm(form): MultipartForm<UploadForm>,
    t: web::ReqData<Token>,
    stat: web::Data<AppState>,
) -> Result<impl Responder> {
    let l = form.files.len();
    let t = t.into_inner();
    let mut res: Vec<String> = Vec::new();
    info!("!|||||||||||_{}_|", l);
    form.files.into_iter().for_each(|v| {
        let fname = v.file_name.unwrap_or("unknown".to_string());
        let path = format!("{}tmp/{}.{}", stat.media_path, t.id, fname);
        info!("saving to {path}");
        match v.file.persist(path) {
            Ok(p) => {
                info!("{:#?}", p);
                res.push(format!("/media/tmp/{}.{}", t.id, fname))
            }
            Err(e) => {
                warn!("{}", e);
                // return Err(Error::InternalServerError);
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
