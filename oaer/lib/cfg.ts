/*
 * cfg.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-07-03 18:22
 * Distributed under terms of the MIT license.
 */

let cfg = {
  uuid: '',
  refresh: '',
  token: '',


  app_data: {},
  ready: false,
  local_user: {},

  host: '',
  _host_nats: '',
  nats_pub_key: '',
  prefix: '/api',
  BaseUrl() {
    return this.host.value + this.prefix
  },
  NatsHost() {
    if (this._host_nats.startsWith('ws')) {
      return this._host_nats
    }
    return 'ws://' + this._host_nats
  },
  media(u: string) {
    return this.host.value + u
  },
  goto(url: string) {
    if (url.startsWith('http')) {
      window.location.href = url
      return
    }
    if (!url.startsWith('/')) {
      url = '/' + url
    }
    window.location.href = this.host.value + url
  },
  Host() {
    return this.host.value || window.location.host
  },
  userFileUrl() {
    return '/fs/u/'
  },
  appFileUrl() {
    return `/fs/a/${this.uuid.value}/`
  },
}


export default cfg
