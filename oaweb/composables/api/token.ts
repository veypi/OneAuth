//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 14:36:07
// Distributed under terms of the MIT license.
//

import webapi from "./webapi"
import * as models from "./models"
export interface TokenSaltOpts {
  username: string
  typ?: string
}
// keep
export function TokenSalt(json: TokenSaltOpts) {
  return webapi.Post<{ id: string, salt: string }>(`/token/salt`, { json })
}
export interface PostOpts {
  user_id: string
  token?: string
  salt?: string
  code?: string
  app_id?: string
  expired_at?: Date
  over_perm?: string
  device?: string
}

// keep
export function Post(json: PostOpts) {
  return webapi.Post<string>(`/token`, { json })
}
export function Get(token_id: string) {
  return webapi.Get<models.Token>(`/token/${token_id}`, {})
}

export interface PatchOpts {
  expired_at?: Date
  over_perm?: string
}
export function Patch(token_id: string, json: PatchOpts) {
  return webapi.Patch<models.Token>(`/token/${token_id}`, { json })
}

export function Delete(token_id: string) {
  return webapi.Delete<models.Token>(`/token/${token_id}`, {})
}

export interface ListOpts {
  user_id: string
  app_id: string
}
export function List(json: ListOpts) {
  return webapi.Get<models.Token>(`/token`, { json })
}

