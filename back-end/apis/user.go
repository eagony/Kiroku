package apis

import (
	"net/http"
	"rinterest/extensions"
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
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err.Error(),
		})
		return
	}

	user := models.User{}
	if err = myDB.Where("id = ?", id).Find(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    user,
	})
}

func (u *UserAPI) updateone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{}

	if err = myDB.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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
			"Bind error": err.Error(),
		})
		return
	}
	// 密码加密
	user.Password = utils.EncryptPassword([]byte(user.Password))

	if err := myDB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Create error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": " create success",
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
			"Bind error": err.Error(),
		})
		return
	}

	user := models.User{}
	extensions.MySQL().Where("username = ?", loginForm.Username).Find(&user)

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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "生成凭证失败.",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   token,
	})
}

// 批量查询用户的一对多模型

func (u *UserAPI) todolist(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err.Error(),
		})
		return
	}

	user := models.User{}

	extensions.MySQL().Preload("ToDos").Where("id = ?", userID).Find(&user)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    user.ToDos,
	})
}

func (u *UserAPI) diarylist(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Convert id error": err.Error(),
		})
		return
	}

	user := models.User{}

	myDB.Order("created_at desc").Preload("Diaries").Where("id = ?", userID).Find(&user)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    user.Diaries,
	})
}
