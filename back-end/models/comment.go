package models

import "gorm.io/gorm"

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string `json:"content" gorm:"type:varchar(1024)"`
	// 评论是否已读，默认否
	Readed bool `json:"readed" gorm:"not null; default:false;"`
	// 评论是否被屏蔽，默认否
	Disabled bool `json:"disabled" gorm:"not null; default:false;"`

	// 外键，评论者id
	UserID uint `json:"user_id"`
	// 外键，被评论的博客id
	BlogID uint `json:"blog_id"`
	// 外键，自引用的评论id
	ParentID uint `json:"parent_id"`
	// 自引用的一对多
	Children []Comment `json:"children" gorm:"many2many:comment_children;foreignKey:ParentID;references:ParentID;"`
}

func (c Comment) getdescendants() []Comment {
	// 获取一级评论的所有子孙评论
	comments := []Comment{}
	var descendants func(c Comment)
	descendants = func(c Comment) {
		if c.Children != nil {

		}
	}
	descendants(c)
	return comments
}
