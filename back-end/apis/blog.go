package apis

import (
	"net/http"
	"rinterest/middlewares"
	"rinterest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BlogAPI 博客API
type BlogAPI struct{}

// Register ...
func (b *BlogAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/blogs", middlewares.JWT(), b.newone)
	rg.GET("/blogs/:id", b.getone)
	rg.GET("/publicblogs", b.getpublic)
	rg.DELETE("/blogs/:id", middlewares.JWT(), b.deleteone)

	rg.GET("/users/:id/blogs", middlewares.JWT(), b.getallbyuserid)
}

func (b *BlogAPI) newone(c *gin.Context) {
	blog := models.Blog{}
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := myDB.Create(&blog).Error; err != nil {
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

func (b *BlogAPI) getone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	blog := models.Blog{}
	if err = myDB.Where("id = ?", id).Find(&blog).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()

	if blog.Invisibility == "private" && currentUserID != blog.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "不能偷看别人的私密博客哦！",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    blog,
	})
}

func (b *BlogAPI) deleteone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query",
		})
		return
	}

	blog := models.Blog{}
	if err = myDB.Where("id = ?", id).Find(&blog).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	if currentUserID != blog.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "怎么能擅自删除别人的博客呢！",
		})
		return
	}
	if err := myDB.Delete(&blog).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "博客删除成功",
	})
}

func (b *BlogAPI) getallbyuserid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param",
		})
		return
	}
	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	if currentUserID != uint(id) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "不能偷看别人的博客列表哦！",
		})
		return
	}
	bloges := []models.Blog{}
	if err = myDB.Preload("Tags").Preload("Comments").Where("user_id = ?", id).Order("created_at desc").Find(&bloges).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    bloges,
	})
}

func (b *BlogAPI) getpublic(c *gin.Context) {
	bloges := []models.Blog{}
	if err := myDB.Preload("Tags").Preload("Comments").Where("invisibility = ?", "public").Order("created_at desc").Find(&bloges).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    bloges,
	})
}
