/*
 * main.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-07-04 14:34
 * Distributed under terms of the MIT license.
 */

import './css/oaer.scss'

export class OAer {
  domid: string
  id: string
  dom: {
    frame: HTMLDivElement
    b0?: HTMLDivElement
    b1?: HTMLDivElement
  }
  constructor(id: string, domid?: string) {
    this.id = id
    this.domid = domid || 'oaer'
    this.dom = {
      frame: document.querySelector(`#${this.domid}`)!
    }
    this.dom.frame.classList.add('voa')
    this.build_b0()
    this.dom.frame.appendChild(this.dom.b0!)
    this.offclick()
  }
  build_b0() {
    this.dom.b0 = document.createElement('div')
    this.dom.b0.classList.add('voa-off')
    this.dom.b0.onclick = () => {
      this.dom.b1?.classList.add('voa-animate-scale-off')
      this.build_b1()
      this.dom.frame.replaceChild(this.dom.b1!, this.dom.b0!)
    }
    this.dom.b0.innerHTML = `
  <span>
    登录
  </span>
`
  }
  build_b1() {
    this.dom.b1 = document.createElement('div')
    this.dom.b1.classList.add('voa-on')
    this.dom.b1.innerHTML = `
        <img style="" class="" src="" />
`
  }
  offclick() {
  }
}



