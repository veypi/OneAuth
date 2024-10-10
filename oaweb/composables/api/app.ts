//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 00:18:06
// Distributed under terms of the MIT license.
//

import webapi from "./webapi"
import * as models from "./models"
export interface GetOpts { 
  name: String
}
export function Get(app_id: String, json: GetOpts) {
  return webapi.Get<models.App>(`/app/${app_id}`, { json })
}

export function AppUserDelete(app_user_id: String, app_id: String) {
  return webapi.Delete<models.AppUser>(`/app/${app_id}/app_user/${app_user_id}`, {  })
}

export function ResourcePatch(app_id: String) {
  return webapi.Patch<models.Resource>(`/app/${app_id}/resource`, {  })
}

export interface PatchOpts { 
  name?: String
  icon?: String
  des?: String
  participate?: String
  init_role_id?: String
}
export function Patch(app_id: String, json: PatchOpts) {
  return webapi.Patch<models.App>(`/app/${app_id}`, { json })
}

export interface PostOpts { 
  name: String
  icon: String
  des: String
  participate: String
}
export function Post(json: PostOpts) {
  return webapi.Post<models.App>(`/app`, { json })
}

export interface AppUserListOpts { 
  user_id?: String
  status?: String
}
export function AppUserList(app_id: String, json: AppUserListOpts) {
  return webapi.Get<models.AppUser>(`/app/${app_id}/app_user`, { json })
}

export interface ResourceListQuery { 
  created_at?: Date
  updated_at?: Date
}
export function ResourceList(app_id: String, query: ResourceListQuery) {
  return webapi.Get<models.Resource>(`/app/${app_id}/resource`, { query })
}

export interface ResourcePostOpts { 
  name: String
  des: String
}
export function ResourcePost(app_id: String, json: ResourcePostOpts) {
  return webapi.Post<models.Resource>(`/app/${app_id}/resource`, { json })
}

export function Delete(app_id: String) {
  return webapi.Delete<models.App>(`/app/${app_id}`, {  })
}

export interface ListOpts { 
  name?: String
}
export function List(json: ListOpts) {
  return webapi.Get<models.App>(`/app`, { json })
}

export interface AppUserGetOpts { 
  user_id: String
}
export function AppUserGet(app_user_id: String, app_id: String, json: AppUserGetOpts) {
  return webapi.Get<models.AppUser>(`/app/${app_id}/app_user/${app_user_id}`, { json })
}

export function ResourceGet(app_id: String) {
  return webapi.Get<models.Resource>(`/app/${app_id}/resource`, {  })
}

export interface AppUserPatchOpts { 
  status?: String
}
export function AppUserPatch(app_user_id: String, app_id: String, json: AppUserPatchOpts) {
  return webapi.Patch<models.AppUser>(`/app/${app_id}/app_user/${app_user_id}`, { json })
}

export interface AppUserPostOpts { 
  status: String
  user_id: String
}
export function AppUserPost(app_id: String, json: AppUserPostOpts) {
  return webapi.Post<models.AppUser>(`/app/${app_id}/app_user`, { json })
}

export interface ResourceDeleteOpts { 
  name: String
}
export function ResourceDelete(app_id: String, json: ResourceDeleteOpts) {
  return webapi.Delete<models.Resource>(`/app/${app_id}/resource`, { json })
}

