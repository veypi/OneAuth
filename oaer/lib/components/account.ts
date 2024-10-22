/*
 * account.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-10-22 22:07
 * Distributed under terms of the GPL license.
 */
import Base from './base'

export default class extends Base {
  doms: { [key: string]: HTMLElement }
  main: HTMLElement
  constructor() {
    super()
    this.main = this.build({
      class: 'voa-account'
    })
    let u = {
      username: 'asd',
      icon: 'https://public.veypi.com/img/avatar/0001.jpg'
    }
    this.main.innerHTML = `
<div class="voa-account-header">
  <div class="vah-1">my account</div>
  <div class="vah-2">account center</div>
</div>
<div class="voa-account-body">
  <div class="vab-ico">
        <img style="" class="" src="${u.icon}" />
  </div>
  <div class="vab-info">
    <div class="vabi-1"><span>昵称：</span> <span>${u.username}</span> </div>
    <div class="vabi-2"><span>昵称：</span> <span>${u.username}</span> </div>
    <div class="vabi-3"><span>昵称：</span> <span>${u.username}</span> </div>
    <div class="vabi-4"><span>昵称：</span> <span>${u.username}</span> </div>
  </div>
</div>
`
  }
}
