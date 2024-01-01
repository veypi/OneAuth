//
// mod.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-11-07 00:07
// Distributed under terms of the MIT license.
//

mod app;
mod usr;
use std::future::{ready, Ready};

use actix_web::{
    dev::{forward_ready, Service, ServiceRequest, ServiceResponse, Transform},
    web, Error,
};
use futures_util::future::LocalBoxFuture;

pub fn routes(cfg: &mut web::ServiceConfig) {
    cfg.service(
        actix_web::web::scope("u")
            .app_data(web::Data::new(usr::client()))
            .service(web::resource("/{tail:.*}").to(usr::dav_handler)),
    );
    cfg.service(
        actix_web::web::scope("a")
            .app_data(web::Data::new(app::client()))
            .service(web::resource("/{id}/{tail:.*}").to(app::dav_handler)),
    );
}

pub struct FsWrap;

impl<S, B> Transform<S, ServiceRequest> for FsWrap
where
    S: Service<ServiceRequest, Response = ServiceResponse<B>, Error = Error>,
    S::Future: 'static,
    B: 'static,
{
    type Response = ServiceResponse<B>;
    type Error = Error;
    type InitError = ();
    type Transform = FsMiddleware<S>;
    type Future = Ready<Result<Self::Transform, Self::InitError>>;

    fn new_transform(&self, service: S) -> Self::Future {
        ready(Ok(FsMiddleware { service }))
    }
}

pub struct FsMiddleware<S> {
    service: S,
}

impl<S, B> Service<ServiceRequest> for FsMiddleware<S>
where
    S: Service<ServiceRequest, Response = ServiceResponse<B>, Error = Error>,
    S::Future: 'static,
    B: 'static,
{
    type Response = ServiceResponse<B>;
    type Error = Error;
    type Future = LocalBoxFuture<'static, Result<Self::Response, Self::Error>>;

    forward_ready!(service);

    fn call(&self, req: ServiceRequest) -> Self::Future {
        println!("start fs: {}", req.path());
        let reqheaders = req.headers().clone();
        let is_preflight = is_request_preflight(&req);
        let fut = self.service.call(req);

        Box::pin(async move {
            let mut res = fut.await?;
            if is_preflight {
                let mut rt = actix_web::HttpResponse::Ok();
                if let Some(o) = reqheaders.get("Access-Control-Request-Headers") {
                    rt.insert_header(("Access-Control-Allow-Headers", o.to_str().unwrap_or("")));
                };
                if let Some(o) = reqheaders.get("Access-Control-Request-Method") {
                    rt.insert_header(("Access-Control-Allow-Methods", o.to_str().unwrap_or("")));
                };
                if let Some(o) = reqheaders.get("Origin") {
                    rt.insert_header(("Access-Control-Allow-Origin", o.to_str().unwrap_or("")));
                };
                rt.insert_header((
                    "Access-Control-Expose-Headers",
                    "access-control-allow-origin, content-type",
                ));
                rt.insert_header(("Access-Control-Allow-Credentials", "true"));
                rt.insert_header(("WWW-Authenticate", "Basic realm=\"file\""));
                res = ServiceResponse::new(
                    res.request().to_owned(),
                    rt.message_body(res.into_body())?,
                );
                return Ok(res);
            } else if let Some(o) = reqheaders.get("Origin") {
                res.headers_mut().insert(
                    http::header::HeaderName::try_from("Access-Control-Allow-Origin").unwrap(),
                    o.to_owned(),
                );
            };
            Ok(res)
        })
    }
}

/// Try to parse header value as HTTP method.
fn header_value_try_into_method(hdr: &http::header::HeaderValue) -> Option<http::Method> {
    hdr.to_str()
        .ok()
        .and_then(|meth| http::Method::try_from(meth).ok())
}

fn is_request_preflight(req: &ServiceRequest) -> bool {
    // check request method is OPTIONS
    if req.method() != http::Method::OPTIONS {
        return false;
    }

    // check follow-up request method is present and valid
    if req
        .headers()
        .get(http::header::ACCESS_CONTROL_REQUEST_METHOD)
        .and_then(header_value_try_into_method)
        .is_none()
    {
        return false;
    }

    true
}
