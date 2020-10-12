package apis

import (
	"kiroku/extensions"
	"kiroku/models"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// 启用MySQL数据库
	db = extensions.MySQL()
	// 启用PostgreSQL数据库
	// db = extensions.Postgresql()
	db.AutoMigrate(&models.User{}, &models.ToDo{}, &models.Tag{}, &models.Diary{}, &models.Blog{}, &models.Comment{})
}
