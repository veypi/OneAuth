/*
 * gocliRequest.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-08-02 17:14
 * Distributed under terms of the MIT license.
 */


import axios, { AxiosError, type AxiosResponse } from 'axios';

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)

axios.defaults.withCredentials = true
const proxy = axios.create({
  withCredentials: true,
  headers: {
    'content-type': 'application/json;charset=UTF-8',
  },
});


// 请求拦截
const beforeRequest = (config: any) => {
  // 设置 token
  const token = util.getToken()
  config.retryTimes = 3
  // NOTE  添加自定义头部
  token && (config.headers.Authorization = `Bearer ${token}`)
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
  return Promise.resolve(data)
}

const responseFailed = (error: AxiosError) => {
  const { response, config } = error
  if (!window.navigator.onLine) {

    alert('没有网络')
    return Promise.reject(new Error('请检查网络连接'))
  }

  console.log(response)
  let needRetry = false
  if (!needRetry) {
    return Promise.reject(response?.data || response?.headers.error)
  };
  // @ts-ignore
  const { __retryCount = 0, retryDelay = 1000, retryTimes } = config;
  // 在请求对象上设置重试次数
  // @ts-ignore
  config.__retryCount = __retryCount;
  // 判断是否超过了重试次数
  if (__retryCount >= retryTimes) {
    return Promise.reject(response?.data || response?.headers.error)
  }
  // 增加重试次数
  // @ts-ignore
  config.__retryCount++;
  // 延时处理
  const delay = new Promise<void>((resolve) => {
    setTimeout(() => {
      resolve();
    }, retryDelay);
  });
  // 重新发起请求
  return delay.then(function() {
    return proxy.request(config as any);
  });
}

proxy.interceptors.response.use(responseSuccess, responseFailed)


interface data {
  params?: any
  req?: any
  headers?: any
}

export const webapi = {
  get<T>(url: string, req: data): Promise<T> {
    return proxy.request<T, any>({ method: 'get', url: url, headers: req.headers, data: req.req, params: req.params })
  },
  head<T>(url: string, req: data): Promise<T> {
    return proxy.request<T, any>({ method: 'head', url: url, headers: req.headers, data: req.req, params: req.params })
  },
  delete<T>(url: string, req: data): Promise<T> {
    return proxy.request<T, any>({ method: 'delete', url: url, headers: req.headers, data: req.req, params: req.params })
  },

  post<T>(url: string, req: data): Promise<T> {
    return proxy.request<T, any>({ method: 'post', url: url, headers: req.headers, data: req.req, params: req.params })
  },
  put<T>(url: string, req: data): Promise<T> {
    return proxy.request<T, any>({ method: 'put', url: url, headers: req.headers, data: req.req, params: req.params })
  },
  patch<T>(url: string, req: data): Promise<T> {
    return proxy.request<T, any>({ method: 'patch', url: url, headers: req.headers, data: req.req, params: req.params })
  },
}


export default webapi


