package models

import "github.com/jinzhu/gorm"

// Tag 标签
type Tag struct {
	gorm.Model
	Icon  string `json:"icon" gorm:"type:varchar(32)"`
	Text  string `json:"text" gorm:"type:varchar(64)"`
	Color string `json:"color" gorm:"type:varchar(32)"`
}
