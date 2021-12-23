import axios from 'axios'
import evt from '../evt'
import {Cfg} from './setting'


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
        auth_token: Cfg.token.value || decodeURIComponent(getQueryVariable('token') as string),
        uuid: Cfg.uuid.value,
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
            Cfg.token.value = res.headers.auth_token
        }
        if ('redirect_url' in res.headers) {
            window.location.href = res.headers.redirect_url
            return
        }
        if (method === 'HEAD') {
            success(res.headers)
        } else {
            success(res.data)
        }
    })
        .catch((e: any) => {
            if (e.response && e.response.status === 401) {
                evt.emit('logout')
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
