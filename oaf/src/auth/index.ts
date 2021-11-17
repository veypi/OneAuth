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

const level = {
    None: 0,
    Do: 1,
    Part: 1,
    Read: 2,
    Create: 3,
    Update: 4,
    Delete: 5,
    All: 6
}

export interface Auth {
    rid: string
    ruid: string
    level: number
}

class authLevel {
    level = level.None
    constructor(level: number) {
        this.level = level
    }
    CanDo(): boolean {
        return this.level >= level.Do
    }
    CanRead(): boolean {
        return this.level >= level.Read
    }
    CanCreate(): boolean {
        return this.level >= level.Create
    }
    CanUpdate(): boolean {
        return this.level >= level.Update
    }
    CanDelete(): boolean {
        return this.level >= level.Delete
    }
    CanDoAny(): boolean {
        return this.level >= level.All
    }
}

export class Auths {
    private readonly list: [Auth]

    constructor(auths: [Auth]) {
        this.list = auths
    }

    Get(rid: string, ruid: string): authLevel {
        let l = level.None
        for (let i of this.list) {
            if (i.rid == rid && (i.ruid === '' || i.ruid === ruid) && i.level > l) {
                l = i.level
            }
        }
        return new authLevel(l)
    }
}


export function NewAuths(a: any) {
    return new Auths(a)
}
