/*
* @name: index
* @author: veypi <i@veypi.com>
* @date: 2021-12-18 13:16
* @description：index
*/

import { App } from 'vue'
import OAer from './main.vue'
import './assets/icon.js'
import { Cfg, api } from './api'

export { OAer, Cfg, api }

export default {
  installed: false,
  install(vue: App, options?: any): void {
    if (this.installed) {
      return
    }
    this.installed = true
    vue.component('OAer', OAer)
  },
}
