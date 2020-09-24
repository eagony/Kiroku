package apis

import (
	"net/http"
	"rinterest/extensions"
	"rinterest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BlogAPI 博客API
type BlogAPI struct{}

// Register ...
func (b *BlogAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/blogs", b.newone)
	rg.GET("/blogs/:id", b.getone)
	rg.GET("/publicblogs", b.getpublic)

	rg.GET("/users/:id/blogs", b.getallbyuserid)
}

func (b *BlogAPI) newone(c *gin.Context) {
	blog := models.Blog{}
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Bind error": err.Error(),
		})
		return
	}

	if err := extensions.MySQL().Create(&blog).Error; err != nil {
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
	if err = extensions.MySQL().Where("id = ?", id).Find(&blog).Error; err != nil {
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

func (b *BlogAPI) getallbyuserid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err.Error(),
		})
		return
	}

	bloges := []models.Blog{}
	if err = extensions.MySQL().Where("user_id = ?", id).Order("created_at desc").Find(&bloges).Error; err != nil {
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
	if err := extensions.MySQL().Where("invisibility = ?", "public").Order("created_at desc").Find(&bloges).Error; err != nil {
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
