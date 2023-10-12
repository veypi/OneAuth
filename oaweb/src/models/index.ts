/*
* @name: index
* @author: veypi <i@veypi.com>
* @date: 2021-11-18 17:36
* @description：index
*/

import { RouteLocationRaw } from 'vue-router';
import { AccessLevel } from './auth';

export { type Auths, type modelsSimpleAuth, NewAuths, R, AccessLevel, LevelOptions } from './auth'

export type Dict = { [key: string]: any }

export enum ArgType {
  Text = 'text',
  Password = 'password',
  Bool = 'bool',
  Select = 'select',
  Radio = 'radio',
  Number = 'number',
  Region = 'region',
  NumList = 'numList',
  StrList = 'strList',
  Table = 'table',
  Grid = 'grid',
  File = 'file',
  Img = 'img'
}

export const ArgTypesTrans = {
  [ArgType.Text]: '文本',
  [ArgType.Password]: '密码',
  [ArgType.Select]: '选择器',
  [ArgType.Radio]: '单选框',
  [ArgType.Number]: '数字',
  [ArgType.Region]: '区间',
  [ArgType.NumList]: '数组',
  [ArgType.StrList]: '文本集合',
  [ArgType.Table]: '表格',
  [ArgType.Grid]: '矩阵',
  [ArgType.File]: '文件',
  [ArgType.Img]: '图片',
  [ArgType.Bool]: '开关',
}

export interface DocItem {
  name: string
  url: string
  version?: string

}
export interface DocGroup {
  name: string
  icon: string
  items?: DocItem[]
}

export interface MenuLink {
  title: string;
  caption?: string;
  to?: RouteLocationRaw;
  link?: string;
  icon?: string;
}

export interface modelsBread {
  Index: number
  Name: string
  Type?: string
  RName: string
  RParams?: any
  RQuery?: any
}


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
  app: modelsApp,
  user: modelsUser,
  app_id: string
  user_id: string
  status: AUStatus
}

export interface modelsUser {
  id: string
  created: string
  updated: string
  username: string
  nickname: string
  email: string
  phone: string
  icon: string
  status: number
  used: number
  space: number
  au: AUStatus
}

export interface modelsAccess {
  id: number
  created: string
  updated: string
  name: string
  app_id: string,
  role_id?: string,
  user_id?: string,
  level: AccessLevel,
  rid?: string
}

export interface modelsRole {
  created: string
  updated: string
  app_id: string
  id: string
  name: string
  des: string
  user_count: number
}

export interface modelsResource {
  created: string
  updated: string
  app_id: string
  name: string
  des: string
}
