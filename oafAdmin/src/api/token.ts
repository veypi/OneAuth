/*
* @name: token
* @author: veypi <i@veypi.com>
* @date: 2021-11-26 19:22
* @description：token
*/

import {Interface} from '@/api/interface'
import ajax from './ajax'

export default (uuid: string) => {
    return {
        local: '/api/app/' + uuid + '/token/',
        get() {
            return new Interface(ajax.get, this.local)
        },
    }
}
