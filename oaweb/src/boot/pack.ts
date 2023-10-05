/*
 * pack.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-26 20:38
 * Distributed under terms of the MIT license.
 */



import { boot } from 'quasar/wrappers'
import '@veypi/msg/index.css'
import '../assets/icon.js'
// import { Cfg } from '@veypi/oaer'
import { Cfg } from '@veypi/oaer'
// import { Cfg } from '/Users/veypi/test/oaer'

Cfg.host.value = 'http://' + window.location.host
Cfg.token.value = localStorage.getItem('auth_token') || ''
Cfg.uuid.value = 'FR9P5t8debxc11aFF'

// "async" is optional;
// more info on params: https://v2.quasar.dev/quasar-cli/boot-files
export default boot(async (/* { app, router, ... } */) => {
  // something to do
})
