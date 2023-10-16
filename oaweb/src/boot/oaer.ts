/*
 * oaer.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-16 21:20
 * Distributed under terms of the MIT license.
 */


// import '@veypi/oaer'
import oaer from '@veypi/oaer'
import bus from 'src/libs/bus'
import util from 'src/libs/util'

oaer.set({
  token: util.getToken(),
  host: 'http://' + window.location.host,
  uuid: 'FR9P5t8debxc11aFF',
})

bus.on('token', (t: any) => {
  oaer.set({ token: t })
})
