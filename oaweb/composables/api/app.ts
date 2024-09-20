import webapi from "./gocliRequest"
import * as components from "./mainComponents"
export * from "./mainComponents"

/**
 * @description 
 * @param params
 * @param req
 */
export function login(pa: string, req: components.AppReq, params: components.AppReqParams) {
	return webapi.get<components.AppResp>(`/api/app/login/${pa}`, {"params":params, "req":req})
}
