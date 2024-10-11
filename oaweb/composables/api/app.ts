//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 14:36:07
// Distributed under terms of the MIT license.
//

import webapi from "./webapi"
import * as models from "./models"
export interface GetOpts { 
  name: string
}
export function Get(app_id: string, json: GetOpts) {
  return webapi.Get<models.App>(`/app/${app_id}`, { json })
}

export interface PatchOpts { 
  name?: string
  icon?: string
  des?: string
  participate?: string
  init_role_id?: string
}
export function Patch(app_id: string, json: PatchOpts) {
  return webapi.Patch<models.App>(`/app/${app_id}`, { json })
}

export function Delete(app_id: string) {
  return webapi.Delete<models.App>(`/app/${app_id}`, {  })
}

export interface PostOpts { 
  name: string
  icon: string
  des: string
  participate: string
}
export function Post(json: PostOpts) {
  return webapi.Post<models.App>(`/app`, { json })
}

export interface ListOpts { 
  name?: string
}
export function List(json: ListOpts) {
  return webapi.Get<models.App>(`/app`, { json })
}

export interface AppUserGetOpts { 
  user_id: string
}
export function AppUserGet(app_user_id: string, app_id: string, json: AppUserGetOpts) {
  return webapi.Get<models.AppUser>(`/app/${app_id}/app_user/${app_user_id}`, { json })
}

export interface AppUserPatchOpts { 
  status?: string
}
export function AppUserPatch(app_user_id: string, app_id: string, json: AppUserPatchOpts) {
  return webapi.Patch<models.AppUser>(`/app/${app_id}/app_user/${app_user_id}`, { json })
}

export function AppUserDelete(app_user_id: string, app_id: string) {
  return webapi.Delete<models.AppUser>(`/app/${app_id}/app_user/${app_user_id}`, {  })
}

export interface AppUserListOpts { 
  user_id?: string
  status?: string
}
export function AppUserList(app_id: string, json: AppUserListOpts) {
  return webapi.Get<models.AppUser>(`/app/${app_id}/app_user`, { json })
}

export interface AppUserPostOpts { 
  status: string
  user_id: string
}
export function AppUserPost(app_id: string, json: AppUserPostOpts) {
  return webapi.Post<models.AppUser>(`/app/${app_id}/app_user`, { json })
}

export interface ResourceListQuery { 
  created_at?: Date
  updated_at?: Date
}
export function ResourceList(app_id: string, query: ResourceListQuery) {
  return webapi.Get<models.Resource>(`/app/${app_id}/resource`, { query })
}

export interface ResourcePostOpts { 
  name: string
  des: string
}
export function ResourcePost(app_id: string, json: ResourcePostOpts) {
  return webapi.Post<models.Resource>(`/app/${app_id}/resource`, { json })
}

export interface ResourceDeleteOpts { 
  name: string
}
export function ResourceDelete(app_id: string, json: ResourceDeleteOpts) {
  return webapi.Delete<models.Resource>(`/app/${app_id}/resource`, { json })
}

export function ResourceGet(app_id: string) {
  return webapi.Get<models.Resource>(`/app/${app_id}/resource`, {  })
}

export function ResourcePatch(app_id: string) {
  return webapi.Patch<models.Resource>(`/app/${app_id}/resource`, {  })
}

