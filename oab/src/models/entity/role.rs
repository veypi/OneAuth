//! `SeaORM` Entity. Generated by sea-orm-codegen 0.12.3

use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(
    Clone, Debug, PartialEq, DeriveEntityModel, Eq, Serialize, Deserialize, Default, sqlx :: FromRow,
)]
#[sea_orm(table_name = "role")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created: DateTime,
    pub updated: DateTime,
    pub app_id: String,
    pub name: String,
    pub des: Option<String>,
    pub user_count: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::access::Entity")]
    Access,
    #[sea_orm(
        belongs_to = "super::app::Entity",
        from = "Column::AppId",
        to = "super::app::Column::Id",
        on_update = "NoAction",
        on_delete = "NoAction"
    )]
    App,
    #[sea_orm(has_many = "super::user_role::Entity")]
    UserRole,
}

impl Related<super::access::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Access.def()
    }
}

impl Related<super::app::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::App.def()
    }
}

impl Related<super::user_role::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UserRole.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        super::user_role::Relation::User.def()
    }
    fn via() -> Option<RelationDef> {
        Some(super::user_role::Relation::Role.def().rev())
    }
}

impl ActiveModelBehavior for ActiveModel {}
