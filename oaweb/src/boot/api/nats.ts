/*
 * nats.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-19 21:36
 * Distributed under terms of the MIT license.
 */


import ajax from './axios'

export default {
  local: './nats/',
  general() {
    return ajax.get(this.local + 'varz')
  },
  conns() {
    return ajax.get(this.local + 'connz', { subs: true })
  },
}
