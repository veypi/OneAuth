/*
 * main.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-07-04 14:34
 * Distributed under terms of the MIT license.
 */

import './assets/css/oaer.scss'


export class OAer {
  domid: string
  id: string
  dom: {
    mask: HTMLDivElement
    frame: HTMLDivElement
    b0?: HTMLDivElement
    b1?: HTMLDivElement
    login_box?: HTMLDivElement
  }
  constructor(id: string, domid?: string) {
    this.id = id
    this.domid = domid || 'oaer'
    this.dom = {
      frame: document.querySelector(`#${this.domid}`)!,
      mask: document.createElement('div'),
    }
    this.dom.mask.id = 'voa-mask'
    document.body.appendChild(this.dom.mask)
    this.dom.frame.classList.add('voa')
    this.build_b0()
    this.dom.frame.appendChild(this.dom.b0!)
    this.offclick()
  }
  build_b0() {
    this.dom.b0 = document.createElement('div')
    this.dom.b0.classList.add('hover-line-b')
    this.dom.b0.onclick = () => {
      console.log('click b0')
      this.build_login()
      // this.dom.b1?.classList.add('voa-animate-scale-off')
      // this.build_b1()
      // this.dom.frame.replaceChild(this.dom.b1!, this.dom.b0!)
    }
    this.dom.b0.innerHTML = `
    登录
`
  }
  build_login() {
    if (this.dom.login_box) {
      return
    }
    this.dom.login_box = document.createElement('div')
    this.dom.login_box.classList.add('voa-modal', 'voa-scale-up')
    this.dom.login_box.innerHTML = `
<div class="voa-login-box">
    <svg class="close" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path d="M18 6L6 18M6 6L18 18" stroke="black" stroke-width="2"/>
    </svg>
  <div class="header">
    <div class="voa-logo"></div>
    <div class="txt">OneAuth</div>
  </div>
  <div class="username">
    <input autocomplete="username" placeholder="username, phone or Email">
  </div>
  <div class="password">
    <input autocomplete="password" type='password' placeholder="password">
  </div>
  <button class='ok voa-btn'>
    login
  </button>
  <div class="last">
    <div class="icos">
      <div class="github"></div>
      <div class="wechat"></div>
      <div class="google"></div>
    </div>
    <div class="txt">
      <div>Create Account</div>
      <div>Forgot Password?</div>
    </div>
  </div>
</div>
`
    document.body.appendChild(this.dom.login_box)
    document.querySelector('.voa-login-box .close')?.addEventListener('click', () => {
      this.dom.login_box?.classList.add('voa-scale-off')
      setTimeout(() => {
        this.dom.login_box?.remove()
        this.dom.login_box = undefined
      }, 300)
    })
    let uin = document.querySelector('.voa-login-box .username input') as HTMLInputElement
    console.log(uin)
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



