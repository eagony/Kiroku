package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *User) String() string {
	return fmt.Sprintf(
		"Username: %v\n,Email: %v\n",
		u.Username,
		u.Password)
}
func init() {
	RegisterSingleton("User", func() General {
		return new(User)
	})
}
