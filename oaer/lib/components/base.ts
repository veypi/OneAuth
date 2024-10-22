/*
 * base.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-10-22 18:21
 * Distributed under terms of the GPL license.
 */
interface buildOpts {
  id?: string
  typ?: 'div'
  class?: string
  style?: string
  innerHtml?: string
  onclick?: any
  children?: HTMLElement[]
}

export default class {
  class_prefix: string
  constructor(p?: string) {
    this.class_prefix = p || 'voa-'
  }
  build(opts: buildOpts) {
    let dom = document.createElement(opts.typ || 'div')
    if (opts.id) {
      dom.id = opts.id
    }
    if (opts.class) {
      this.addClass(dom, opts.class)
    }
    if (opts.innerHtml) {
      dom.innerHTML = opts.innerHtml
    }
    if (opts.onclick) {
      dom.onclick = opts.onclick
    }
    if (opts.children) {
      for (let c in opts.children) {
        dom.appendChild(opts.children[c])
      }
    }
    if (opts.style) {
      const regex = /([a-zA-Z-]+)\s*:\s*([^;]+);?/g;
      let match;
      while ((match = regex.exec(opts.style)) !== null) {
        const key = match[1].trim();
        const value = match[2].trim();
        console.log([key, value])
        dom.style.setProperty(key, value)
      }
    }
    return dom
  }
  addClass(dom: HTMLElement, c: string) {
    let items = c.split(' ')
    for (let i of items) {
      if (i.startsWith(this.class_prefix)) {
        dom.classList.add(i)
      } else {
        dom.classList.add(this.class_prefix + i)
      }
    }
  }
  removeClass(dom: HTMLElement, c: string) {
    let items = c.split(' ')
    for (let i of items) {
      if (i.startsWith(this.class_prefix)) {
        dom.classList.remove(i)
      } else {
        dom.classList.remove(this.class_prefix + i)
      }
    }
  }
}
