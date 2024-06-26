//
// auth.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-09-01 17:39
// Distributed under terms of the Apache license.
//

use std::cell::RefCell;
use std::pin::Pin;
use std::rc::Rc;
use std::task::{Context, Poll};

use actix_web::body::MessageBody;
use actix_web::dev::{Service, ServiceRequest, ServiceResponse, Transform};

use actix_web::{Error, HttpMessage};
use futures_util::future::{ok, Ready};
use futures_util::Future;
use tracing::warn;

use crate::models;

// custom request auth middleware
pub struct Auth {
    pub key: String,
}

impl<S, B> Transform<S, ServiceRequest> for Auth
where
    S: Service<ServiceRequest, Response = ServiceResponse<B>, Error = Error> + 'static,
    S::Future: 'static,
    B: MessageBody + 'static,
{
    type Response = ServiceResponse<B>;
    type Error = Error;
    type Transform = AuthMiddleware<S>;
    type InitError = ();
    type Future = Ready<Result<Self::Transform, Self::InitError>>;

    fn new_transform(&self, service: S) -> Self::Future {
        ok(AuthMiddleware {
            key: self.key.clone(),
            service: Rc::new(RefCell::new(service)),
        })
    }
}

pub struct AuthMiddleware<S> {
    service: Rc<RefCell<S>>,
    key: String,
}

impl<S, B> Service<ServiceRequest> for AuthMiddleware<S>
where
    S: Service<ServiceRequest, Response = ServiceResponse<B>, Error = Error> + 'static,
    S::Future: 'static,
    B: MessageBody + 'static,
{
    type Response = ServiceResponse<B>;
    type Error = Error;
    type Future = Pin<Box<dyn Future<Output = Result<Self::Response, Self::Error>>>>;

    fn poll_ready(&self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.service.poll_ready(cx)
    }

    fn call(&self, req: ServiceRequest) -> Self::Future {
        let svc = self.service.clone();
        let key = self.key.clone();

        Box::pin(async move {
            match req.headers().get("auth_token") {
                Some(h) => match models::Token::from(h.to_str().unwrap_or(""), &key) {
                    Ok(t) => {
                        req.extensions_mut().insert(t.id.clone());
                        req.extensions_mut().insert(t);
                    }
                    Err(e) => warn!("{}", e),
                },
                None => {}
            }
            // let value = HeaderValue::from_str("").unwrap();
            // let token = req.headers().get("auth_token").unwrap_or(&value);
            // let token = models::Token::from(token.to_str().unwrap_or(""));
            // match token {
            //     Ok(t) => {
            //         req.extensions_mut().insert(t.id.clone());
            //         req.extensions_mut().insert(t);
            //     }
            // };
            Ok(svc.call(req).await?)
        })
    }
}
