import webapi from "./gocliRequest"
import * as components from "./mainComponents"
export * from "./mainComponents"

/**
 * @description 
 * @param req
 */
export function reg(req: components.RegReq) {
	return webapi.post<null>(`/api/user`, {"req":req})
}

/**
 * @description 
 * @param params
 */
export function login(id: string, params: components.LoginReqParams) {
	return webapi.head<null>(`/api/user/${id}`, {"params":params})
}

/**
 * @description 
 */
export function list() {
	return webapi.get<Array<components.UserResp>>(`/api/user`, {})
}

/**
 * @description 
 */
export function get(id: number) {
	return webapi.get<components.UserResp>(`/api/user/${id}`, {})
}
