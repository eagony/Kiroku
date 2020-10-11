package models

import "gorm.io/gorm"

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
	// 统计数据
	Views uint `json:"views" gorm:"size:8; not null; default:0"`
	Likes uint `json:"likes" gorm:"size:8; not null; default:0"`

	// 外键
	UserID uint `json:"user_id"`
	// 标签，一对多
	Tags []Tag `json:"tags" gorm:"many2many:blog_tags;"`
	// 博客评论列表，一对多
	Comments []Comment `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
