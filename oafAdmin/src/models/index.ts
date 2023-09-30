/*
* @name: index
* @author: veypi <i@veypi.com>
* @date: 2021-11-18 17:36
* @description：index
*/

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

    // Creator: number
    // Des: string
    // EnableEmail: boolean
    // EnablePhone: boolean
    // EnableRegister: true
    // EnableUser: boolean
    // EnableUserKey: boolean
    // EnableWx: boolean
    // Hide: boolean
    // Host: string
    // Icon: string
    // InitRole?: null
    // InitRoleID: number
    // Name: string
    // UUID: string
    // UserCount: number
    // UserKeyUrl: string
    // UserRefreshUrl: string
    // UserStatus: string
    // Users: null
}

export interface modelsUser {
    // Index 前端缓存
    Index?: number
    Apps: modelsApp[]
    Auths: null
    CreatedAt: string
    DeletedAt: null
    ID: number
    Icon: string
    Position: string
    Roles: null
    Status: string
    UpdatedAt: string
    Username: string
    Email: string
    Nickname: string
    Phone: string
}

export interface modelsSimpleAuth {
    level: number
    name: string
    rid: string
}

export interface modelsAuth {
    App?: modelsApp
    AppUUID: string
    CreatedAt: string
    DeletedAt: null
    ID: number
    Level: number
    RID: string
    RUID: string
    Resource?: modelsResource
    ResourceID: number
    Role?: modelsRole
    RoleID: number
    UpdatedAt: string
    User?: modelsUser
    UserID?: number
}

export interface modelsRole {
    App?: modelsApp
    AppUUID: string
    Auths: null
    CreatedAt: string
    DeletedAt: null
    ID: number
    Name: string
    Tag: string
    UpdatedAt: string
    UserCount: number
}

export interface modelsResource {
    App?: modelsApp
    AppUUID: string
    CreatedAt: string
    DeletedAt: null
    Des: string
    ID: number
    Name: string
    UpdatedAt: string
}