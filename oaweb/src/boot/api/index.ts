/*
 * index.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-09-22 20:17
 * Distributed under terms of the MIT license.
 */

import app from "./app";
import role from "./role";
import token from "./token";
import user from "./user";
import resource from "./resource";
import access from './access';



const api = {
  user,
  app,
  token,
  role,
  resource,
  access,
}


export default api;

