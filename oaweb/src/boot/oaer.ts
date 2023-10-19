/*
 * oaer.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-16 21:20
 * Distributed under terms of the MIT license.
 */


// import '@veypi/oaer'
import oaer from '@veypi/oaer'
import '@veypi/oaer/dist/index.css'
import cfg from 'src/cfg'
import bus from 'src/libs/bus'
import util from 'src/libs/util'

oaer.set({
  token: util.getToken(),
  host: cfg.host,
  uuid: cfg.id,
})

bus.on('token', (t: any) => {
  oaer.set({ token: t })
})
