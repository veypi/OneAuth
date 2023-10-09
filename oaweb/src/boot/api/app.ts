/*
 * app.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-30 17:31
 * Distributed under terms of the MIT license.
 */


import ajax from './axios'

export default {
  local: './app/',
  self() {
    return ajax.get(this.local, { option: 'oa' })
  },
  getKey(uuid: string) {
    return ajax.get(this.local + uuid, { option: 'key' })
  },
  create(name: string, icon: string) {
    return ajax.post(this.local, { name, icon })
  },
  get(uuid: string) {
    return ajax.get(this.local + uuid)
  },
  list() {
    return ajax.get(this.local)
  },
  update(uuid: string, props: any) {
    return ajax.patch(this.local + uuid, props)
  },
  user(uuid: string) {
    if (uuid === '') {
      uuid = '-'
    }
    return {
      local: this.local + uuid + '/user/',
      list(id: string, data?: any) {
        return ajax.get(this.local + id, data)
      },
      add(uid: string) {
        return ajax.post(this.local + uid)
      },
      update(uid: string, status: number) {
        return ajax.patch(this.local + uid, { status })
      },
    }
  },
}
