package extensions

import (
	"fmt"
	"rinterest/models"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysql *gorm.DB

func init() {
	conn, err := gorm.Open("mysql", "root:123456@/rinterest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("Open MySQL failed.。")
	}
	fmt.Println("Open MySQL succeed.")
	mysql = conn
	mysql.Debug().AutoMigrate(&models.User{}, models.ToDo{})
}

// MySQL ...
func MySQL() *gorm.DB {
	return mysql
}
