/*
 * main.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-07-04 14:34
 * Distributed under terms of the MIT license.
 */

import './assets/css/oaer.scss'
import bus from './bus'

import ui from './components'

export class OAer {
  host: string
  domid?: string
  ui?: ui
  constructor(host: string, domid?: string) {
    this.host = host
    if (domid) {
      this.domid = domid
      this.ui = new ui(document.querySelector(`#${this.domid}`)!)
    }
  }
  login() {
    bus.emit('login')
  }
  logout() {
    bus.emit('logout')
  }
  onlogout(fc: () => void) {
    bus.on('logout', fc)
  }
}


