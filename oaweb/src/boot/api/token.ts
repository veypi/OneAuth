/*
 * token.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-30 17:37
 * Distributed under terms of the MIT license.
 */


import ajax from './axios'

export default (uuid: string) => {
  return {
    local: './app/' + uuid + '/token/',
    get() {
      return ajax.get(this.local)
    },
  }
}
