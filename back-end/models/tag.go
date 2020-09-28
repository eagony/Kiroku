package models

import "github.com/jinzhu/gorm"

// Tag 标签
type Tag struct {
	gorm.Model
	Icon   string `json:"icon" gorm:"type:varchar(32)"`
	Text   string `json:"text" gorm:"type:varchar(64)"`
	Color  string `json:"color" gorm:"type:varchar(32)"`
	UseFor string `json:"use_for" gorm:"type:enum('diary', 'blog'); default:'diary'"`

	// 多对多的反向引用，实现按标签筛选
	// Diaries []Diary `gorm:"many2many:diary_tags;"`
	// Blogs   []Blog  `gorm:"many2many:blog_tags;"`
}
