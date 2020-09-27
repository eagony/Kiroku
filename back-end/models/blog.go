package models

import "github.com/jinzhu/gorm"

// Blog 博客模型
type Blog struct {
	gorm.Model
	// 标题
	Title string `json:"title" gorm:"type:varchar(256)"`
	// 摘要
	Summary string `json:"summary" gorm:"type:varchar(512)"`
	// 正文
	Content string `json:"content" gorm:"type:varchar(10240)"`
	// 可见性，默认为public
	Invisibility string `json:"invisibility" gorm:"type:enum('public', 'private', 'protected'); default:'public'"`

	// 外键
	UserID uint `json:"user_id"`
	// 标签，一对多
	Tags []Tag `gorm:"many2many:blog_tags;"`
}
