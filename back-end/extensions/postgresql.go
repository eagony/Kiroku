package extensions

import (
	"gorm.io/gorm"
)

var postgresqlConn *gorm.DB

// func init() {
// 	// 加载环境变量
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file, please create one in the root directory")
// 	}

// 	conn, err := gorm.Open(postgres.New(postgres.Config{
// 		PreferSimpleProtocol: true, // disables implicit prepared statement usage
// 		DSN:                  "host=pgm-bp12x1o674j31212io.pg.rds.aliyuncs.com user=chika password=a_strong_password dbname=postgres port=1921 sslmode=disable TimeZone=Asia/Shanghai",
// 	}), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("Connect to postgres failed.")
// 	}
// 	postgresqlConn = conn
// }

// Postgresql ...
func Postgresql() *gorm.DB {
	return postgresqlConn
}
