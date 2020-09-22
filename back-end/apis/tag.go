package apis

import (
	"net/http"
	"rinterest/extensions"
	"rinterest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TagAPI ...
type TagAPI struct{}

// Register ...
func (t *TagAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/tags", t.newone)
	rg.GET("/tags", t.getall)
}

func (t *TagAPI) newone(c *gin.Context) {
	tag := models.Tag{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Bind error": err.Error(),
		})
		return
	}

	if err := extensions.MySQL().Create(&tag).Error; err != nil {
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
	page, pageErr := strconv.Atoi(c.DefaultQuery("page", "1"))
	if pageErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Param error": pageErr.Error(),
		})
		return
	}

	perPage, perPageErr := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if perPageErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": perPageErr.Error(),
		})
		return
	}

	var data []models.Tag
	extensions.MySQL().Limit(perPage).Offset((page - 1) * perPage).Order("created_at desc").Find(&data)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
		"total":  len(data),
	})
}
