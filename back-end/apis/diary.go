package apis

import (
	"fmt"
	"net/http"
	"rinterest/middlewares"
	"rinterest/models"
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
			"Bind error": err.Error(),
		})
		return
	}
	if err := myDB.Create(&diary).Error; err != nil {
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

	if err := myDB.Delete(&models.Diary{}, id).Error; err != nil {
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
	if err = myDB.Preload("Tags").Where("id = ?", id).Find(&diary).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}
	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	fmt.Println(currentUserID, diary.UserID)
	if diary.UserID != currentUserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "偷看别人的日记是不好的哦！",
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
	if err = myDB.Preload("Tags").Where("user_id = ?", id).Order("created_at desc").Find(&diaries).Error; err != nil {
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
