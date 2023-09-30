/*
 * user.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-22 21:05
 * Distributed under terms of the MIT license.
 */


import { defineStore } from 'pinia';
import { Auths, modelsUser, NewAuths } from 'src/models';
import { useRouter } from 'vue-router';
import { Base64 } from 'js-base64'
import api from 'src/boot/api';

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
      localStorage.removeItem('auth_token')
      const r = useRouter()
      r.push({ name: 'login' })
    },
    fetchUserData() {
      let token = localStorage.getItem('auth_token')?.split('.');
      if (!token || token.length !== 3) {
        return false
      }
      let data = JSON.parse(Base64.decode(token[1]))
      if (data.id) {
        this.auth = NewAuths(data.Auth)
        api.user.get(data.id).then((e: modelsUser) => {
          console.log(e)
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
