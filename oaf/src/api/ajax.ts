import axios from 'axios'
import { store } from '@/store'
import msg from '@veypi/msg'


function getQueryVariable(variable: string) {
    let query = window.location.search.substring(1)
    let vars = query.split('&')
    for (let i = 0; i < vars.length; i++) {
        let pair = vars[i].split('=')
        if (pair[0] == variable) {
            return pair[1]
        }
    }
    return ''
}

function baseRequests(url: string, method: any = 'GET', query: any, data: any, success: any, fail?: Function, header?: any) {
    let headers = {
        auth_token: localStorage.auth_token || decodeURIComponent(getQueryVariable('token') as string),
    }
    if (header) {
        headers = Object.assign(headers, header)
    }
    return axios({
        url: url,
        params: query,
        data: data,
        method: method,
        headers: headers,
    }).then((res: any) => {
        if ('auth_token' in res.headers) {
            localStorage.auth_token = res.headers.auth_token
            store.commit('user/refreshToken', localStorage.auth_token)
        }
        if ('redirect_url' in res.headers) {
            window.location.href = res.headers.redirect_url
            return
        }
        console.log(res)
        if (method === 'HEAD') {
            success(res.headers)
        } else {
            success(res)
        }
    })
        .catch((e: any) => {
            if (typeof fail === 'function') {
                fail(e.response)
                return
            }
            let code = e.response.status
            if (code === 400) {
                msg.Warn(e.response.data)
                return
            } else if (code === 401) {
                console.log(e)
                store.commit('user/logout')
                return
            } else if (code === 500) {
                return
            }
            console.log(e)
        })
}

const ajax = {
    get(url: '', data = {}, success = {}, fail?: Function, header?: any) {
        return baseRequests(url, 'GET', data, {}, success, fail, header)
    },
    head(url: '', data = {}, success = {}, fail?: Function, header?: any) {
        return baseRequests(url, 'HEAD', data, {}, success, fail, header)
    },
    delete(url: '', data = {}, success = {}, fail?: Function, header?: any) {
        return baseRequests(url, 'DELETE', data, {}, success, fail, header)
    },
    post(url: '', data = {}, success = {}, fail?: Function, header?: any) {
        return baseRequests(url, 'POST', {}, data, success, fail, header)
    },
    put(url: '', data = {}, success = {}, fail?: Function, header?: any) {
        return baseRequests(url, 'PUT', {}, data, success, fail, header)
    },
    patch(url: '', data = {}, success = {}, fail?: Function, header?: any) {
        return baseRequests(url, 'PATCH', {}, data, success, fail, header)
    },
}

export default ajax
