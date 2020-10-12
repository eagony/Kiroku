package apis

import (
	"net/http"
	"kiroku/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// StatisticAPI 统计数据API
type StatisticAPI struct{}

// Register ...
func (s *StatisticAPI) Register(rg *gin.RouterGroup) {
	rg.GET("/statistic/blogs/:id/views", s.blogviewsplus)
	rg.GET("/statistic/blogs/:id/likes", s.bloglikesplus)
}

// 统计博客访问量
func (s *StatisticAPI) blogviewsplus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	blog := models.Blog{}
	if err := db.First(&blog, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	blog.Views++
	db.Save(&blog)
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
	})
}

// 统计博客点赞数
func (s *StatisticAPI) bloglikesplus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	blog := models.Blog{}
	if err := db.First(&blog, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	blog.Likes++
	db.Save(&blog)
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
	})
}
