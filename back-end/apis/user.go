package apis

import (
	"net/http"
	"rinterest/middlewares"
	"rinterest/models"
	"rinterest/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserAPI 用户API
type UserAPI struct{}

// Register ...
func (u *UserAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/users/register", u.register)
	rg.POST("/users/login", u.login)
	rg.GET("/users/:id", middlewares.JWT(), u.getone)
	rg.PUT("/users/:id", middlewares.JWT(), u.updateone)
}

func (u *UserAPI) getone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param",
		})
		return
	}

	user := models.User{}
	if err = myDB.Where("id = ?", id).Find(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Password = ""
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    user,
	})
}

func (u *UserAPI) updateone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param",
		})
		return
	}

	user := models.User{}
	if err = myDB.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	if currentUserID != user.ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "怎么能擅自更改别人的信息呢！",
		})
		return
	}

	if err = c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	myDB.Save(&user)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "资料更新成功。",
	})
}

func (u *UserAPI) register(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 密码加密
	user.Password = utils.EncryptPassword([]byte(user.Password))

	if err := myDB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = ""
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
		"data":    user,
	})
}

func (u *UserAPI) login(c *gin.Context) {
	loginForm := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{}
	myDB.Where("username = ?", loginForm.Username).Find(&user)

	if user.Username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "无此用户。",
		})
		return
	}

	if !utils.ComparePassword([]byte(user.Password), []byte(loginForm.Password)) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "密码错误。",
		})
		return
	}
	token, err := middlewares.NewToken(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "生成凭证失败.",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   token,
	})
}
