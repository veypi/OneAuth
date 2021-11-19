/*
* @name: index
* @author: veypi <i@veypi.com>
* @date: 2021-11-18 17:36
* @description：index
*/

export interface modelsBread {
    Index: number
    Name: string
    Type?: string
    RName: string
    RParams?: any
    RQuery?: any
}


export interface modelsApp {
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: null
    Creator: number
    Des: string
    EnableEmail: boolean
    EnablePhone: boolean
    EnableRegister: true
    EnableUser: boolean
    EnableUserKey: boolean
    EnableWx: boolean
    Hide: boolean
    Host: string
    Icon: string
    InitRole: null
    InitRoleID: number
    Name: string
    UUID: string
    UserCount: number
    UserKeyUrl: string
    UserRefreshUrl: string
    UserStatus: string
    Users: null
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
    Level: number
    RID: string
    RUID: string
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
