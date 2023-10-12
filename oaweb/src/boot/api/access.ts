/*
 * access.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-12 19:32
 * Distributed under terms of the MIT license.
 */


import ajax from './axios'

export default (app_id: string) => {
  return {
    local: `./app/${app_id}/access/`,
    create(name: string, props?: { name?: string, level?: number, role_id?: string, user_id?: string, rid?: string }) {
      return ajax.post(this.local, Object.assign({ name, level: 0 }, props))
    },
    get(id: string) {
      return ajax.get(this.local + id)
    },
    list(props?: { name?: string, role_id?: string, user_id?: string }) {
      return ajax.get(this.local, props)
    },
    update(uuid: string, props: any) {
      return ajax.patch(this.local + uuid, props)
    },
    del(id: string) {
      return ajax.delete(this.local + id)
    },
  }
}
