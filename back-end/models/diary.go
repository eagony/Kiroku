package models

import "github.com/jinzhu/gorm"

// Diary 日记模型
type Diary struct {
	gorm.Model
	// 内容
	Content string `json:"content" gorm:"type:varchar(4096)"`

	// 外键
	UserID uint `json:"user_id"`
	// 标签，一对多
	Tags []Tag `gorm:"many2many:diary_tags;"`
}
