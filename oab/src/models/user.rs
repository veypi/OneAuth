//
// user.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-06-02 23:03
// Distributed under terms of the MIT license.
//
//

use actix_web::ResponseError;

use chrono::{prelude::*, Duration};
use jsonwebtoken::{decode, encode, DecodingKey, EncodingKey, Header, Validation};
use serde::{Deserialize, Serialize};
use tracing::info;

use crate::{Error, Result};
use aes_gcm::aead::{Aead, NewAead};
use aes_gcm::{Aes256Gcm, Key, Nonce};
use block_padding::{Padding, Pkcs7};

use generic_array::typenum::U32;
use generic_array::GenericArray;
use rand::distributions::Alphanumeric;
use rand::{thread_rng, Rng};

fn rand_str(l: usize) -> String {
    thread_rng()
        .sample_iter(&Alphanumeric)
        .take(l)
        .map(char::from)
        .collect()
}

// fn padding<u8, N: ArrayLength<u8>>(s: &str) -> GenericArray<u8, N> {
//     let msg = s.as_bytes();
//     let pos = msg.len();
//     let mut block: GenericArray<u8, U32> = [0xff; 32].into();
//     block[..pos].copy_from_slice(msg);
//     Pkcs7::pad(&mut block, pos);
//     assert_eq!(&block[..], b"test\x04\x04\x04\x04");
//     let res = Pkcs7::unpad(&block).unwrap();
//     assert_eq!(res, msg);
//     block
// }

#[derive(Debug, Serialize, Deserialize, sqlx::FromRow)]
pub struct User {
    pub id: String,
    pub created: Option<NaiveDateTime>,
    pub updated: Option<NaiveDateTime>,
    pub delete_flag: bool,

    pub username: String,
    pub nickname: Option<String>,
    pub email: Option<String>,
    pub phone: Option<String>,
    pub icon: Option<String>,
    pub real_code: Option<String>,
    pub check_code: Option<Vec<u8>>,
    pub status: i32,
    pub used: i32,
    pub space: i32,
}

impl User {
    pub fn token(&self) -> Token {
        let t = Token {
            iss: "oa".to_string(),
            aud: "".to_string(),
            exp: (Utc::now() + Duration::days(4)).timestamp(),
            iat: Utc::now().timestamp(),
            id: self.id.clone(),
        };
        t
    }
    pub fn check_pass(&self, p: &str) -> Result<()> {
        let p = p.as_bytes();
        let mut key_block: GenericArray<u8, U32> = [0xff; 32].into();
        key_block[..p.len()].copy_from_slice(p);
        Pkcs7::pad(&mut key_block, p.len());
        // key 32 Byte
        let key = Key::from_slice(&key_block);
        let cipher = Aes256Gcm::new(&key);

        // 12 Byte
        // 96-bits; unique per message
        let nonce = Nonce::from_slice(&self.id.as_bytes()[..12]);

        let plaintext = match cipher.decrypt(nonce, self.check_code.as_ref().unwrap().as_slice()) {
            Ok(p) => p,
            Err(_) => return Err(Error::ArgInvalid("password".to_string())),
        };
        let plaintext = std::str::from_utf8(&plaintext).unwrap();
        if plaintext.eq(self.real_code.as_ref().unwrap()) {
            Ok(())
        } else {
            Err(Error::ArgInvalid("password".to_string()))
        }
    }
    pub fn update_pass(&mut self, p: &str) -> Result<()> {
        if p.len() < 6 || p.len() > 32 {
            return Err(Error::ArgInvalid("password".to_string()));
        }
        let p = p.as_bytes();
        let mut key_block: GenericArray<u8, U32> = [0xff; 32].into();
        key_block[..p.len()].copy_from_slice(p);
        Pkcs7::pad(&mut key_block, p.len());
        // key 32 Byte
        let key = Key::from_slice(&key_block);
        let cipher = Aes256Gcm::new(&key);

        // 12 Byte
        // 96-bits; unique per message
        let nonce = Nonce::from_slice(&self.id.as_bytes()[..12]);

        let real = rand_str(32);

        let ciphertext = cipher.encrypt(nonce, real.as_bytes().as_ref())?;
        self.check_code = Some(ciphertext);
        self.real_code = Some(real);
        // let plaintext = cipher.decrypt(nonce, ciphertext.as_ref())?;
        // let plaintext = std::str::from_utf8(&plaintext).unwrap();
        // info!("123123{:?}\n{:?}\n", real, plaintext);

        Ok(())
    }
}

impl Default for User {
    fn default() -> Self {
        Self {
            id: uuid::Uuid::new_v4().to_string().replace("-", ""),
            created: None,
            updated: None,
            delete_flag: false,

            username: "".to_string(),
            nickname: None,
            email: None,
            phone: None,
            icon: None,
            check_code: None,
            real_code: None,
            status: 0,
            used: 0,
            space: 300,
        }
    }
}

impl actix_web::Responder for User {
    type Body = actix_web::body::BoxBody;

    fn respond_to(self, _req: &actix_web::HttpRequest) -> actix_web::HttpResponse<Self::Body> {
        match serde_json::to_string(&self) {
            Ok(body) => match actix_web::HttpResponse::Ok()
                .content_type(actix_web::http::header::ContentType::json())
                .message_body(body)
            {
                Ok(res) => res.map_into_boxed_body(),
                Err(err) => Error::from(err).error_response(),
            },
            Err(_err) => Error::SerdeError.error_response(),
        }
    }
}
#[derive(Debug, Serialize, Deserialize)]
pub struct Token {
    pub iss: String, // Optional. token 发行者
    pub aud: String, // Optional. token 使用者
    pub exp: i64,    // Required  失效时间
    pub iat: i64,    // Optional. 发布时间
    pub id: String,  // 用户id
}

impl Token {
    pub fn from(t: &str) -> Result<Self> {
        let token = decode::<Self>(
            t,
            &DecodingKey::from_secret("secret".as_ref()),
            &Validation::default(),
        )?;
        Ok(token.claims)
    }
    pub fn to_string(&self) -> Result<String> {
        let token = encode(
            &Header::default(),
            self,
            &EncodingKey::from_secret("secret".as_ref()),
        )?;
        Ok(token)
    }
}
