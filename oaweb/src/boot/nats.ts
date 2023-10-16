/*
 * nats.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-16 16:18
 * Distributed under terms of the MIT license.
 */

import axios from 'axios'
import util from '../libs/util'
import { connect, StringCodec } from '../libs/nats.ws'




const sc = StringCodec();
const nc = await connect({
  servers: 'ws://127.0.0.1:4221',
  authenticator: function(nonce?: string) {
    let nkey = 'UCXFAAVMCPTATZUZX6H24YF6FI3NKPQBPLM6BNN2EDFPNSUUEZPNFKEL'
    // let nre = nkeyAuthenticator(nkey_seed)
    let res = {
      nkey: nkey,
      sig: async function() {
        let response = await axios.post('/api/app/nats/token/', { token: util.getToken(), nonce: nonce });
        console.log(response)
        return response.data
      }
    };
    return res
  } as any
})

nc.publish('msg', '123')
const sub = nc.subscribe("msg");
(async () => {
  for await (const m of sub) {
    console.log(`[${sub.getProcessed()}]: ${sc.decode(m.data)}`);
  }
  console.log("subscription closed");
})();
