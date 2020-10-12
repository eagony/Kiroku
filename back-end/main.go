package main

import (
	"kiroku/apis"
	"kiroku/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
}

func main() {
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	// 开启跨域访问
	router.Use(middlewares.Cors())
	// 路由分组
	apiV1 := router.Group("api/v1")
	// 文件服务
	apiV1.StaticFS("/data/images", http.Dir("data/images"))
	// 注册路由到接口
	new(apis.UserAPI).Register(apiV1)
	new(apis.ToDoAPI).Register(apiV1)
	new(apis.TagAPI).Register(apiV1)
	new(apis.DiaryAPI).Register(apiV1)
	new(apis.BlogAPI).Register(apiV1)
	new(apis.ImageAPI).Register(apiV1)
	new(apis.StatisticAPI).Register(apiV1)
	new(apis.CommentAPI).Register(apiV1)

	port := os.Getenv("PORT")

	if os.Getenv("SSL") == "ON" {
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{}
		// Generated using sh generate-certificate.sh
		SSLKeys.CERT = "./cert/MyCA.cer"
		SSLKeys.KEY = "./cert/MyCA.key"
		router.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		router.Run(":" + port)
	}
}
