package apis

import (
	"net/http"
	"rinterest/middlewares"
	"rinterest/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// CommentAPI ...
type CommentAPI struct{}

// Register ...
func (co *CommentAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/comments", middlewares.JWT(), co.newone)
	rg.DELETE("/comments/:id", middlewares.JWT(), co.deleteone)
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
	if err := myDB.Omit(clause.Associations).Create(&comment).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    comment,
	})
}

func (co *CommentAPI) deleteone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query params",
		})
		return
	}

	comment := models.Comment{}
	if err = myDB.Where("id = ?", id).Find(&comment).Error; err != nil {
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

	if err := myDB.Delete(&comment).Error; err != nil {
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
			"error": "invalid query param",
		})
		return
	}

	comments := []models.Comment{}
	if err = myDB.Where("blog_id = ? AND invisibility = ?", id, "public").Order("created_at desc").Find(&comments).Error; err != nil {
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
			"error": "invalid query param",
		})
		return
	}

	comments := []models.Comment{}
	if err = myDB.Where("user_id = ? AND invisibility = ?", id, "public").Order("created_at desc").Find(&comments).Error; err != nil {
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
