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

	// 用户代办列表，一对多
	ToDos []ToDo
	// 用户日记列表，一对多
	Diaries []Diary
}

func (u *User) String() string {
	return ""
}
func init() {
	RegisterSingleton("User", func() General {
		return new(User)
	})
}
