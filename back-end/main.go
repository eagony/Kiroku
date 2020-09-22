package main

import (
	"rinterest/apis"
	"rinterest/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	router := gin.Default()
	router.Use(middlewares.Cors())

	apiV1 := router.Group("api/v1")
	new(apis.UserAPI).Register(apiV1)
	new(apis.ToDoAPI).Register(apiV1)
	new(apis.TagAPI).Register(apiV1)
	new(apis.DiaryAPI).Register(apiV1)

	router.Run(":8000") // 监听并在 0.0.0.0:8080 上启动服务
}
