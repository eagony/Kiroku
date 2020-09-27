package models

import (
	"github.com/jinzhu/gorm"
)

// ToDo 代办模型
type ToDo struct {
	gorm.Model
	Text string `json:"text" gorm:"type:varchar(1024)"`
	Done bool   `json:"done"`

	// 外键
	UserID uint `json:"user_id"`
}

// TableName 指定表名
func (td *ToDo) TableName() string {
	return "todos"
}
