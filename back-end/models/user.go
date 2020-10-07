package models

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"type:varchar(32)"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
	Role      string `json:"role" gorm:"type:enum('user', 'admin'); default:'user'"`
	Email     string `json:"email" gorm:"size:64"`
	Phone     string `json:"phone" gorm:"size:16"`
	Avatar    string `json:"avatar" gorm:"type:varchar(128)"`
	Signature string `json:"signature" gorm:"type:varchar(255)"`

	// 用户代办列表，一对多
	ToDos []ToDo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// 用户日记列表，一对多
	Diaries []Diary `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// 用户博客列表，一对多
	Blogs []Blog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// 用户评论列表，一对多
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
