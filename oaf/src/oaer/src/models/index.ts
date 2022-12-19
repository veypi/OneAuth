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

export enum AppJoin {
    Auto = 0,
    Disabled = 1,
    Applying = 2,
}


export interface App {
    created: string
    updated: string
    delete_flag: boolean
    des: string
    hide: boolean
    icon: string
    id: string
    join_method: AppJoin
    name: string
    redirect: string
    role_id: string
    status: number
    user_count: number

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
    Apps: App[]
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

export interface modelsAuth {
    App?: App
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
    App?: App
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
    App?: App
    AppUUID: string
    CreatedAt: string
    DeletedAt: null
    Des: string
    ID: number
    Name: string
    UpdatedAt: string
}
