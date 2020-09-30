package apis

import (
	"net/http"
	"rinterest/middlewares"
	"rinterest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ToDoAPI ...
type ToDoAPI struct{}

// Register ...
func (t *ToDoAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/todos", middlewares.JWT(), t.newone)
	rg.DELETE("/todos/:id", middlewares.JWT(), t.deleteone)
	rg.PUT("/todos/:id", middlewares.JWT(), t.updateone)

	rg.GET("/users/:id/todos", middlewares.JWT(), t.getallbyuserid)
}

func (t *ToDoAPI) newone(c *gin.Context) {
	todo := models.ToDo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Bind error": err.Error(),
		})
		return
	}

	if err := myDB.Create(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Create error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": " create success",
		"data":    todo,
	})
}

func (t *ToDoAPI) deleteone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err,
		})
		return
	}

	if err := myDB.Delete(&models.ToDo{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Delete error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功。",
		"status":  "OK",
	})
}

func (t *ToDoAPI) updateone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	todo := models.ToDo{}

	if err = myDB.First(&todo, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	myDB.Save(&todo)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "更新成功",
	})
}

func (t *ToDoAPI) getallbyuserid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err.Error(),
		})
		return
	}

	todos := []models.ToDo{}
	if err = myDB.Where("user_id = ?", id).Order("created_at asc").Find(&todos).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    todos,
	})
}
