export interface Access { 
  created_at: Date
  updated_at: Date
  app_id: string
  user_id?: string
  role_id?: string
  name: string
  tid: string
  level: number
}
export interface App { 
  id: string
  created_at: Date
  updated_at: Date
  name: string
  icon: string
  des: string
  participate: string
  init_role_id?: string
  init_role?: Role
  init_url: string
  user_count: number
}
export interface AppUser { 
  id: string
  created_at: Date
  updated_at: Date
  app_id: string
  user_id: string
  status: string
}
export interface Resource { 
  created_at: Date
  updated_at: Date
  app_id: string
  name: string
  des: string
}
export interface Role { 
  id: string
  created_at: Date
  updated_at: Date
  name: string
  des: string
  app_id: string
  user_count: number
}
export interface Token { 
  id: string
  created_at: Date
  updated_at: Date
  user_id: string
  app_id: string
  expired_at: Date
  over_perm: string
  device: string
  ip: string
}
export interface User { 
  id: string
  created_at: Date
  updated_at: Date
  username: string
  nickname: string
  icon: string
  email: string
  phone: string
  status: number
}
export interface UserRole { 
  id: string
  created_at: Date
  updated_at: Date
  user_id: string
  role_id: string
  app_id: string
  status: string
}
