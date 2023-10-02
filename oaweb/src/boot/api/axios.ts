/*
 * axios.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-22 20:22
 * Distributed under terms of the MIT license.
 */

import axios, { AxiosError, AxiosInstance, AxiosResponse } from 'axios';
import msg from '@veypi/msg'

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)

const proxy = axios.create({
  baseURL: '/api/',
  withCredentials: true,
  headers: {
    'content-type': 'application/json;charset=UTF-8',
  },
});


// 请求拦截
const beforeRequest = (config: any) => {
  // 设置 token
  const token = localStorage.getItem('auth_token')
  // NOTE  添加自定义头部
  token && (config.headers.auth_token = token)
  // config.headers['auth_token'] = ''
  return config
}
proxy.interceptors.request.use(beforeRequest)

// 响应拦截器
const responseSuccess = (response: AxiosResponse) => {
  // eslint-disable-next-line yoda
  // 这里没有必要进行判断，axios 内部已经判断
  // const isOk = 200 <= response.status && response.status < 300
  let data = response.data
  if (response.config.method === 'head') {
    data = JSON.parse(JSON.stringify(response.headers))
  }
  return Promise.resolve(data)
}

const responseFailed = (error: AxiosError) => {
  const { response } = error
  if (!window.navigator.onLine) {

    alert('没有网络')
    return Promise.reject(new Error('请检查网络连接'))
  }
  console.log(response)
  return Promise.reject(response?.data || response?.headers.error)
}

proxy.interceptors.response.use(responseSuccess, responseFailed)


const ajax = {
  get(url: string, data = {}, header?: any) {
    return proxy.get<any, any>(url, { params: data, headers: header })
  },
  head(url: string, data = {}, header?: any) {
    return proxy.head<any, any>(url, { params: data, headers: header })
  },
  delete(url: string, data = {}, header?: any) {
    return proxy.delete<any, any>(url, { params: data, headers: header })
  },

  post(url: string, data = {}, header?: any) {
    return proxy.post<any, any>(url, data, { headers: header })
  },
  put(url: string, data = {}, header?: any) {
    return proxy.put<any, any>(url, data, { headers: header })
  },
  patch(url: string, data = {}, header?: any) {
    return proxy.patch<any, any>(url, data, { headers: header })
  },
}

export default ajax

