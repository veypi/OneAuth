package models

type Message struct {
	BaseModel
	UserID   uint   `json:"user_id"`
	User     *User  `json:"user"`
	Title    string `json:"title"`
	Redirect string `json:"redirect"`
	Content  string `json:"content"`
	From     string `json:"from"`
}
