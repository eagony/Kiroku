package apis

import (
	"net/http"
	"kiroku/middlewares"
	"kiroku/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CommentAPI ...
type CommentAPI struct{}

// Register ...
func (co *CommentAPI) Register(rg *gin.RouterGroup) {
	// 需要验证
	rg.POST("/comments", middlewares.JWT(), co.newone)
	rg.DELETE("/comments/:id", middlewares.JWT(), co.deleteone)

	// 无需验证
	rg.GET("/blogs/:id/comments", co.getallbyblogid)
	rg.GET("/users/:id/comments", co.getallbyuserid)
}

func (co *CommentAPI) newone(c *gin.Context) {
	comment := models.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// TODO: 自引用一对多的创建
	if err := db.Create(&comment).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
	})
}

func (co *CommentAPI) deleteone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param.",
		})
		return
	}

	comment := models.Comment{}
	if err = db.Where("id = ?", id).Find(&comment).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	if currentUserID != comment.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "怎么能擅自删除别人的评论呢！",
		})
		return
	}

	if err := db.Delete(&comment).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "succed",
	})
}

func (co *CommentAPI) getallbyblogid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param.",
		})
		return
	}

	comments := []models.Comment{}
	if err = db.Where("blog_id = ? AND invisibility = ?", id, "public").Order("created_at desc").Find(&comments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    comments,
	})
}

func (co *CommentAPI) getallbyuserid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param.",
		})
		return
	}

	comments := []models.Comment{}
	if err = db.Where("user_id = ? AND invisibility = ?", id, "public").Order("created_at desc").Find(&comments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    comments,
	})
}
