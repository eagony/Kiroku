package models

import "github.com/jinzhu/gorm"

// Diary 日记模型
type Diary struct {
	gorm.Model
	// 标签
	Tags string `json:"tags" gorm:"type:varchar(256)"`
	// 日记内容
	Content string `json:"content" gorm:"type:varchar(4096)"`
	// 用户外键
	UserID uint `json:"user_id"`
}
