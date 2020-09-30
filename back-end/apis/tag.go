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
	// 后面限制管理员才能使用增加接口
	rg.POST("/tags", t.newone)
	rg.GET("/tags", middlewares.JWT(), t.getall)
}

func (t *TagAPI) newone(c *gin.Context) {
	tag := models.Tag{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Bind error": err.Error(),
		})
		return
	}

	if err := myDB.Create(&tag).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Create error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": " create success",
		"data":    tag,
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
