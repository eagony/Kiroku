package models

import (
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`

	// One to Many
	ToDos []ToDo
}

func (u *User) String() string {
	return ""
}
func init() {
	RegisterSingleton("User", func() General {
		return new(User)
	})
}
