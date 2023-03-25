/*
* @name: resource
* @author: veypi <i@veypi.com>
* @date: 2021-11-18 15:52
* @description：resource
*/

import {Interface} from './interface'
import ajax from './ajax'
import {BaseUrl} from './setting'

export default (uuid: string) => {
    return {
        local: BaseUrl + 'app/' + uuid + '/resource/',
        list() {
            return new Interface(ajax.get, this.local)
        },
        update(id: number, props: {}) {
            return new Interface(ajax.patch, this.local + id, props)
        },
        create(Name: string, Des: string) {
            return new Interface(ajax.post, this.local, {Name, Des})
        },
        delete(id: number) {
            return new Interface(ajax.delete, this.local + id)
        },
    }
}
