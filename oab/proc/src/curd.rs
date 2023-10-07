//
// curd.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-10-06 23:47
// Distributed under terms of the MIT license.
//

use proc_macro2::{Ident, Span};
use quote::{quote, ToTokens};
use syn::{AttributeArgs, ItemFn, NestedMeta, ReturnType};

pub struct CrudWrap {
    func: ItemFn,
    args: Crud,
    method: i32,
}

impl CrudWrap {
    pub fn new(args: AttributeArgs, func: ItemFn, method: i32) -> syn::Result<Self> {
        let args = Crud::new(args)?;

        Ok(Self { func, args, method })
    }
}

impl ToTokens for CrudWrap {
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
        let model_name = &self.args.model;

        let builder_fields = self.args.attrs.iter().map(|field| {
            quote! {
                if let Some(#field) = _data.#field {
                    obj.#field = sea_orm::Set(#field.into())
                };
            }
        });

        let stream = quote! {
            #(#fn_attrs)*
            #func_vis #fn_async fn #fn_name #fn_generics(
                #fn_args
            ) -> #fn_output {
                let _id = &id.clone();
                let _data = data.clone();
                let _db = &stat.db().clone();
                let f = || async move #func_block;
                let res = f().await;
                match res {
                    Err(e) => Err(e),
                    Ok(res) => {
                        let obj = crate::models::#model_name::Entity::find_by_id(_id).one(_db).await?;
                        let mut obj: crate::models::#model_name::ActiveModel = match obj {
                            Some(o) => o.into(),
                            None => return Err(Error::NotFound(_id.to_owned())),
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
}

struct Crud {
    model: syn::Ident,
    attrs: Vec<syn::Ident>,
}

impl Crud {
    fn new(args: AttributeArgs) -> syn::Result<Self> {
        let mut model: Option<syn::Ident> = None;
        let mut attrs: Vec<syn::Ident> = Vec::new();
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

                _ => {
                    return Err(syn::Error::new_spanned(arg, "Unknown attribute."));
                }
            }
        }
        match model {
            Some(model) => Ok(Self { model, attrs }),
            None => Err(syn::Error::new(
                Span::call_site(),
                "The #[crud(..)] macro requires one `model` name",
            )),
        }
    }
}
