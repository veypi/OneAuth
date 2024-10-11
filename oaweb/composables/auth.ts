/*
 * auth.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-10-11 12:01
 * Distributed under terms of the MIT license.
 */

export interface modelsSimpleAuth {
  level: number
  name: string
  tid: string
}


export const R = {
  // 应用管理配置权限
  App: 'app',
  // 用户管理和绑定应用权限
  User: 'user',
  // 权限资源定义权限
  Resource: 'resource',
  // 角色管理和绑定用户权限
  Role: 'role',
  // 权限管理和绑定角色权限
  Auth: 'auth',
}


export enum AccessLevel {
  None = 0,
  Do = 1,
  Read = 1,
  Create = 2,
  Update = 3,
  Delete = 4,
  All = 5
}

export const LevelOptions = [
  { key: 0, name: '无权限' },
  { key: 1, name: '读取数据权限' },
  { key: 2, name: '创建数据权限' },
  { key: 3, name: '更新数据权限' },
  { key: 4, name: '删除数据权限' },
  { key: 5, name: '管理员权限' },
]

class authLevel {
  level = AccessLevel.None
  constructor(level: number) {
    this.level = level
  }
  CanDo(): boolean {
    return this.level >= AccessLevel.Do
  }
  CanRead(): boolean {
    return this.level >= AccessLevel.Read
  }
  CanCreate(): boolean {
    return this.level >= AccessLevel.Create
  }
  CanUpdate(): boolean {
    return this.level >= AccessLevel.Update
  }
  CanDelete(): boolean {
    return this.level >= AccessLevel.Delete
  }
  CanDoAny(): boolean {
    return this.level >= AccessLevel.All
  }
}

export class auths {
  private readonly list: modelsSimpleAuth[]

  constructor(auths: modelsSimpleAuth[]) {
    this.list = auths
  }

  Get(name: string, rid: string): authLevel {
    let l = AccessLevel.None
    for (let i of this.list) {
      if (i.name == name && (!i.tid || i.tid === rid) && i.level > l) {
        l = i.level
      }
    }
    return new authLevel(l)
  }
}

export interface Auths {
  Get(name: string, rid: string): authLevel
}


export function NewAuths(a: modelsSimpleAuth[]): Auths {
  return new auths(a)
}

