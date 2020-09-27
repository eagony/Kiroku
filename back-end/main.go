package main

import (
	"log"
	"net/http"
	"os"
	"rinterest/apis"
	"rinterest/middlewares"

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
	router.Use(middlewares.Cors())

	apiV1 := router.Group("api/v1")
	apiV1.StaticFS("/data/images", http.Dir("data/images"))
	new(apis.UserAPI).Register(apiV1)
	new(apis.ToDoAPI).Register(apiV1)
	new(apis.TagAPI).Register(apiV1)
	new(apis.DiaryAPI).Register(apiV1)
	new(apis.BlogAPI).Register(apiV1)
	new(apis.ImageAPI).Register(apiV1)

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
