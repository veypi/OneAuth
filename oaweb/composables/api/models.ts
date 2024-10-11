export interface Access { 
  app_id: string
  app?: App
  user_id?: string
  user?: User
  role_id?: string
  role?: Role
  name: string
  t_id: string
  level: number
}
export interface App { 
  name: string
  icon: string
  des: string
  participate: string
  init_role_id?: string
  init_role?: Role
  init_url: string
  user_count: number
  key: string
}
export interface AppUser { 
  app_id: string
  app?: App
  user_id: string
  user?: User
  status: string
}
export interface Resource { 
  app_id: string
  app?: App
  name: string
  des: string
}
export interface Role { 
  name: string
  des: string
  app_id: string
  app?: App
  user_count: number
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
  user_id: string
  user?: User
  app_id: string
  app?: App
  expired_at: Date
  over_perm: string
  device: string
}
export interface User { 
  username: string
  nickname: string
  icon: string
  email: string
  phone: string
  status: number
  salt: string
  code: string
}
export interface UserRole { 
  user_id: string
  user?: User
  role_id: string
  role?: Role
  app_id: string
  app?: App
  status: string
}
