package models

import "gorm.io/gorm"

// Diary 日记模型
type Diary struct {
	gorm.Model
	// 内容
	Content string `json:"content" gorm:"type:varchar(4096)"`
	// 可见性，默认为private、
	Invisibility string `json:"invisibility" gorm:"type:enum('public', 'private', 'protected'); default:'private'"`

	// 外键
	UserID uint `json:"user_id"`
	// 标签，多对多
	Tags []Tag `json:"tags" gorm:"many2many:diary_tags;"`
}
