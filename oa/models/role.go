package models

type Role struct {
	BaseModel
	Name      string    `json:"name" methods:"post,*patch,*list" parse:"json"`
	Des       string    `json:"des" methods:"post,*patch" parse:"json"`
	AppID     string    `json:"app_id" gorm:"index;type:varchar(32)" methods:"post,*patch" parse:"json"`
	App       *App      `json:"-" gorm:"foreignKey:AppID;references:ID"`
	UserCount uint      `json:"user_count"`
	Access    []*Access `json:"-"`
}
