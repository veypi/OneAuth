/*
* @name: setting
* @author: veypi <i@veypi.com>
* @date: 2021-11-17 15:45
* @description：setting
* @update: 2021-11-17 15:45
*/

import { ref } from 'vue'

export let Cfg = {
  token: ref(''),
  uuid: ref(''),
  host: ref(''),
  prefix: '/api',
  BaseUrl() {
    return this.host.value + this.prefix
  },
  goto(url: string) {
    if (!url.startsWith('/')) {
      url = '/' + url
    }
    window.location.href = this.host.value + '/#' + url
  },
  userFileUrl() {
    return (this.host.value || window.location.host) + '/file/'
  },
}
