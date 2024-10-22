/*
 * slide.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-10-22 17:57
 * Distributed under terms of the GPL license.
 */
import Base from './base'
import bus from '../bus'
import account from './account'

/*
mask
  slide
    header
    body
      main
      footer
 *
 * */
export default class extends Base {
  mask: HTMLDivElement
  slide: HTMLDivElement
  header: HTMLDivElement
  body: HTMLElement
  main: HTMLElement
  footer: HTMLElement
  constructor() {
    super()
    this.header = this.build({
      class: 'slide-header animate-slow',
    })
    this.footer = this.build({
      class: 'slide-footer',
      innerHtml: 'logout',
      onclick: () => {
        bus.emit('logout')
      }
    })
    this.main = this.build({
      class: 'slide-main',
      children: [new account().main]
    })
    this.body = this.build({
      class: 'slide-body animate-slow',
      style: 'animation-delay: 300ms',
      children: [this.main, this.footer]
    })
    this.slide = this.build({
      id: 'voa-slide',
      class: 'slide',
      children: [this.header, this.body]
    })
    this.mask = this.build({
      class: 'slide-mask',
      style: 'visibility: hidden',
      children: [this.slide],
      onclick: (e: MouseEvent) => {
        if (e.target === e.currentTarget) {
          this.hide()
        }
      }
    })
    document.body.appendChild(this.mask)
  }
  show() {
    this.mask.style.visibility = 'visible'
    this.addClass(this.header, 'slidein-right')
    this.addClass(this.body, 'slidein-up')
  }
  hide() {
    this.mask.style.visibility = 'hidden'
    this.removeClass(this.header, 'slidein-right')
    this.removeClass(this.body, 'slidein-up')
  }
}
