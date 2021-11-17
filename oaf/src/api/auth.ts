import {Interface} from './interface'
import ajax from './ajax'
import {BaseUrl} from './setting'

export default (uuid: string) => {
    return {
        local: BaseUrl + 'app/' + uuid + '/auth/',
        get(id: number) {
            return new Interface(ajax.get, this.local + id)
        },
        list(appUUID: string, user_id: number) {
            return new Interface(ajax.get, this.local, {uuid: appUUID, id: user_id})
        },
    }
}

