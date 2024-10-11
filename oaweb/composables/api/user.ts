//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 00:18:06
// Distributed under terms of the MIT license.
//

import webapi from "./webapi"
import * as models from "./models"
export interface PostOpts {
  username: string
  nickname?: string
  icon?: string
  email?: string
  phone?: string
  salt: string
  code: string
}
export function Post(json: PostOpts) {
  return webapi.Post<models.User>(`/user`, { json })
}

export function UserRoleGet(user_role_id: string, user_id: string) {
  return webapi.Get<models.UserRole>(`/user/${user_id}/user_role/${user_role_id}`, {  })
}

export interface UserRolePatchOpts {
  status?: string
}
export function UserRolePatch(user_role_id: string, user_id: string, json: UserRolePatchOpts) {
  return webapi.Patch<models.UserRole>(`/user/${user_id}/user_role/${user_role_id}`, { json })
}

export interface UserRoleDeleteOpts {
  role_id: string
  app_id: string
}
export function UserRoleDelete(user_role_id: string, user_id: string, json: UserRoleDeleteOpts) {
  return webapi.Delete<models.UserRole>(`/user/${user_id}/user_role/${user_role_id}`, { json })
}

export interface UserRolePostOpts {
  status: string
  role_id: string
  app_id: string
}
export function UserRolePost(user_id: string, json: UserRolePostOpts) {
  return webapi.Post<models.UserRole>(`/user/${user_id}/user_role`, { json })
}

export interface UserLoginOpts {
  pwd: String
  typ: String
}

export function Get(user_id: string) {
  return webapi.Get<models.User>(`/user/${user_id}`, {  })
}

export function Delete(user_id: string) {
  return webapi.Delete<models.User>(`/user/${user_id}`, {  })
}

export interface UserRoleListOpts {
  status?: string
}
export function UserRoleList(user_id: string, json: UserRoleListOpts) {
  return webapi.Get<models.UserRole>(`/user/${user_id}/user_role`, { json })
}

export interface PatchOpts {
  username?: string
  nickname?: string
  icon?: string
  email?: string
  phone?: string
  status?: number
}
export function Patch(user_id: string, json: PatchOpts) {
  return webapi.Patch<models.User>(`/user/${user_id}`, { json })
}

export interface ListOpts {
  username?: string
  nickname?: string
  email?: string
  phone?: string
  status?: number
}
export function List(json: ListOpts) {
  return webapi.Get<models.User>(`/user`, { json })
}

