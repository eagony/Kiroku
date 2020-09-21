package models

import (
	"github.com/jinzhu/gorm"
)

// ToDo 代办模型
type ToDo struct {
	gorm.Model
	Text string `json:"text" gorm:"type:varchar(1024)"`
	Done bool   `json:"done"`

	// Foreignkey
	UserID uint `json:"user_id"`
}

func (td *ToDo) String() string {
	return ""
}

// TableName ...
func (td ToDo) TableName() string {
	return "todos"
}

func init() {
	RegisterSingleton("ToDo", func() General {
		return new(ToDo)
	})
}
