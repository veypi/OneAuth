//
// access.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-09-16 00:13
// Distributed under terms of the Apache license.
//

use proc_macro2::{Ident, Span};
use quote::{quote, ToTokens};
use syn::{AttributeArgs, ItemFn, NestedMeta, ReturnType};

pub struct AccessWrap {
    cb_fn: Option<Ident>,
    func: ItemFn,
    access: Access,
}

impl AccessWrap {
    pub fn new(args: AttributeArgs, func: ItemFn, cb_fn: Option<&str>) -> syn::Result<Self> {
        let cb_fn: Option<Ident> = match cb_fn {
            Some(cb) => Some(syn::parse_str(cb)?),
            None => None,
        };
        let args = Access::new(args)?;

        Ok(Self {
            cb_fn,
            func,
            access: args,
        })
    }
}

impl ToTokens for AccessWrap {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let func_vis = &self.func.vis;
        let func_block = &self.func.block;

        let fn_sig = &self.func.sig;
        let fn_attrs = &self.func.attrs;
        let fn_name = &fn_sig.ident;
        let fn_generics = &fn_sig.generics;
        let fn_args = &fn_sig.inputs;
        let fn_async = &fn_sig.asyncness.unwrap();
        let fn_output = match &fn_sig.output {
            ReturnType::Type(ref _arrow, ref ty) => ty.to_token_stream(),
            ReturnType::Default => {
                quote! {()}
            }
        };

        let resp = quote!(Err(Error::NotAuthed));

        let permissions = &self.access.domain;
        // let args = quote! {
        //     #(#permissions,)*
        // };
        let stream = match &self.cb_fn {
            Some(cb_fn) => {
                let condition = match &self.access.did {
                    Some(did) => {
                        quote! {
                            let _auth_did = #did;
                            if _auth_token.#cb_fn(#permissions, _auth_did)
                        }
                    }
                    None => {
                        quote! {
                        if _auth_token.#cb_fn(#permissions, "")
                        }
                    }
                };
                quote! {
                    #(#fn_attrs)*
                    #func_vis #fn_async fn #fn_name #fn_generics(
                        _auth_token: Option<actix_web::web::ReqData<crate::models::Token>>,
                        #fn_args
                    ) -> #fn_output {
                        let _auth_token = match _auth_token {
                            Some(_auth_token) => _auth_token.into_inner(),
                            None => {
                                return #resp
                            }
                        };
                        #condition {
                            let f = || async move #func_block;
                            f().await
                        } else {
                            #resp
                        }
                    }
                }
            }
            None => {
                quote! {
                    #(#fn_attrs)*
                    #func_vis #fn_async fn #fn_name #fn_generics(
                        _auth_token: Option<actix_web::web::ReqData<crate::models::Token>>,
                        #fn_args
                    ) -> #fn_output {
                        if _auth_token.is_some() {
                            let f = || async move #func_block;
                            f().await
                        } else {
                            #resp
                        }
                    }
                }
            }
        };

        let _stream = tokens.extend(stream);
    }
}

struct Access {
    domain: syn::LitStr,
    did: Option<syn::Expr>,
}

impl Access {
    fn new(args: AttributeArgs) -> syn::Result<Self> {
        let mut domain: Option<syn::LitStr> = None;
        let mut did = None;
        for arg in args {
            match arg {
                NestedMeta::Lit(syn::Lit::Str(lit)) => {
                    domain = Some(lit);
                }
                NestedMeta::Meta(syn::Meta::NameValue(syn::MetaNameValue {
                    path,
                    lit: syn::Lit::Str(lit_str),
                    ..
                })) => {
                    if path.is_ident("id") {
                        let expr = lit_str.parse().unwrap();
                        did = Some(expr);
                    } else {
                        return Err(syn::Error::new_spanned(
                            path,
                            "Unknown identifier. Available: 'id'",
                        ));
                    }
                }
                _ => {
                    return Err(syn::Error::new_spanned(arg, "Unknown attribute."));
                }
            }
        }
        match domain {
            Some(domain) => Ok(Self { domain, did }),
            None => Err(syn::Error::new(
                Span::call_site(),
                "The #[access(..)] macro requires one `auth` argument",
            )),
        }
    }
}
