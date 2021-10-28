// @ts-ignore
import axios from 'axios'
import {store} from '../store'


function baseRequests(url: string, method: any = 'GET', query: any, data: any, success: any, fail?: Function) {
    return axios({
        url: url,
        params: query,
        data: data,
        method: method,
        headers: {
            auth_token: localStorage.auth_token
        }
    }).then((res: any) => {
            if ('auth_token' in res.headers) {
                localStorage.auth_token = res.headers.auth_token
            }
            if (method === 'HEAD') {
                success(res.headers)
            } else {
                success(res.data)
            }
        })
        .catch((e: any) => {
            if (e.response && e.response.status === 401) {
                console.log(e)
                store.dispatch('handleLogout')
                return
            }
            console.log(e)
            if (e.response && e.response.status === 500) {
                return
            }
            if (typeof fail === 'function') {
                fail(e.response)
            } else if (e.response && e.response.status === 400) {
                console.log(400)
            } else {
                console.log(e.request)
            }
        })
}

const ajax = {
    get(url: '', data = {}, success = {}, fail?: Function) {
        return baseRequests(url, 'GET', data, {}, success, fail)
    },
    head(url: '', data = {}, success = {}, fail?: Function) {
        return baseRequests(url, 'HEAD', data, {}, success, fail)
    },
    delete(url: '', data = {}, success = {}, fail?: Function) {
        return baseRequests(url, 'DELETE', data, {}, success, fail)
    },
    post(url: '', data = {}, success = {}, fail?: Function) {
        return baseRequests(url, 'POST', {}, data, success, fail)
    },
    put(url: '', data = {}, success = {}, fail?: Function) {
        return baseRequests(url, 'PUT', {}, data, success, fail)
    },
    patch(url: '', data = {}, success = {}, fail?: Function) {
        return baseRequests(url, 'PATCH', {}, data, success, fail)
    }
}

export default ajax
