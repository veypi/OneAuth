//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 00:18:06
// Distributed under terms of the MIT license.
//

import webapi from "./webapi"
import * as models from "./models"
export interface ListOpts { 
  app_id: String
  user_id?: String
  role_id?: String
  name?: String
}
export interface ListQuery { 
  created_at?: Date
  updated_at?: Date
}
export function List(json: ListOpts, query: ListQuery) {
  return webapi.Get<models.Access>(`/access`, { json, query })
}

export interface PostOpts { 
  app_id: String
  user_id?: String
  role_id?: String
  name: String
  t_id: String
  level: Number
}
export function Post(json: PostOpts) {
  return webapi.Post<models.Access>(`/access`, { json })
}

