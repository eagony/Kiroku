package apis

import (
	"net/http"
	"rinterest/middlewares"
	"rinterest/models"

	"github.com/gin-gonic/gin"
)

// TagAPI ...
type TagAPI struct{}

// Register ...
func (t *TagAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/tags", middlewares.JWT(), t.newone)
	rg.GET("/tags", middlewares.JWT(), t.getall)
}

func (t *TagAPI) newone(c *gin.Context) {
	tag := models.Tag{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	currentUserRole := middlewares.GetUserRole()
	defer middlewares.ResetUserRole()
	if currentUserRole != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "非管理员无法添加标签！",
		})
		return
	}

	if err := myDB.Create(&tag).Error; err != nil {
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

func (t *TagAPI) getall(c *gin.Context) {
	// 暂时不分页
	// page, pageErr := strconv.Atoi(c.DefaultQuery("page", "1"))
	// if pageErr != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"Param error": pageErr.Error(),
	// 	})
	// 	return
	// }

	// perPage, perPageErr := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	// if perPageErr != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"Query error": perPageErr.Error(),
	// 	})
	// 	return
	// }

	useFor := c.DefaultQuery("use_for", "diary")
	var tags []models.Tag
	// extensions.MySQL().Limit(perPage).Offset((page - 1) * perPage).Order("created_at desc").Find(&data)
	myDB.Where("use_for = ?", useFor).Order("created_at desc").Find(&tags)
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    tags,
	})
}
