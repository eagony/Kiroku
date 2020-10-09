package apis

import (
	"rinterest/extensions"
	"rinterest/models"

	"gorm.io/gorm"
)

var myDB *gorm.DB

func init() {
	// 启用MySQL数据库
	myDB = extensions.MySQL()
	// 启用PostgreSQL数据库
	// myDB = extensions.Postgresql()
	myDB.AutoMigrate(&models.User{}, &models.ToDo{}, &models.Tag{}, &models.Diary{}, &models.Blog{}, &models.Comment{})
}
