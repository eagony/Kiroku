package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rinterest/extensions"
	"rinterest/models"
	"strconv"
)

type GeneralAPI struct{}

// CreateOne 新建一个
func (g *GeneralAPI) CreateOne(insName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ins := models.Create(insName)
		if err := c.ShouldBindJSON(&ins); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Bind error": err.Error(),
			})
			return
		}
		tx := extensions.MySQL().Begin()
		if err := tx.Create(ins).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Create error": err.Error(),
			})
			return
		}
		tx.Commit()

		// handle success
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  200,
			"message": insName + " create success",
			"data":    ins,
		})
		ins = nil
	}
}

func (g *GeneralAPI) DeleteOne(insName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, idErr := strconv.Atoi(c.Param("id"))
		if idErr != nil || id < 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Convert id error": idErr,
			})
			return
		}

		tx := extensions.MySQL().Begin()

		ins := models.Create(insName)
		tx.First(ins, id)
		tx.Delete(ins)
		tx.Commit()

		c.JSON(http.StatusOK, gin.H{
			"message": "Delete success.",
			"user":    ins.String(),
		})
	}
}

func (g *GeneralAPI) UpdateOne(insName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Convert id error": err.Error(),
			})
			return
		}
		ins := models.Create(insName)

		tx := extensions.MySQL().Begin()

		if err = tx.First(ins, id).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Query error": err.Error(),
			})
			return
		}

		if err = c.ShouldBindJSON(ins); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Bind error": err.Error(),
			})
			return
		}

		tx.Save(ins)
		tx.Commit()

		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "success",
			"data":    ins.String(),
		})
	}
}

func (g *GeneralAPI) GetOne(insName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Convert id error": err.Error(),
			})
			return
		}
		ins := models.Create(insName)

		tx := extensions.MySQL().Begin()

		if err = tx.First(ins, id).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Query error": err.Error(),
			})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "success",
			"data":    ins.String(),
		})
	}
}

func (g *GeneralAPI) GetAll(insName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数
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

		// 查询事务
		tx := extensions.MySQL().Begin()

		var data []models.User
		tx.Limit(perPage).Offset((page - 1) * perPage).Order("created_at desc").Find(&data)

		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   data,
			"total":  len(data),
		})
	}
}
