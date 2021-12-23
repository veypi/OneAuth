/*
* @name: app
* @author: veypi <i@veypi.com>
* @date: 2021-11-17 14:44
* @description：ap
* @update: 2021-11-17 14:44
*/
import {Interface} from './interface'
import ajax from './ajax'
import {Cfg} from './setting'

export default {
    local: () => Cfg.BaseUrl() + 'app/',
    get(uuid: string) {
        return new Interface(ajax.get, this.local() + uuid)
    },
    list() {
        return new Interface(ajax.get, this.local())
    },
    user(uuid: string) {
        if (uuid === '') {
            uuid = '-'
        }
        return {
            local: () => this.local() + uuid + '/user/',
            list(uid: number) {
                return new Interface(ajax.get, this.local() + uid)
            },
        }
    },
}
