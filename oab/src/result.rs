//
// result.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-06-24 18:57
// Distributed under terms of the Apache license.
//

use actix_web::ResponseError;
use actix_web::{
    error,
    http::{header::ContentType, StatusCode},
    HttpResponse,
};

use serde::{Deserialize, Serialize};
use thiserror::Error as ThisError;
use tracing::warn;

pub type Result<T> = std::result::Result<T, Error>;

#[derive(Serialize, Deserialize)]
pub struct JsonResponse<T> {
    pub content: T,
}
impl<T> From<T> for JsonResponse<T> {
    fn from(e: T) -> Self {
        Self { content: e }
    }
}
impl<T> actix_web::Responder for JsonResponse<T>
where
    T: serde::Serialize,
{
    type Body = actix_web::body::BoxBody;
    fn respond_to(self, _req: &actix_web::HttpRequest) -> HttpResponse<Self::Body> {
        HttpResponse::build(StatusCode::OK)
            .insert_header(ContentType::json())
            .body(serde_json::to_string(&self.content).unwrap())
    }
}

// pub type AsyncResult<T> = std::result::Result<T, Box<dyn std::error::Error + Send + Sync>>;

#[derive(ThisError, Debug)]
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
    #[error("invalid 的 Session ID")]
    InvalidSessionId,
    #[error("invalid verify code")]
    InvalidVerifyCode,
    #[error("invalid token")]
    InvalidToken,
    #[error("expired token")]
    ExpiredToken,

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

    #[error(transparent)]
    Other(#[from] anyhow::Error),
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
impl From<sqlx::Error> for Error {
    fn from(e: sqlx::Error) -> Self {
        Error::BusinessException(format!("{:?}", e))
    }
}
impl From<jsonwebtoken::errors::Error> for Error {
    fn from(e: jsonwebtoken::errors::Error) -> Self {
        Error::BusinessException(format!("{:?}", e))
    }
}

impl From<sea_orm::DbErr> for Error {
    fn from(e: sea_orm::DbErr) -> Self {
        Error::BusinessException(format!("{:?}", e))
    }
}
impl From<aes_gcm::Error> for Error {
    fn from(e: aes_gcm::Error) -> Self {
        Error::BusinessException(format!("{:?}", e))
    }
}

impl From<actix_multipart::MultipartError> for Error {
    fn from(e: actix_multipart::MultipartError) -> Self {
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
        warn!("{}", self);
        HttpResponse::build(self.status_code())
            .insert_header(ContentType::html())
            .insert_header(("error", self.to_string()))
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
