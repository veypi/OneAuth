//
// token.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-10-13 02:29
// Distributed under terms of the MIT license.
//

use actix_web::{post, web, Responder};
use nkeys;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter};
use serde::{Deserialize, Serialize};
use tracing::info;

use crate::{
    models::{self, AUStatus, AccessCore, UserPlugin},
    AppState, Error, Result,
};

#[derive(Debug, Deserialize, Serialize)]
pub struct GetOptions {
    app_id: Option<String>,
    token: String,
    nonce: Option<String>,
}

// 转换token
#[post("/app/{aid}/token/")]
pub async fn get(
    aid: web::Path<String>,
    stat: web::Data<AppState>,
    query: web::Json<GetOptions>,
) -> Result<impl Responder> {
    let aid = aid.into_inner();
    let mut key = stat.key.clone();
    info!("{}", key);
    let sid = match &query.app_id {
        Some(i) => {
            if !i.is_empty() {
                match models::app::Entity::find_by_id(i).one(stat.db()).await? {
                    Some(sapp) => key = sapp.key,
                    None => {}
                }
                info!("{}", key);
            };
            i
        }
        _ => "",
    };
    let token = models::Token::from(&query.token, &key)?;
    if aid.starts_with("nats") {
        let nonce = &query.nonce.clone().unwrap();
        let u = nkeys::KeyPair::from_seed(&stat.nats_usr[1].clone()).unwrap();
        let res = base64::encode(u.sign(nonce.as_bytes()).unwrap());
        return Ok(format!("{}@{}", res, &stat.nats_usr[0].clone()));
    };
    if !aid.is_empty() {
        // 从OA token 转向其他app token
        if sid.is_empty() {
            let s = models::app_user::Entity::find()
                .filter(models::app_user::Column::AppId.eq(&aid))
                .filter(models::app_user::Column::UserId.eq(&token.id))
                .one(stat.db())
                .await?;
            if s.is_none() {
                return Err(Error::NotAuthed);
            };
            let s = s.unwrap();
            if s.status == AUStatus::OK as i32 {
                let result = sqlx::query_as::<_, models::AccessCore>(
            "select access.name, access.rid, access.level from access, user_role, role WHERE user_role.user_id = ? && access.role_id=user_role.role_id && role.id=user_role.role_id && role.app_id = ?",
            )
            .bind(&token.id)
            .bind(&aid)
            .fetch_all(stat.sqlx())
            .await?;
                let appobj = models::app::Entity::find_by_id(&aid)
                    .one(stat.db())
                    .await?
                    .unwrap();
                let u = models::user::Entity::find_by_id(&token.id)
                    .one(stat.db())
                    .await?
                    .unwrap();
                let str = u.token(appobj.id.clone(), result).to_string(&appobj.key)?;
                Ok(str)
            } else {
                Err(Error::NotAuthed)
            }
        } else {
            let u = models::user::Entity::find_by_id(&token.id)
                .one(stat.db())
                .await?
                .unwrap();
            let str = u
                .token(
                    stat.uuid.clone(),
                    vec![
                        AccessCore {
                            name: "app".to_string(),
                            rid: None,
                            level: models::AccessLevel::Read,
                        },
                        AccessCore {
                            name: "user".to_string(),
                            rid: None,
                            level: models::AccessLevel::Read,
                        },
                    ],
                )
                .to_string(&stat.key)?;
            Ok(str)
        }
    } else {
        Err(Error::Missing("id".to_string()))
    }
}
