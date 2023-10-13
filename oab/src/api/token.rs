//
// token.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-10-13 02:29
// Distributed under terms of the MIT license.
//

use actix_web::{get, web, Responder};
use proc::access_read;
use sea_orm::{ColumnTrait, EntityTrait, QueryFilter};

use crate::{
    models::{self, AUStatus, Token, UserPlugin},
    AppState, Error, Result,
};

#[get("/app/{aid}/token/")]
#[access_read("app")]
pub async fn get(
    aid: web::Path<String>,
    stat: web::Data<AppState>,
    t: web::ReqData<Token>,
) -> Result<impl Responder> {
    let n = aid.into_inner();
    if !n.is_empty() {
        let s = models::app_user::Entity::find()
            .filter(models::app_user::Column::AppId.eq(&n))
            .filter(models::app_user::Column::UserId.eq(&t.id))
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
            .bind(&t.id)
            .bind(&n)
            .fetch_all(stat.sqlx())
            .await?;
            let appobj = models::app::Entity::find_by_id(&n)
                .one(stat.db())
                .await?
                .unwrap();
            let u = models::user::Entity::find_by_id(&t.id)
                .one(stat.db())
                .await?
                .unwrap();
            let str = u.token(result).to_string(&appobj.key)?;
            // tokio::spawn(async move {
            //     let mut interval = tokio::time::interval(Duration::from_secs(5));
            //     interval.tick().await;
            //     let start = Instant::now();
            //     println!("time:{:?}", start);
            //     loop {
            //         interval.tick().await;
            //         println!("time:{:?}", start.elapsed());
            //     }
            // });

            Ok(str)
        } else {
            Err(Error::NotAuthed)
        }
    } else {
        Err(Error::Missing("id".to_string()))
    }
}
