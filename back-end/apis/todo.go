package apis

import "github.com/gin-gonic/gin"

type ToDoAPI struct {
	G GeneralAPI
}

func (td *ToDoAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/todos", td.G.CreateOne("ToDo"))
	rg.DELETE("/todos/:id", td.G.DeleteOne("ToDo"))
	rg.PUT("/todos/:id", td.G.UpdateOne("ToDo"))
	rg.GET("/todos/:id", td.G.GetOne("ToDo"))
}
