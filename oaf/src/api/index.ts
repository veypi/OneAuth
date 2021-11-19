/*
 * Copyright (C) 2019 light <veypi@light-laptop>
 *
 * Distributed under terms of the MIT license.
 */

import {App} from 'vue'
import role from "./role";
import app from './app'
import user from './user'
import auth from './auth'
import resource from './resource'


const api = {
    user: user,
    app: app,
    auth: auth,
    role: role,
    resource: resource
}

const Api = {
    install(vue: App): void {
        vue.config.globalProperties.$api = api
    }
}
export {Api}
export default api
