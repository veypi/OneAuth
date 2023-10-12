/*
 * role.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-12 15:40
 * Distributed under terms of the MIT license.
 */


import ajax from './axios'

export default (app_id: string) => {
  return {
    local: `./app/${app_id}/role/`,
    create(name: string, props?: { des?: string }) {
      return ajax.post(this.local, Object.assign({ name }, props))
    },
    get(id: string) {
      return ajax.get(this.local + id)
    },
    list() {
      return ajax.get(this.local)
    },
    update(uuid: string, props: any) {
      return ajax.patch(this.local + uuid, props)
    },
    del(id: string) {
      return ajax.delete(this.local + id)
    },
    add(id: string, uid: string) {
      return ajax.get(this.local + `${id}/user/${uid}`)
    },
    drop(id: string, uid: string) {
      return ajax.delete(this.local + `${id}/user/${uid}`)
    }
  }
}
