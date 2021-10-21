import Vue from 'vue'
import OneIcon from './icon.vue'

function loadJS(url: string) {
  const script = document.createElement('script')
  script.type = 'text/javascript'
  script.src = url
  document.getElementsByTagName('head')[0].appendChild(script)
}

export default {
  installed: false,
  install(vue: typeof Vue, options?: { href: '' }): void {
    if (this.installed) {
      return
    }
    this.installed = true
    if (options && options.href) {
      console.log(options.href)
      loadJS(options.href)
    } else {
      console.error('not set iconfont href')
    }
    vue.component('one-icon', OneIcon)
  }
}
