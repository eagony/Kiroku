package extensions

import (
	"fmt"
	"log"
	"os"
	"rinterest/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysql *gorm.DB

func init() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
	// conn, err := gorm.Open("mysql", "root:123456@/rinterest?charset=utf8mb4&parseTime=True&loc=Local")
	conn, err := gorm.Open(os.Getenv("DB_TYPE"),
		fmt.Sprintf("%s:%s@(%s)/%s?%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_OPTIONS")))
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
