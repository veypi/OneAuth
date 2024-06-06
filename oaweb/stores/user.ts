/*
 * user.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-06 20:17
 * Distributed under terms of the MIT license.
 */

import { type Auths } from '@/composables/models'
import { Base64 } from 'js-base64'

export const useUserStore = defineStore('user', {
  state: () => ({
    id: '',
    local: {} as modelsUser,
    auth: {} as Auths,
    ready: false
  }),
  getters: {
  },
  actions: {
    logout(msg?: string) {
      console.log('logout: ' + msg)
      this.ready = false
      util.setToken('')
      useRouter().push('/login')
    },
    fetchUserData() {
      let token = util.getToken().split('.');
      if (!token || token.length !== 3) {
        return false
      }
      let data = JSON.parse(Base64.decode(token[1]))
      if (data.id) {
        let l = 'access to'
        data.access.map((e: any) => l = l + `\n${e.name}.${e.level}`)
        console.log(l)
        this.auth = NewAuths(data.access)
        api.user.get(data.id).then((e: modelsUser) => {
          this.id = e.id
          this.local = e
          this.ready = true
        }).catch((e) => {
          this.logout()
        })
      }
    }
  },
});

