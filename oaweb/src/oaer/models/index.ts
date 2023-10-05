/*
* @name: index
* @author: veypi <i@veypi.com>
* @date: 2021-11-18 17:36
* @description：index
*/


export interface modelsApp {
  created: string
  updated: string
  delete_flag: boolean
  des: string
  hide: boolean
  icon: string
  id: string
  name: string
  redirect: string
  role_id: string
  status: number
  user_count: number

  au: modelsAppUser
}

export enum AUStatus {
  OK = 0,
  Disabled = 1,
  Applying = 2,
  Deny = 3,
}

export interface modelsAppUser {
  app_id: string
  user_id: string
  status: AUStatus
}

export interface modelsUser {
  id: string
  created: string
  updated: string
  delete_flag: boolean
  username: string
  nickname: string
  email: string
  phone: string
  icon: string
  status: number
  used: number
  space: number
}

