/*
 * index.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-10-22 17:51
 * Distributed under terms of the GPL license.
 */

import slide from './slide'
import Base from './base'
import bus from '../bus'
export default class extends Base {
  slide: slide
  frame: HTMLDivElement
  frame_login?: HTMLDivElement
  frame_user?: HTMLDivElement
  constructor(frame: HTMLDivElement) {
    super()
    this.frame = frame
    this.frame.classList.add('voa')
    this.slide = new slide()
    this.mount_user()
    bus.on('logout', () => {
      this.mount_login()
      this.slide.hide()
    })
  }
  mount_login() {
    this.frame_login = this.build({
      class: 'off hover-line-b scale-in',
      innerHtml: '登录',
      onclick: () => {
        console.log('click login')
        this.addClass(this.frame_login!, 'scale-off')
        this.mount_user()
      }
    })
    if (this.frame_user) {
      this.frame.removeChild(this.frame_user)
    }
    this.frame.appendChild(this.frame_login)
  }
  mount_user() {
    let icon = 'https://public.veypi.com/img/avatar/0001.jpg'
    this.frame_user = this.build({
      class: 'on scale-in',
      innerHtml: `
        <img style="" class="" src="${icon}" />
`,
      onclick: () => {
        this.slide.show()
        // this.mount_login()
      }
    })
    if (this.frame_login) {
      this.frame.removeChild(this.frame_login)
    }
    this.frame.appendChild(this.frame_user)
  }
}

