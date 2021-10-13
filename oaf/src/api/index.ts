/*
 * Copyright (C) 2019 light <light@light-laptop>
 *
 * Distributed under terms of the MIT license.
 */

import Vue from 'vue'
import {Base64} from 'js-base64'
import ajax from './ajax'
import store from '@/store'

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
        store.dispatch('handleLogOut')
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
          store.dispatch('handleLogOut')
          // bus.$emit('log_out')
        }
      }
    }
    this.method(this.api, this.data, newSuccess, newFail)
  }
}

const message = {
  count() {
    return new Interface(ajax.get, '/api/message/', {
      count: true,
      status: 'UnRead'
    })
  },
  get_content(id: number) {
    return new Interface(ajax.get, '/api/message/' + Number(id))
  },
  list(status: string) {
    return new Interface(ajax.get, '/api/message/', {status})
  },
  update(id: number, status: string) {
    return new Interface(ajax.patch, '/api/message/' + Number(id), {status})
  }
}

const role = {
  local: '/api/role/',
  get(id: number) {
    return new Interface(ajax.get, this.local + id)
  },
  list() {
    return new Interface(ajax.get, this.local)
  },
  update(id: number, props: any) {
    return new Interface(ajax.patch, this.local + id, props)
  },
  create(props: any) {
    return new Interface(ajax.post, this.local, props)
  },
  del(id: number) {
    return new Interface(ajax.delete, this.local + id)
  },
  bind(id: number, aid: number) {
    return new Interface(ajax.get, this.local + id + '/bind/' + aid)
  },
  unbind(id: number, aid: number) {
    return new Interface(ajax.get, this.local + id + '/unbind/' + aid)
  }
}

const app = {
  local: '/api/app/',
  get(id: string) {
    return new Interface(ajax.get, this.local + id)
  },
  list() {
    return new Interface(ajax.get, this.local)
  }
}

const user = {
  local: '/api/user/',
  register(username: string, password: string, prop?: any) {
    const data = Object.assign({
      username: username,
      password: Base64.encode(password)
    }, prop)
    return new Interface(ajax.post, this.local, data)
  },
  login(username: string, password: string) {
    return new Interface(ajax.head, this.local + username, {
      password: Base64.encode(password)
    })
  }
}

const api = {
  role: role,
  app: app,
  user: user,
  admin: {
    auths() {
      return new Interface(ajax.get, '/api/auth/')
    },
    user: {
      create(props: any) {
        const p = Object.assign({}, props)
        p.password = Base64.encode(props.password)
        return new Interface(ajax.post, '/api/user/', p)
      },
      update(user_id: number, props: any) {
        return new Interface(ajax.patch, '/api/user/' + user_id, props)
      },
      enable(user_id: number) {
        return new Interface(ajax.patch, '/api/user/' + user_id, {
          status: 'ok'
        })
      },
      disable(user_id: number) {
        return new Interface(ajax.patch, '/api/user/' + user_id, {
          status: 'disabled'
        })
      },
      attach_role(user_id: number, props: any) {
        return new Interface(ajax.post, '/api/user/' + user_id + '/role/', props)
      },
      detach_role(user_id: number, id: any) {
        return new Interface(ajax.delete, '/api/user/' + user_id + '/role/' + id)
      },
      reset_pass(user_id: number, password: string) {
        return new Interface(ajax.patch, '/api/user/' + user_id, {password})
      }
    }
  },
  auth: {
    event() {
      return {
        local: '/api/user/event/',
        list() {
          return new Interface(ajax.get, this.local)
        },
        create(title: string, tag: string, start_date: any, end_date: any) {
          return new Interface(ajax.post, this.local, {title, tag, start_date, end_date})
        },
        del(id: number) {
          return new Interface(ajax.delete, this.local + id)
        }
      }
    },
    favorite(name: string, tag: string, ok: boolean) {
      if (ok) {
        return new Interface(ajax.post, '/api/user/favorite', {name, tag})
      }
      return new Interface(ajax.delete, '/api/user/favorite', {name, tag})
    },
    get(id: number) {
      return new Interface(ajax.get, '/api/user/' + id)
    },
    search(username: string) {
      return new Interface(ajax.get, '/api/user/', {
        username
      })
    },
    login(username: string, password: string) {
      return new Interface(ajax.head, '/api/user/' + username, {
        password: Base64.encode(password)
      })
    },
    // @title 职位
    // @domain 部门
    register(username: string, password: string, domain?: string, title?: string) {
      return new Interface(ajax.post, '/api/user/', {
        username: username,
        password: Base64.encode(password),
        domain: domain,
        title: title
      })
    }
  },
  message: message
}

const Api = {
  install(vue: typeof Vue): void {
    vue.prototype.$api = api
  }
}
export {Api}
export default api
