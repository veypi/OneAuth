/*
 * user.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-05 15:37
 * Distributed under terms of the MIT license.
 */



import ajax from './axios'
import cfg from '../cfg'



export default {
  local: () => cfg.BaseUrl() + '/user/',
  search(q: string) {
    return ajax.get(this.local(), { username: q })
  },
  get(id: number) {
    return ajax.get(this.local() + id)
  },
}

