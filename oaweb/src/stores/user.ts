/*
 * user.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-22 21:05
 * Distributed under terms of the MIT license.
 */


import { defineStore } from 'pinia';
import { Auths, modelsUser, NewAuths } from 'src/models';
import { Base64 } from 'js-base64'
import router from 'src/router';
import api from 'src/boot/api';
import util from 'src/libs/util';

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
    logout() {
      this.ready = false
      util.setToken('')
      router.push({ name: 'login' })
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
