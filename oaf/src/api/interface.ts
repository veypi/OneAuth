import {store} from "@/store";

export type SuccessFunction<T> = (e: any) => void;
export type FailedFunction<T> = (e: any) => void;

const Code = {
    42011: '无操作权限',
    22031: '资源不存在 或 您无权操作该资源'
}

export class Interface {
    private readonly method: Function
    private readonly api: string
    private readonly data: any

    constructor(method: Function, api: string, data?: any) {
        this.method = method
        this.api = api
        this.data = data
    }

    Start(success?: SuccessFunction<any>, fail?: FailedFunction<any>) {
        const newFail = function (data: any) {
            if (data) {
                if (data.code === 40001) {
                    // no login
                    store.commit('user/logout')
                    return
                    // @ts-ignore
                } else if (data.code === 42011 && window.$msg) {
                    // @ts-ignore
                    window.$msg.warning('无权限')
                }
            }
            // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
            // @ts-ignore
            if (data && data.code && Code[data.code]) {
            }
            if (fail) {
                fail(data.err)
            } else {
                // @ts-ignore
                window.$msg.warning(data.err)
            }
        }

        const newSuccess = function (data: any) {
            if (Number(data.status) === 1) {
                if (success) {
                    success(data.content)
                } else {
                    // @ts-ignore
                    window.$msg.warning('ok')
                }
            } else {
                newFail(data)
                if (data.code === 41001) {
                    store.commit('user/logout')
                    // bus.$emit('log_out')
                }
            }
        }
        this.method(this.api, this.data, newSuccess, newFail)
    }
}
