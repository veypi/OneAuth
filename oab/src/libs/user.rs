//
// user.rs
// Copyright (C) 2023 veypi <i@veypi.com>
// 2023-10-02 21:00
// Distributed under terms of the MIT license.
//

use crate::{
    models::{self, app, app_user, user_role},
    Error, Result,
};
use sea_orm::{ActiveModelTrait, ConnectionTrait, DatabaseTransaction, EntityTrait};

// 尝试绑定应用
pub async fn connect_to_app(
    uid: String,
    aid: String,
    db: &DatabaseTransaction,
    app_obj: Option<app::Model>,
) -> Result<app_user::Model> {
    let app_obj = match app_obj {
        Some(o) => o,
        None => match app::Entity::find_by_id(&aid).one(db).await? {
            Some(o) => o,
            None => return Err(Error::NotFound(aid.clone())),
        },
    };
    let m = match app_obj.join_method.into() {
        models::AppJoin::Disabled => return Err(Error::AppDisabledRegister),
        models::AppJoin::Auto => models::AUStatus::OK,
        models::AppJoin::Applying => models::AUStatus::Applying,
    };
    let au = app_user::ActiveModel {
        app_id: sea_orm::ActiveValue::Set(aid.clone()),
        user_id: sea_orm::ActiveValue::Set(uid.clone()),
        status: sea_orm::ActiveValue::Set(m.clone() as i32),
        ..Default::default()
    };
    let au = au.insert(db).await?;
    if m == models::AUStatus::OK {
        after_connected_to_app(uid, app_obj, db).await?;
    }
    Ok(au)
}

// 成功绑定应用后操作
pub async fn after_connected_to_app(
    uid: String,
    obj: app::Model,
    db: &DatabaseTransaction,
) -> Result<()> {
    if obj.role_id.is_some() {
        user_role::ActiveModel {
            user_id: sea_orm::ActiveValue::Set(uid.clone()),
            role_id: sea_orm::ActiveValue::Set(obj.role_id.unwrap().clone()),
            ..Default::default()
        }
        .insert(db)
        .await?;
    };
    let sql = format!(
        "update app set user_count = user_count + 1 where id = '{}'",
        obj.id
    );
    db.execute(sea_orm::Statement::from_string(
        sea_orm::DatabaseBackend::MySql,
        sql,
    ))
    .await?;
    Ok(())
}
