/*
 * app.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-30 17:26
 * Distributed under terms of the MIT license.
 */


import { defineStore } from 'pinia';
import { useQuasar } from 'quasar'

export const useAppStore = defineStore('app', {
  state: () => ({
    id: '',
    is_dark: false,
    title: '',
  }),
  getters: {
  },
  actions: {
    toggle_mode() {

    }
  },
});
