package models

import "github.com/jinzhu/gorm"

// Blog 博客模型
type Blog struct {
	gorm.Model
	// 标签
	Tags string `json:"tags" gorm:"type:varchar(256)"`
	// 标题
	Title string `json:"title" gorm:"type:varchar(256)"`
	// 摘要
	Summary string `json:"summary" gorm:"type:varchar(512)"`
	// 正文
	Content string `json:"content" gorm:"type:varchar(10240)"`
	// 可见性，默认为public
	Invisibility string `json:"invisibility" gorm:"type:enum('public', 'private', 'protected'); default:'public'"`

	// 用户外键
	UserID uint `json:"user_id"`
}
