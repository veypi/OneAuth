import {BaseUrl} from './setting'
import {Interface} from './interface'
import ajax from './ajax'

export default (uuid: string) => {
    return {
        local: BaseUrl +'app/' + uuid + '/role/',
        get(id: number) {
            return new Interface(ajax.get, this.local + id)
        },
        list() {
            return new Interface(ajax.get, this.local)
        },
        create(uuid: string, name: string) {
            return new Interface(ajax.post, this.local, {
                uuid: uuid,
                name: name,
            })
        },
        bind(id: number, aid: number) {
        },
    }
}
