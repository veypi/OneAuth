
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
    private readonly header: any

    constructor(method: Function, api: string, data?: any, headers?: any) {
        this.method = method
        this.api = api
        this.data = data
        this.header = headers
    }

    Start(success?: SuccessFunction<any>, fail?: FailedFunction<any>) {
        this.method(this.api, this.data, success, fail, this.header)
    }
}
