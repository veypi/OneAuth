/*
 * Copyright (C) 2019 light <light@light-laptop>
 *
 * Distributed under terms of the MIT license.
 */

import {App} from 'vue'
import ajax from './ajax'
import {store} from '../store'
import {Base64} from 'js-base64'



export type SuccessFunction<T> = (e: any) => void;
export type FailedFunction<T> = (e: any) => void;

const Code = {
    42011: '无操作权限',
    22031: '资源不存在 或 您无权操作该资源'
}

class Interface {
    private readonly method: Function
    private readonly api: string
    private readonly data: any

    constructor(method: Function, api: string, data?: any) {
        this.method = method
        this.api = api
        this.data = data
    }

    Start(success: SuccessFunction<any>, fail?: FailedFunction<any>) {
        const newFail = function (data: any) {
            if (data && data.code === 40001) {
                // no login
                store.dispatch('handleLogout')
                return
            }
            // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
            // @ts-ignore
            if (data && data.code > 0 && Code[data.code]) {
            }
            if (fail) {
                fail(data)
            }
        }

        const newSuccess = function (data: any) {
            if (Number(data.status) === 1) {
                if (success) {
                    success(data.content)
                }
            } else {
                newFail(data)
                if (data.code === 41001) {
                    store.dispatch('handleLogout')
                    // bus.$emit('log_out')
                }
            }
        }
        this.method(this.api, this.data, newSuccess, newFail)
    }
}

const app = {
    local: '/api/app/',
    self() {
        return new Interface(ajax.get, this.local, {is_self: true})
    },
    get(id: string) {
        return new Interface(ajax.get, this.local + id)
    },
    list() {
        return new Interface(ajax.get, this.local)
    }
}

const user = {
    local: '/api/user/',
    register(username: string, password: string, uuid: string, prop?: any) {
        const data = Object.assign({
            username: username,
            uuid: uuid,
            password: Base64.encode(password)
        }, prop)
        return new Interface(ajax.post, this.local, data)
    },
    login(username: string, password: string, uuid: string) {
        return new Interface(ajax.head, this.local + username, {
            uid_type: 'username',
            uuid: uuid,
            password: Base64.encode(password)
        })
    }
}


const api = {
    user: user,
    app: app
}

const Api = {
    install(vue: App): void {
        vue.config.globalProperties.$api = api
    }
}
export {Api}
export default api
