package models

type Role struct {
	BaseModel
	Name      string `json:"name" methods:"post,put,*patch,*list" parse:"json"`
	Des       string `json:"des" methods:"post,put,*patch" parse:"json"`
	AppID     string `json:"app_id" gorm:"index;type:varchar(32)" methods:"post,put,*patch" parse:"json"`
	App       *App   `json:"app" gorm:"foreignKey:ID;references:AppID"`
	UserCount uint   `json:"user_count"`
}
