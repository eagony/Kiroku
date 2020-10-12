package apis

import (
	"net/http"
	"kiroku/middlewares"
	"kiroku/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DiaryAPI ...
type DiaryAPI struct{}

// Register ...
func (d *DiaryAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/diaries", middlewares.JWT(), d.newone)
	rg.GET("/diaries/:id", middlewares.JWT(), d.getone)
	rg.DELETE("/diaries/:id", middlewares.JWT(), d.deleteone)
	rg.GET("/users/:id/diaries", middlewares.JWT(), d.getallbyuserid)
}

func (d *DiaryAPI) newone(c *gin.Context) {
	diary := models.Diary{}
	if err := c.ShouldBindJSON(&diary); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Create(&diary).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    diary,
	})
}

func (d *DiaryAPI) deleteone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param.",
		})
		return
	}

	diary := models.Diary{}
	if err = db.Where("id = ?", id).Find(&diary).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	if currentUserID != diary.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "怎么能擅自删除别人的日记呢！",
		})
		return
	}

	if err := db.Delete(&diary).Error; err != nil {
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

func (d *DiaryAPI) getone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param.",
		})
		return
	}

	diary := models.Diary{}
	if err = db.Where("id = ?", id).Find(&diary).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	if diary.Invisibility == "private" && currentUserID != diary.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "不能偷看别人的私密日记哦！",
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
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param.",
		})
		return
	}

	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	if currentUserID != uint(id) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "不能偷看别人的日记列表哦！",
		})
		return
	}

	diaries := []models.Diary{}
	if err = db.Preload("Tags").Where("user_id = ?", id).Order("created_at desc").Find(&diaries).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    diaries,
	})
}

func (d *DiaryAPI) getpublic(c *gin.Context) {
	diaries := []models.Diary{}
	if err := db.Preload("Tags").Where("invisibility = ?", "public").Order("created_at desc").Find(&diaries).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    diaries,
	})
}
