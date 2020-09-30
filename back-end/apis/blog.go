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
	rg.GET("/blogs/:id", middlewares.JWT(), b.getone)
	rg.GET("/publicblogs", b.getpublic)
	rg.DELETE("/blogs/:id", middlewares.JWT(), b.deleteone)

	rg.GET("/users/:id/blogs", middlewares.JWT(), b.getallbyuserid)
}

func (b *BlogAPI) newone(c *gin.Context) {
	blog := models.Blog{}
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Bind error": err.Error(),
		})
		return
	}

	if err := myDB.Create(&blog).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Create error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": " create success",
		"data":    blog,
	})
}

func (b *BlogAPI) getone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err.Error(),
		})
		return
	}

	blog := models.Blog{}
	if err = myDB.Where("id = ?", id).Find(&blog).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
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
	if err != nil || id < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := myDB.Delete(&models.Blog{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (b *BlogAPI) getallbyuserid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err.Error(),
		})
		return
	}

	bloges := []models.Blog{}
	if err = myDB.Preload("Tags").Where("user_id = ?", id).Order("created_at desc").Find(&bloges).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
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
	if err := myDB.Preload("Tags").Where("invisibility = ?", "public").Order("created_at desc").Find(&bloges).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    bloges,
	})
}
