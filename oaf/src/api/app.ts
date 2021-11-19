/*
* @name: app
* @author: veypi <i@veypi.com>
* @date: 2021-11-17 14:44
* @description：ap
* @update: 2021-11-17 14:44
*/
import {Interface} from './interface'
import ajax from './ajax'
import {BaseUrl} from './setting'

export default {
    local: BaseUrl + 'app/',
    self() {
        return new Interface(ajax.get, this.local, {option: 'oa'})
    },
    getKey(uuid: string) {
        return new Interface(ajax.get, this.local + uuid, {option: 'key'})
    },
    create(name: string, icon: string) {
        return new Interface(ajax.post, this.local, {name, icon})
    },
    get(uuid: string) {
        return new Interface(ajax.get, this.local + uuid)
    },
    list() {
        return new Interface(ajax.get, this.local)
    },
    update(uuid: string, props: any) {
        return new Interface(ajax.patch, this.local + uuid, props)
    },
    user(uuid: string) {
        if (uuid === '') {
            uuid = '-'
        }
        return {
            local: this.local + uuid + '/user/',
            list(uid: number) {
                return new Interface(ajax.get, this.local + uid)
            },
            add(uid: number) {
                return new Interface(ajax.post, this.local + uid)
            },
            update(uid: number, status: string) {
                return new Interface(ajax.patch, this.local + uid, {status})
            },
        }
    },
}
