//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 00:18:06
// Distributed under terms of the MIT license.
//

import webapi from "./webapi"
import * as models from "./models"
export function Delete(token_id: String) {
  return webapi.Delete<models.Token>(`/token/${token_id}`, {})
}

export interface ListOpts {
  user_id: String
  app_id: String
}
export function List(json: ListOpts) {
  return webapi.Get<models.Token>(`/token`, { json })
}

export interface TokenSaltOpts {
  username: String
  typ?: String
}

// keep
export function TokenSalt(json: TokenSaltOpts) {
  return webapi.Post<String>(`/token/salt`, { json })
}

export interface PostOpts {
  user_id: String
  token?: String
  salt?: String
  code?: String
  app_id?: String
  expired_at?: Date
  over_perm?: String
  device?: String
}
export function Post(json: PostOpts) {
  return webapi.Post<models.Token>(`/token`, { json })
}

export function Get(token_id: String) {
  return webapi.Get<models.Token>(`/token/${token_id}`, {})
}

export interface PatchOpts {
  expired_at?: Date
  over_perm?: String
}
export function Patch(token_id: String, json: PatchOpts) {
  return webapi.Patch<models.Token>(`/token/${token_id}`, { json })
}

