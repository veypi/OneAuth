//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 00:18:06
// Distributed under terms of the MIT license.
//

import webapi from "./webapi"
import * as models from "./models"
export interface PostOpts {
  username: String
  nickname?: String
  icon?: String
  email?: String
  phone?: String
  salt: String
  code: String
}
export function Post(json: PostOpts) {
  return webapi.Post<models.User>(`/user`, { json })
}

export function UserRoleGet(user_role_id: String, user_id: String) {
  return webapi.Get<models.UserRole>(`/user/${user_id}/user_role/${user_role_id}`, {})
}

export interface UserRolePatchOpts {
  status?: String
}
export function UserRolePatch(user_role_id: String, user_id: String, json: UserRolePatchOpts) {
  return webapi.Patch<models.UserRole>(`/user/${user_id}/user_role/${user_role_id}`, { json })
}

export interface UserRoleDeleteOpts {
  role_id: String
  app_id: String
}
export function UserRoleDelete(user_role_id: String, user_id: String, json: UserRoleDeleteOpts) {
  return webapi.Delete<models.UserRole>(`/user/${user_id}/user_role/${user_role_id}`, { json })
}

export interface UserRolePostOpts {
  status: String
  role_id: String
  app_id: String
}
export function UserRolePost(user_id: String, json: UserRolePostOpts) {
  return webapi.Post<models.UserRole>(`/user/${user_id}/user_role`, { json })
}

export interface UserLoginOpts {
  pwd: String
  typ: String
}

export function Get(user_id: String) {
  return webapi.Get<models.User>(`/user/${user_id}`, {})
}

export function Delete(user_id: String) {
  return webapi.Delete<models.User>(`/user/${user_id}`, {})
}

export interface UserRoleListOpts {
  status?: String
}
export function UserRoleList(user_id: String, json: UserRoleListOpts) {
  return webapi.Get<models.UserRole>(`/user/${user_id}/user_role`, { json })
}

export interface PatchOpts {
  username?: String
  nickname?: String
  icon?: String
  email?: String
  phone?: String
  status?: Number
}
export function Patch(user_id: String, json: PatchOpts) {
  return webapi.Patch<models.User>(`/user/${user_id}`, { json })
}

export interface ListOpts {
  username?: String
  nickname?: String
  email?: String
  phone?: String
  status?: Number
}
export function List(json: ListOpts) {
  return webapi.Get<models.User>(`/user`, { json })
}

