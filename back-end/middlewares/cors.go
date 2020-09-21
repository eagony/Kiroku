package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 开启跨域函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置返回格式是json
		c.Set("content-type", "application/json")
		// 允许访问所有域
		c.Header("Access-Control-Allow-Origin", "*")
		// 跨域请求是否需要带cookie信息
		c.Header("Access-Control-Allow-Credentials", "true")
		// 支持的跨域请求方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		// header 类型
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, AccessToken, X-CSRF-Token, Authorization, Token, origin")
		// 响应报头指示哪些报头可以公开为通过列出他们的名字的响应的一部分
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		//放行所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
