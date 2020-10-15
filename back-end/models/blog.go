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
	Content string `json:"content" gorm:"type:text;"`
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

// BlogPreview 博客预览模型，不需要博客内容等数据，优化传输速度
type BlogPreview struct {
	ID      uint
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Views   uint   `json:"views"`
	Likes   uint   `json:"likes"`
	UserID  uint   `json:"user_id"`
	Tags    []Tag  `json:"tags" gorm:"many2many:blog_tags;"`
	// Comments []Comment `json:"comments" gorm:"foreignKey:BlogID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
