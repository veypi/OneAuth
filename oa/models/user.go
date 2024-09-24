package models

type User struct {
	BaseModel
	Username string `json:"username" gorm:"type:varchar(100);unique;default:not null" methods:"post,*patch,*list" parse:"json"`
	Nickname string `json:"nickname" gorm:"type:varchar(100)" methods:"*post,*patch,*list" parse:"json"`
	Icon     string `json:"icon" methods:"*post,*patch" parse:"json"`

	Email string `json:"email" gorm:"unique;type:varchar(50);default:null" methods:"*post,*patch,*list" parse:"json"`
	Phone string `json:"phone" gorm:"type:varchar(30);unique;default:null" methods:"*post,*patch,*list" parse:"json"`

	Status uint `json:"status" methods:"*patch,*list" parse:"json"`

	Salt string `json:"-" gorm:"type:varchar(32)" methods:"post" parse:"json"`
	Code string `json:"-" gorm:"type:varchar(256)" methods:"post" parse:"json"`
}

type UserRole struct {
	BaseModel
	UserID string `json:"user_id" methods:"post,delete" parse:"path"`
	User   *User  `json:"user"`
	RoleID string `json:"role_id" methods:"post,delete" parse:"path"`
	Role   *Role  `json:"role"`
	Status string `json:"status" methods:"post,*patch,*list" parse:"json"`
}
