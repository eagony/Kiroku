package apis

import "github.com/gin-gonic/gin"

type UserAPI struct {
	G GeneralAPI
}

func (u *UserAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/users", new(GeneralAPI).CreateOne("User"))
	rg.DELETE("/users/:id", u.G.DeleteOne("User"))
	rg.PUT("/users/:id", u.G.UpdateOne("User"))
	rg.GET("/users/:id", u.G.GetOne("User"))
	rg.GET("/users", u.G.GetAll("User"))
}
