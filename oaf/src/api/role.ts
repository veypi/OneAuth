import {BaseUrl} from './setting'
import {Interface} from './interface'
import ajax from './ajax'

export default (uuid: string) => {
    return {
        local: BaseUrl + 'app/' + uuid + '/role/',
        get(id: number) {
            return new Interface(ajax.get, this.local + id)
        },
        list() {
            return new Interface(ajax.get, this.local)
        },
        update(id: number, props) {
            return new Interface(ajax.patch, this.local + id, props)
        },
        create(Name: string, Tag: string) {
            return new Interface(ajax.post, this.local, {
                Name,Tag
            })
        },
        bind(id: number, aid: number) {
        },
    }
}
