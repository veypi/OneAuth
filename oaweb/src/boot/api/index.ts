/*
 * index.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-22 20:17
 * Distributed under terms of the MIT license.
 */

import app from "./app";
import token from "./token";
import user from "./user";




const api = {
  user: user,
  app: app,
  token: token
}


export default api;

