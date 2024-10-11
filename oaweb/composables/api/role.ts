//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 00:18:06
// Distributed under terms of the MIT license.
//

import webapi from "./webapi"
import * as models from "./models"
export interface ListOpts { 
  name?: string
}
export function List(json: ListOpts) {
  return webapi.Get<models.Role>(`/role`, { json })
}

export function Get(role_id: string) {
  return webapi.Get<models.Role>(`/role/${role_id}`, {  })
}

export interface PatchOpts { 
  name?: string
  des?: string
  app_id?: string
}
export function Patch(role_id: string, json: PatchOpts) {
  return webapi.Patch<models.Role>(`/role/${role_id}`, { json })
}

export function Delete(role_id: string) {
  return webapi.Delete<models.Role>(`/role/${role_id}`, {  })
}

export interface PostOpts { 
  name: string
  des: string
  app_id: string
}
export function Post(json: PostOpts) {
  return webapi.Post<models.Role>(`/role`, { json })
}

