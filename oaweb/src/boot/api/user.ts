
/*
 * user.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-22 20:18
 * Distributed under terms of the MIT license.
 */


import { Base64 } from 'js-base64'
import ajax from './axios'

export default {
  local: './user/',
  register(username: string, password: string, prop?: any) {
    const data = Object.assign({
      username: username,
      password: Base64.encode(password),
    }, prop)
    return ajax.post(this.local, data)
  },
  login(username: string, password: string) {
    return ajax.head(this.local + username, {
      typ: 'username',
      password: Base64.encode(password),
    })
  },
  search(q: string) {
    return ajax.get(this.local, { username: q })
  },
  get(id: number) {
    return ajax.get(this.local + id)
  },
  list(props?: { name?: string, role_id?: string, app_id?: string }) {
    return ajax.get(this.local, props)
  },
  update(id: string, props: any) {
    props.test = { a: 1 }
    return ajax.patch(this.local + id, props)
  },
}

