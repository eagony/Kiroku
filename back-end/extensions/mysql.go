package extensions

import (
	"fmt"
	"rinterest/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysql *gorm.DB

func init() {
	conn, err := gorm.Open("mysql", "root:kaguya@/test_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("Open MySQL failed.。")
	}
	fmt.Println("Open MySQL succeed.")
	mysql = conn
	mysql.Debug().AutoMigrate(&models.User{}, models.ToDo{})
}

func MySQL() *gorm.DB {
	return mysql
}
