//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 14:36:07
// Distributed under terms of the MIT license.
//

import webapi from "./webapi"
import * as models from "./models"
export interface ListOpts {
  app_id: string
  user_id?: string
  role_id?: string
  name?: string
}
export interface ListQuery {
  created_at?: Date
  updated_at?: Date
}
export function List(json: ListOpts, query: ListQuery) {
  return webapi.Get<[models.Access]>(`/access`, { json, query })
}

export interface PostOpts {
  app_id: string
  user_id?: string
  role_id?: string
  name: string
  t_id: string
  level: number
}
export function Post(json: PostOpts) {
  return webapi.Post<models.Access>(`/access`, { json })
}

