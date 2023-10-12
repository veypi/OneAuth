//
// lib.rs
// Copyright (C) 2022 veypi <i@veypi.com>
// 2022-09-16 00:07
// Distributed under terms of the Apache license.
//
//
use proc_macro::TokenStream;
use quote::{quote, ToTokens};
use syn::{parse_macro_input, AttributeArgs, ItemFn};
mod access;
mod crud;
use access::AccessWrap;
use crud::CrudWrap;

#[proc_macro_attribute]
pub fn have_access(args: TokenStream, input: TokenStream) -> TokenStream {
    check_permissions(None, args, input)
}

#[proc_macro_attribute]
pub fn access_read(args: TokenStream, input: TokenStream) -> TokenStream {
    check_permissions(Some("can_read"), args, input)
}

#[proc_macro_attribute]
pub fn access_create(args: TokenStream, input: TokenStream) -> TokenStream {
    check_permissions(Some("can_create"), args, input)
}

#[proc_macro_attribute]
pub fn access_update(args: TokenStream, input: TokenStream) -> TokenStream {
    check_permissions(Some("can_update"), args, input)
}

#[proc_macro_attribute]
pub fn access_delete(args: TokenStream, input: TokenStream) -> TokenStream {
    check_permissions(Some("can_delete"), args, input)
}

fn check_permissions(cb_fn: Option<&str>, args: TokenStream, input: TokenStream) -> TokenStream {
    let args = parse_macro_input!(args as AttributeArgs);
    let func = parse_macro_input!(input as ItemFn);

    match AccessWrap::new(args, func, cb_fn) {
        Ok(ac) => ac.into_token_stream().into(),
        Err(err) => err.to_compile_error().into(),
    }
}

#[proc_macro_attribute]
pub fn crud_update(args: TokenStream, input: TokenStream) -> TokenStream {
    derive_crud(3, args, input)
}

#[proc_macro_attribute]
pub fn crud_test(args: TokenStream, input: TokenStream) -> TokenStream {
    derive_crud(0, args, input)
}

fn derive_crud(method: i32, args: TokenStream, input: TokenStream) -> TokenStream {
    let args = parse_macro_input!(args as AttributeArgs);
    let func = parse_macro_input!(input as ItemFn);

    match CrudWrap::new(args, func, method) {
        Ok(ac) => ac.into_token_stream().into(),
        Err(err) => err.to_compile_error().into(),
    }
}

#[proc_macro_derive(MyDisplay)]
#[doc(hidden)]
pub fn display(input: TokenStream) -> TokenStream {
    // Parse the string representation
    let ast: syn::DeriveInput = syn::parse(input).unwrap();

    match ast.data {
        syn::Data::Enum(ref enum_data) => {
            let name = &ast.ident;
            impl_display(name, enum_data).into()
        }
        _ => panic!("#[derive(Display)] works only on enums"),
    }
}

fn impl_display(name: &syn::Ident, data: &syn::DataEnum) -> proc_macro2::TokenStream {
    let variants = data
        .variants
        .iter()
        .map(|variant| impl_display_for_variant(name, variant));

    quote! {
        impl ::std::fmt::Display for #name {
            fn fmt(&self, f: &mut ::std::fmt::Formatter) -> ::std::result::Result<(), ::std::fmt::Error> {
                match *self {
                    #(#variants)*
                }
            }
        }
    }
}

fn impl_display_for_variant(name: &syn::Ident, variant: &syn::Variant) -> proc_macro2::TokenStream {
    let id = &variant.ident;
    match variant.fields {
        syn::Fields::Unit => match &variant.discriminant {
            // print true value of enummember
            // enum {
            // a = 1
            //
            // }
            Some((_, value)) => {
                quote! {
                    #name::#id => {
                        f.write_str(stringify!(#value).to_lowercase().as_str())
                    }
                }
            }
            _ => {
                // print lowercase name of enummember
                quote! {
                    #name::#id => {
                        f.write_str(stringify!(#id).to_lowercase().as_str())
                    }
                }
            }
        },
        syn::Fields::Unnamed(ref fields) => match fields.unnamed.len() {
            0 => {
                quote! {
                    #name::#id() => {
                        f.write_str(stringify!(#id))?;
                        f.write_str("()")
                    }
                }
            }
            1 => {
                quote! {
                    #name::#id(ref inner) => {
                        ::std::fmt::Display::fmt(inner, f)
                    }
                }
            }
            _ => {
                panic!(
                    "#[derive(Display)] does not support tuple variants with more than one \
                            fields"
                )
            }
        },
        _ => panic!("#[derive(Display)] works only with unit and tuple variants"),
    }
}
