package models

type User struct {
	BaseModel
	Username string `json:"username" gorm:"varchar(100);unique;default:not null" methods:"post,put,*patch,*list" parse:"json"`
	Nickname string `json:"nickname" methods:"post,put,*patch,*list" parse:"json"`
	Icon     string `json:"icon" methods:"post,put,*patch" parse:"json"`

	Email string `json:"email" gorm:"varchar(20);unique;default:null" methods:"post,put,*patch,*list" parse:"json"`
	Phone string `json:"phone" gorm:"varchar(50);unique;default:null" methods:"post,put,*patch,*list" parse:"json"`

	Status uint `json:"status" methods:"post,put,*patch,*list" parse:"json"`

	RealCode  string `json:"-"`
	CheckCode string `json:"-"`
}

type UserRole struct {
	BaseModel
	UserID string `json:"user_id" methods:"post,*patch,delete" parse:"path"`
	User   *User  `json:"user"`
	RoleID string `json:"role_id" methods:"post,delete" parse:"path"`
	Role   *Role  `json:"role"`
	Status string `json:"status" methods:"post,put,*patch,*list" parse:"json"`
}
