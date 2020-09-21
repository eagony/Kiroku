package main

import (
	"github.com/gin-gonic/gin"
	"rinterest/apis"
)

func init() {

}

func main() {
	r := gin.Default()
	apiV1 := r.Group("api/v1")

	new(apis.UserAPI).Register(apiV1)
	new(apis.ToDoAPI).Register(apiV1)

	r.Run(":8000") // 监听并在 0.0.0.0:8080 上启动服务
}
