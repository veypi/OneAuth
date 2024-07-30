/*
 * Copyright (C) 2019 light <veypi@light-laptop>
 *
 * Distributed under terms of the MIT license.
 */

import user from './user'
import app from './app'
import cfg from '../cfg'
import ajax from './axios'
import nats from './nats'
import tsdb from './tsdb'


const api = {
  nats: nats,
  tsdb: tsdb,
  user: user,
  app: app,
  info: () => {
    return ajax.get(cfg.BaseUrl() + '/info')
  },
  refresh_token: () => {
    ajax.post(cfg.BaseUrl() + '/app/' + cfg.uuid + '/token/', { app_id: cfg.uuid, token: cfg.refresh }).then(e => {
      cfg.token = e
      // bus.emit('sync', e)
    }).catch(e => {
      console.warn(e)
      // bus.emit('logout', 'get token failed ' + e)
    })
  }
}

export { api }
export default api
