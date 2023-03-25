import {Interface} from './interface'
import ajax from './ajax'
import {BaseUrl} from './setting'

export default (uuid: string) => {
    return {
        local: BaseUrl + 'app/' + uuid + '/auth/',
        get(id: number) {
            return new Interface(ajax.get, this.local + id)
        },
        del(id: number) {
            return new Interface(ajax.delete, this.local + id)
        },
        update(id: number, ResourceID: number, RUID: string, Level: number) {
            return new Interface(ajax.patch, this.local + id, {
                ResourceID,
                RUID,
                Level,
            })
        },
        create(ResourceID: number, UserID: number | null, RoleID: number | null, RUID: string, Level: number) {
            return new Interface(ajax.post, this.local, {
                ResourceID,
                UserID,
                RoleID,
                RUID,
                Level,
            })
        },
        listOfUser(user_id: number) {
            return new Interface(ajax.get, this.local, {uid: user_id})
        },
        listOfRole(id: number) {
            return new Interface(ajax.get, this.local, {rid: id})
        },
    }
}

