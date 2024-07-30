/*
* @name: app
* @author: veypi <i@veypi.com>
* @date: 2021-11-17 14:44
* @description：ap
* @update: 2021-11-17 14:44
*/
import ajax from './axios'
import cfg from '../cfg'

export default {
  local: () => cfg.BaseUrl() + '/app/',
  getKey(uuid: string) {
    return ajax.get(this.local() + uuid, { option: 'key' })
  },
  get(uuid: string) {
    return ajax.get(this.local() + uuid)
  },
  list() {
    return ajax.get(this.local())
  },
  users(uuid: string, user_id: string, data?: any) {
    if (uuid === '') {
      uuid = '-'
    }
    return ajax.get(this.local() + uuid + '/user/' + user_id, data)
  },
}
