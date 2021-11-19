import {Base64} from "js-base64";
import {Interface} from './interface'
import ajax from './ajax'
import {BaseUrl} from './setting'

export default {
    local: BaseUrl + 'user/',
    register(username: string, password: string, prop?: any) {
        const data = Object.assign({
            username: username,
            password: Base64.encode(password)
        }, prop)
        return new Interface(ajax.post, this.local, data)
    },
    login(username: string, password: string) {
        return new Interface(ajax.head, this.local + username, {
            UidType: 'username',
            password: Base64.encode(password)
        })
    },
    get(id: number) {
        return new Interface(ajax.get, this.local + id)
    },
    list() {
        return new Interface(ajax.get, this.local)
    },
    update(id: number, props: any) {
        return new Interface(ajax.patch, this.local + id, props)
    }
}
