package apis

import (
	"net/http"
	"rinterest/extensions"
	"rinterest/middlewares"
	"rinterest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DiaryAPI ...
type DiaryAPI struct{}

// Register ...
func (d *DiaryAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/diaries", d.newone)
	rg.GET("/diaries/:id", d.getone)
	rg.DELETE("/diaries/:id", d.deleteone)

	rg.GET("/users/:id/diaries", middlewares.JWT(), d.getallbyuserid)
}

func (d *DiaryAPI) newone(c *gin.Context) {
	diary := models.Diary{}
	if err := c.ShouldBindJSON(&diary); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Bind error": err.Error(),
		})
		return
	}

	if err := extensions.MySQL().Create(&diary).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Create error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": " create success",
		"data":    diary,
	})
}

func (d *DiaryAPI) deleteone(c *gin.Context) {
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil || id < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": idError,
		})
		return
	}

	if err := extensions.MySQL().Delete(&models.Diary{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Delete error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete succed.",
		"status":  "OK",
		"data":    "",
	})
}

func (d *DiaryAPI) getone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err.Error(),
		})
		return
	}

	diary := models.Diary{}
	if err = extensions.MySQL().Where("id = ?", id).Find(&diary).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    diary,
	})
}

func (d *DiaryAPI) getallbyuserid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err.Error(),
		})
		return
	}

	diaries := []models.Diary{}
	if err = extensions.MySQL().Where("user_id = ?", id).Order("created_at desc").Find(&diaries).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    diaries,
	})
}
