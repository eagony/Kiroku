package apis

import (
	"rinterest/extensions"
	"rinterest/models"

	"gorm.io/gorm"
)

var myDB *gorm.DB

func init() {
	myDB = extensions.MySQL()
	// myDB = extensions.Postgresql()
	myDB.AutoMigrate(&models.User{}, &models.ToDo{}, &models.Tag{}, &models.Diary{}, &models.Blog{}, &models.Comment{})
}
