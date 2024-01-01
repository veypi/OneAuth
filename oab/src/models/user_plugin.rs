//
// user.rs
// Copyright (C) 2022 veypi <veypi@qq.com>
// 2022-06-02 23:03
// Distributed under terms of the MIT license.
//
//

use chrono::{prelude::*, Duration};
use jsonwebtoken::{decode, encode, DecodingKey, EncodingKey, Header, Validation};
use serde::{Deserialize, Serialize};

use crate::{Error, Result};
use aes_gcm::aead::{Aead, NewAead};
use aes_gcm::{Aes256Gcm, Key, Nonce};
use block_padding::{Padding, Pkcs7};

use generic_array::typenum::U32;
use generic_array::GenericArray;
use proc::MyDisplay;
use rand::distributions::Alphanumeric;
use rand::{thread_rng, Rng};
use serde_repr::*;

pub fn rand_str(l: usize) -> String {
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
//

#[derive(Debug, Serialize, Deserialize, Clone, sqlx::FromRow)]
pub struct AccessCore {
    pub name: String,
    pub rid: Option<String>,
    pub level: AccessLevel,
}

pub trait UserPlugin {
    fn token(&self, aid: String, ac: Vec<AccessCore>) -> Token;
    fn check_pass(&self, p: &str) -> Result<()>;
    fn update_pass(&mut self, p: &str) -> Result<()>;
}

// impl User {
impl UserPlugin for super::entity::user::Model {
    fn token(&self, aid: String, ac: Vec<AccessCore>) -> Token {
        let default_ico = "/media/".to_string();
        let t = Token {
            iss: "oa".to_string(),
            aud: "".to_string(),
            exp: (Utc::now() + Duration::days(4)).timestamp(),
            iat: Utc::now().timestamp(),
            id: self.id.clone(),
            aid: aid,
            icon: self.icon.as_ref().unwrap_or(&default_ico).to_string(),
            access: Some(ac),
            nickname: self
                .nickname
                .as_ref()
                .unwrap_or(&self.username.clone())
                .to_string(),
        };
        t
    }
    fn check_pass(&self, p: &str) -> Result<()> {
        let p = p.as_bytes();
        let mut key_block: GenericArray<u8, U32> = [0xff; 32].into();
        key_block[..p.len()].copy_from_slice(p);
        Pkcs7::pad(&mut key_block, p.len());
        // key 32 Byte
        let key = Key::from_slice(&key_block.as_slice());
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
    fn update_pass(&mut self, p: &str) -> Result<()> {
        if p.len() < 6 || p.len() > 32 {
            return Err(Error::ArgInvalid("password".to_string()));
        }
        let p = p.as_bytes();
        let mut key_block: GenericArray<u8, U32> = [0xff; 32].into();
        key_block[..p.len()].copy_from_slice(p);
        Pkcs7::pad(&mut key_block, p.len());
        // key 32 Byte
        let key = Key::from_slice(&key_block.as_slice());
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

#[derive(
    MyDisplay,
    Debug,
    Deserialize_repr,
    Serialize_repr,
    Clone,
    sqlx::Type,
    PartialEq,
    Eq,
    PartialOrd,
    Ord,
)]
#[repr(i64)]
pub enum AccessLevel {
    No = 0,
    Read = 1,
    Create = 2,
    Update = 3,
    Delete = 4,
    ALL = 5,
}

impl From<i32> for AccessLevel {
    fn from(v: i32) -> Self {
        match v {
            x if x == AccessLevel::No as i32 => AccessLevel::No,
            x if x == AccessLevel::Read as i32 => AccessLevel::Read,
            x if x == AccessLevel::Create as i32 => AccessLevel::Create,
            x if x == AccessLevel::Update as i32 => AccessLevel::Update,
            x if x == AccessLevel::Delete as i32 => AccessLevel::Delete,
            x if x == AccessLevel::ALL as i32 => AccessLevel::ALL,
            _ => AccessLevel::No,
        }
    }
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Token {
    pub iss: String, // Optional. token 发行者
    pub aud: String, // Optional. token 使用者
    pub exp: i64,    // Required  失效时间
    pub iat: i64,    // Optional. 发布时间
    pub id: String,  // 用户id
    pub nickname: String,
    pub icon: String,
    pub aid: String,
    pub access: Option<Vec<AccessCore>>,
}

impl Token {
    pub fn from(t: &str, key: &str) -> Result<Self> {
        let token = decode::<Self>(
            t,
            &DecodingKey::from_secret(key.as_ref()),
            &Validation::default(),
        )?;
        if token.claims.is_valid() {
            Ok(token.claims)
        } else {
            Err(Error::ExpiredToken)
        }
    }
    pub fn is_valid(&self) -> bool {
        if self.exp > Utc::now().timestamp() {
            true
        } else {
            false
        }
    }
    pub fn to_string(&self, key: &str) -> Result<String> {
        let token = encode(
            &Header::default(),
            self,
            &EncodingKey::from_secret(key.as_ref()),
        )?;
        Ok(token)
    }
    // pub fn to_string(&self) -> Result<String> {
    //     let token = encode(
    //         &Header::default(),
    //         self,
    //         &EncodingKey::from_secret(self._key.as_ref()),
    //     )?;
    //     Ok(token)
    // }

    fn check(&self, domain: &str, did: &str, l: AccessLevel) -> bool {
        match &self.access {
            Some(ac) => {
                for ele in ac {
                    if ele.name == domain && ele.level >= l {
                        match &ele.rid {
                            Some(temp) => {
                                if temp == did {
                                    return true;
                                }
                            }
                            None => return true,
                        }
                    }
                }
                false
            }
            None => false,
        }
    }
    pub fn can_read(&self, domain: &str, did: &str) -> bool {
        self.check(domain, did, AccessLevel::Read)
    }

    pub fn can_create(&self, domain: &str, did: &str) -> bool {
        self.check(domain, did, AccessLevel::Create)
    }

    pub fn can_update(&self, domain: &str, did: &str) -> bool {
        self.check(domain, did, AccessLevel::Update)
    }

    pub fn can_delete(&self, domain: &str, did: &str) -> bool {
        self.check(domain, did, AccessLevel::Delete)
    }
}
