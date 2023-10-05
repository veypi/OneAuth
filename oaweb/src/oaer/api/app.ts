/*
* @name: app
* @author: veypi <i@veypi.com>
* @date: 2021-11-17 14:44
* @description：ap
* @update: 2021-11-17 14:44
*/
import ajax from './axios'
import { Cfg } from './setting'

export default {
  local: () => Cfg.BaseUrl() + '/app/',
  self() {
    return ajax.get(this.local(), { option: 'oa' })
  },
  getKey(uuid: string) {
    return ajax.get(this.local() + uuid, { option: 'key' })
  },
  create(name: string, icon: string) {
    return ajax.post(this.local(), { name, icon })
  },
  get(uuid: string) {
    return ajax.get(this.local() + uuid)
  },
  list() {
    return ajax.get(this.local())
  },
  update(uuid: string, props: any) {
    return ajax.patch(this.local() + uuid, props)
  },
  user(uuid: string) {
    if (uuid === '') {
      uuid = '-'
    }
    return {
      local: () => this.local() + uuid + '/user/',
      list(id: string) {
        return ajax.get(this.local() + id)
      },
      add(uid: number) {
        return ajax.post(this.local() + uid)
      },
      update(uid: number, status: string) {
        return ajax.patch(this.local() + uid, { status })
      },
    }
  },
}
