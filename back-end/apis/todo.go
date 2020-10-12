package apis

import (
	"net/http"
	"kiroku/middlewares"
	"kiroku/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ToDoAPI ...
type ToDoAPI struct{}

// Register ...
func (t *ToDoAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/todos", middlewares.JWT(), t.newone)
	rg.PUT("/todos/:id", middlewares.JWT(), t.updateone)
	rg.DELETE("/todos/:id", middlewares.JWT(), t.deleteone)
	rg.GET("/users/:id/todos", middlewares.JWT(), t.getallbyuserid)
}

func (t *ToDoAPI) newone(c *gin.Context) {
	todo := models.ToDo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Create(&todo).Error; err != nil {
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

func (t *ToDoAPI) deleteone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param.",
		})
		return
	}

	todo := models.ToDo{}
	if err = db.Where("id = ?", id).Find(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	if currentUserID != todo.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "怎么能擅自删除别人的日记呢,调皮！",
		})
		return
	}

	if err := db.Delete(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "删除成功。",
	})
}

func (t *ToDoAPI) updateone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param.",
		})
		return
	}

	todo := models.ToDo{}
	if err = db.First(&todo, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	if currentUserID != todo.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "怎么能擅自修改别人的代办呢！",
		})
		return
	}

	if err = c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.Save(&todo)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "更新成功",
	})
}

func (t *ToDoAPI) getallbyuserid(c *gin.Context) {
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
			"error": "不能偷看别人的代办列表哦！",
		})
		return
	}

	todos := []models.ToDo{}
	if err = db.Where("user_id = ?", id).Order("created_at asc").Find(&todos).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    todos,
	})
}
