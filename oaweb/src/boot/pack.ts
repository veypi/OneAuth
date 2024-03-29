/*
 * pack.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-26 20:38
 * Distributed under terms of the MIT license.
 */



import { boot } from 'quasar/wrappers'
import '@veypi/msg/index.css'
import { conf } from '@veypi/msg'
import '../assets/icon.js'
import '@veypi/oaer/dist/index.css'
import 'cherry-markdown/dist/cherry-markdown.css';

import oafs from 'src/libs/oafs'
import { Cfg } from '@veypi/oaer'
import util from 'src/libs/util.js'
import evt from 'src/libs/evt.js'


oafs.setCfg({ token: util.getToken(), app_id: 'FR9P5t8debxc11aFF' })
Cfg.token.value = util.getToken()

conf.timeout = 5000
Cfg.host.value = 'http://' + window.location.host
Cfg.uuid.value = 'FR9P5t8debxc11aFF'

evt.on('token', (t: any) => {
  oafs.setCfg({ token: t })
  Cfg.token.value = t
})

// "async" is optional;
// more info on params: https://v2.quasar.dev/quasar-cli/boot-files
export default boot(async (/* { app, router, ... } */) => {
  // something to do
})
