package extensions

import (
	"rinterest/models"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysql *gorm.DB

func init() {
	conn, err := gorm.Open("mysql", "root:123456@/rinterest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("Connect to MySQL has failed.")
	}
	mysql = conn
	mysql.Debug().AutoMigrate(&models.User{}, &models.ToDo{}, &models.Tag{}, &models.Diary{}, &models.Blog{})
}

// MySQL ...
func MySQL() *gorm.DB {
	return mysql
}
