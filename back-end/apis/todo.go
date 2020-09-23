package apis

import (
	"net/http"
	"rinterest/extensions"
	"rinterest/middlewares"
	"rinterest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ToDoAPI ...
type ToDoAPI struct {
	G GeneralAPI
}

// Register ...
func (t *ToDoAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/todos", t.newone)
	rg.DELETE("/todos/:id", t.G.DeleteOne("ToDo"))
	rg.PUT("/todos/:id", t.G.UpdateOne("ToDo"))
	rg.GET("/todos/:id", t.G.GetOne("ToDo"))

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

	if err := extensions.MySQL().Create(&todo).Error; err != nil {
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

func (t *ToDoAPI) getallbyuserid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err.Error(),
		})
		return
	}

	todos := []models.ToDo{}
	if err = extensions.MySQL().Where("user_id = ?", id).Order("created_at asc").Find(&todos).Error; err != nil {
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
