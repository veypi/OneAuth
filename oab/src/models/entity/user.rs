//! `SeaORM` Entity. Generated by sea-orm-codegen 0.12.3

use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(
    Clone, Debug, PartialEq, DeriveEntityModel, Eq, Serialize, Deserialize, Default, sqlx :: FromRow,
)]
#[sea_orm(table_name = "user")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created: Option<DateTime>,
    pub updated: Option<DateTime>,
    pub delete_flag: i8,
    #[sea_orm(unique)]
    pub username: String,
    pub nickname: Option<String>,
    #[sea_orm(unique)]
    pub email: Option<String>,
    #[sea_orm(unique)]
    pub phone: Option<String>,
    pub icon: Option<String>,
    #[sea_orm(column_name = "_real_code")]
    #[serde(skip)]
    pub real_code: Option<String>,
    #[sea_orm(
        column_name = "_check_code",
        column_type = "Binary(BlobSize::Blob(Some(48)))",
        nullable
    )]
    #[serde(skip)]
    pub check_code: Option<Vec<u8>>,
    pub status: i32,
    pub used: i32,
    pub space: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::access::Entity")]
    Access,
    #[sea_orm(has_many = "super::app_user::Entity")]
    AppUser,
    #[sea_orm(has_many = "super::user_role::Entity")]
    UserRole,
}

impl Related<super::access::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Access.def()
    }
}

impl Related<super::app_user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AppUser.def()
    }
}

impl Related<super::user_role::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UserRole.def()
    }
}

impl Related<super::app::Entity> for Entity {
    fn to() -> RelationDef {
        super::app_user::Relation::App.def()
    }
    fn via() -> Option<RelationDef> {
        Some(super::app_user::Relation::User.def().rev())
    }
}

impl Related<super::role::Entity> for Entity {
    fn to() -> RelationDef {
        super::user_role::Relation::Role.def()
    }
    fn via() -> Option<RelationDef> {
        Some(super::user_role::Relation::User.def().rev())
    }
}

impl ActiveModelBehavior for ActiveModel {}
