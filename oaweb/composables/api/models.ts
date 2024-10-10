export interface Access { 
  app_id: String
  app?: App
  user_id?: String
  user?: User
  role_id?: String
  role?: Role
  name: String
  t_id: String
  level: Number
}
export interface App { 
  name: String
  icon: String
  des: String
  participate: String
  init_role_id?: String
  init_role?: Role
  init_url: String
  user_count: Number
  key: String
}
export interface AppUser { 
  app_id: String
  app?: App
  user_id: String
  user?: User
  status: String
}
export interface Resource { 
  app_id: String
  app?: App
  name: String
  des: String
}
export interface Role { 
  name: String
  des: String
  app_id: String
  app?: App
  user_count: Number
  access: any
}
export interface Testb { 
  test_id: String
  test?: Test
  name: String
}
export interface Test { 
  name: String
}
export interface Token { 
  user_id: String
  user?: User
  app_id: String
  app?: App
  expired_at: Date
  over_perm: String
  device: String
}
export interface User { 
  username: String
  nickname: String
  icon: String
  email: String
  phone: String
  status: Number
  salt: String
  code: String
}
export interface UserRole { 
  user_id: String
  user?: User
  role_id: String
  role?: Role
  app_id: String
  app?: App
  status: String
}
