//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 00:18:06
// Distributed under terms of the MIT license.
//

import webapi from "./webapi"
import * as models from "./models"
export interface ListOpts { 
  name?: String
}
export function List(json: ListOpts) {
  return webapi.Get<models.Role>(`/role`, { json })
}

export function Get(role_id: String) {
  return webapi.Get<models.Role>(`/role/${role_id}`, {  })
}

export interface PatchOpts { 
  name?: String
  des?: String
  app_id?: String
}
export function Patch(role_id: String, json: PatchOpts) {
  return webapi.Patch<models.Role>(`/role/${role_id}`, { json })
}

export function Delete(role_id: String) {
  return webapi.Delete<models.Role>(`/role/${role_id}`, {  })
}

export interface PostOpts { 
  name: String
  des: String
  app_id: String
}
export function Post(json: PostOpts) {
  return webapi.Post<models.Role>(`/role`, { json })
}

