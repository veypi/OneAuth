/*
 * Copyright (C) 2019 light <veypi@light-laptop>
 *
 * Distributed under terms of the MIT license.
 */

import user from './user'
import app from './app'
import {Cfg} from './setting'


const api = {
    user: user,
    app: app,
}

export {api, Cfg}
export default api
