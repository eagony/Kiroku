package apis

import (
	"net/http"
	"rinterest/extensions"
	"rinterest/models"

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
