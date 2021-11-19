package models

type Message struct {
	BaseModel
	UserID   uint
	User     *User
	Title    string
	Redirect string
	Content  string
	From     string
}
