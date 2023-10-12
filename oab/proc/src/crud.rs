//
// curd.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-10-06 23:47
// Distributed under terms of the MIT license.
//

use proc_macro2::Span;
use quote::{format_ident, quote, ToTokens};
use syn::{AttributeArgs, ItemFn, NestedMeta, ReturnType};

pub struct CrudWrap {
    func: ItemFn,
    args: Crud,
    method: i32,
}

impl CrudWrap {
    pub fn new(args: AttributeArgs, func: ItemFn, method: i32) -> syn::Result<Self> {
        let args = Crud::new(args, method)?;

        Ok(Self { func, args, method })
    }
    fn update(&self, tokens: &mut proc_macro2::TokenStream) {
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
        let model_name = &self.args.model;

        let builder_fields = self.args.props.iter().map(|field| {
            quote! {
                if let Some(#field) = _data.#field {
                    obj.#field = sea_orm::Set(#field.into())
                };
            }
        });
        let (args_fields, filter_fields) = match self.args.filters.len() {
            1 => {
                let k = &self.args.filters[0];
                let _k = format_ident!("_{}", k);
                (
                    vec![quote! {let #_k = _id; }],
                    vec![quote! {
                    filter(crate::models::#model_name::Column::#k.eq(#_k))
                        }],
                )
            }
            _ => (
                self.args
                    .filters
                    .iter()
                    .enumerate()
                    .map(|(idx, k)| {
                        let _k = format_ident!("_{}", k);
                        quote! {
                            let #_k = &_id[#idx];
                        }
                    })
                    .collect(),
                self.args
                    .filters
                    .iter()
                    .map(|k| {
                        let _k = format_ident!("_{}", k);
                        quote! {
                        filter(crate::models::#model_name::Column::#k.eq(#_k))
                            }
                    })
                    .collect(),
            ),
        };
        let stream = quote! {
            #(#fn_attrs)*
            #func_vis #fn_async fn #fn_name #fn_generics(
                #fn_args
            ) -> #fn_output {
                let _id = id.clone();
                let _data = data.clone();
                let _db = &stat.db().clone();
                #(#args_fields)*
                let f = || async move #func_block;
                let res = f().await;
                match res {
                    Err(e) => Err(e),
                    Ok(res) => {
                        let obj = crate::models::#model_name::Entity::find().#(#filter_fields).*.one(_db).await?;
                        let mut obj: crate::models::#model_name::ActiveModel = match obj {
                            Some(o) => o.into(),
                            None => return Err(Error::NotFound("".into())),
                        };
                        #(#builder_fields)*
                        let obj = obj.update(_db).await?;
                        Ok(actix_web::web::Json(obj))
                    }
                }
            }
        };

        let _stream = tokens.extend(stream);
    }

    fn copy(&self, _: &mut proc_macro2::TokenStream) {
        let _ = self.args.attrs.len();
    }
}

impl ToTokens for CrudWrap {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        match self.method {
            3 => self.update(tokens),
            // 3 => self.update(tokens),
            _ => self.copy(tokens),
        }
    }
}

struct Crud {
    model: syn::Ident,
    attrs: Vec<syn::Ident>,
    filters: Vec<syn::Ident>,
    props: Vec<syn::Ident>,
}

impl Crud {
    fn new(args: AttributeArgs, method: i32) -> syn::Result<Self> {
        let mut model: Option<syn::Ident> = None;
        let mut attrs: Vec<syn::Ident> = Vec::new();
        let mut filters: Vec<syn::Ident> = Vec::new();
        let mut props: Vec<syn::Ident> = Vec::new();
        if method == 0 {
            return Ok(Self {
                model: format_ident!("a"),
                attrs,
                filters,
                props,
            });
        }
        for arg in args {
            match arg {
                // NestedMeta::Lit(syn::Lit::Str(lit)) => {
                // }
                NestedMeta::Meta(syn::Meta::Path(i)) => match i.get_ident() {
                    Some(i) => {
                        if None == model {
                            model = Some(i.to_owned());
                        } else {
                            attrs.push(i.to_owned());
                        }
                    }
                    None => {}
                },
                NestedMeta::Meta(syn::Meta::NameValue(syn::MetaNameValue {
                    path,
                    lit: syn::Lit::Str(lit_str),
                    ..
                })) => {
                    if path.is_ident("filter") {
                        filters = lit_str
                            .value()
                            .replace(" ", "")
                            .split(",")
                            .into_iter()
                            .map(|l| {
                                return format_ident!("{}", l);
                            })
                            .collect();
                    } else if path.is_ident("props") {
                        props = lit_str
                            .value()
                            .replace(" ", "")
                            .split(",")
                            .into_iter()
                            .map(|l| {
                                return format_ident!("{}", l);
                            })
                            .collect();
                    } else {
                        return Err(syn::Error::new_spanned(
                            path,
                            "Unknown identifier. Available: filter, props ",
                        ));
                    }
                }

                _ => {
                    return Err(syn::Error::new_spanned(arg, "Unknown attribute."));
                }
            }
        }
        match model {
            Some(model) => Ok(Self {
                model,
                attrs,
                filters,
                props,
            }),
            None => Err(syn::Error::new(
                Span::call_site(),
                "The #[crud(..)] macro requires one `model` name",
            )),
        }
    }
}
