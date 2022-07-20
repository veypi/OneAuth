//
// result.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-06-24 18:57
// Distributed under terms of the Apache license.
//

use actix_web::http::header;
use actix_web::middleware::ErrorHandlerResponse;
use actix_web::ResponseError;
use actix_web::{
    dev, error,
    http::{header::ContentType, StatusCode},
    HttpResponse,
};

use serde::{Deserialize, Serialize};
use thiserror::Error as ThisError;

pub type Result<T> = std::result::Result<T, Error>;

// pub type AsyncResult<T> = std::result::Result<T, Box<dyn std::error::Error + Send + Sync>>;

#[derive(Clone, ThisError, Debug, Deserialize, Serialize)]
pub enum Error {
    // system
    // EnvVarError,
    #[error("Parsing listening address failed")]
    ParseListeningAddressFailed,
    #[error("Data save failed")]
    SledSaveFailed,
    #[error("Database(1) error")]
    SledDbError,
    #[error("Database(2) error")]
    SqliteDbError,
    #[error("Deserialize / Serialize failed")]
    SerdeError,

    #[error("register disabled")]
    AppDisabledRegister,

    #[error("missing {0}")]
    Missing(String),
    #[error("invalid arg {0}")]
    ArgInvalid(String),
    #[error("duplicated arg {0}")]
    ArgDuplicated(String),

    #[error("not found {0}")]
    NotFound(String),

    #[error("timeout")]
    Timeout,

    #[error("bad request")]
    BadRequest,
    #[error("Method not allowed")]
    MethodNotAllowed,
    #[error("Internal server error")]
    InternalServerError,

    // business
    #[error("无效的 Session ID")]
    InvalidSessionId,
    #[error("invalid verify code")]
    InvalidVerifyCode,
    #[error("no access")]
    NotAuthed,
    #[error("login failed")]
    LoginFailed,
    #[error("Registration failed")]
    RegisterFailed,
    #[error("Already registered")]
    AlreadyRegistered,
    #[error("Saving post failed")]
    SavePostFailed,
    #[error("Can not find post you requested")]
    CannotFoundPost,
    #[error("Can not find tag you requested")]
    CannotFoundTag,
    #[error("Upload failed")]
    UploadFailed,
    #[error("Upload file not found")]
    FileNotFound,
    #[error("Unknown file type")]
    UnknownFileType,
    #[error("Unsupported file type {0}")]
    UnsupportedFileType(String),
    #[error("Creating thumbnail failed")]
    CreateThumbnailFailed,
    #[error("Reading post id data by tag failed")]
    ReadPostIdDataByTagFailed,
    #[error("Saving post id data by tag failed")]
    SavePostIdDataByTagFailed,
    #[error("Tag not found")]
    TagNotFound,

    #[error("{0}")]
    BusinessException(String),
    #[error("invalid header (expected {expected:?}, found {found:?})")]
    InvalidHeader { expected: String, found: String },

    #[error("unknown error")]
    Unknown,
}

impl From<std::io::Error> for Error {
    fn from(e: std::io::Error) -> Self {
        Error::UnsupportedFileType(format!("{:?}", e))
    }
}
impl From<actix_web::Error> for Error {
    fn from(e: actix_web::Error) -> Self {
        Error::BusinessException(format!("{:?}", e))
    }
}
impl From<jsonwebtoken::errors::Error> for Error {
    fn from(e: jsonwebtoken::errors::Error) -> Self {
        Error::BusinessException(format!("{:?}", e))
    }
}
impl From<rbatis::error::Error> for Error {
    fn from(e: rbatis::error::Error) -> Self {
        Error::BusinessException(format!("{:?}", e))
    }
}
impl From<aes_gcm::Error> for Error {
    fn from(e: aes_gcm::Error) -> Self {
        Error::BusinessException(format!("{:?}", e))
    }
}

impl From<Box<dyn std::fmt::Display>> for Error {
    fn from(e: Box<dyn std::fmt::Display>) -> Self {
        Error::BusinessException(format!("{}", e))
    }
}

impl actix_web::Responder for Error {
    type Body = actix_web::body::BoxBody;
    fn respond_to(self, _req: &actix_web::HttpRequest) -> HttpResponse<Self::Body> {
        self.error_response()
    }
}

impl error::ResponseError for Error {
    fn error_response(&self) -> HttpResponse {
        HttpResponse::build(self.status_code())
            .insert_header(ContentType::html())
            .body(self.to_string())
    }

    fn status_code(&self) -> StatusCode {
        match *self {
            Error::InternalServerError => StatusCode::INTERNAL_SERVER_ERROR,
            Error::BadRequest => StatusCode::BAD_REQUEST,
            Error::Timeout => StatusCode::GATEWAY_TIMEOUT,
            Error::NotFound(_) => StatusCode::NOT_FOUND,
            Error::NotAuthed => StatusCode::UNAUTHORIZED,
            _ => StatusCode::BAD_REQUEST,
        }
    }
}