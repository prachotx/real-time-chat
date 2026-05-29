package model

type User struct {
	GormModel
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
