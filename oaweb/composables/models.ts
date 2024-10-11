/*
 * models.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-10-11 12:02
 * Distributed under terms of the MIT license.
 */


export * from './api/models'
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

