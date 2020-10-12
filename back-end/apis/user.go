package apis

import (
	"fmt"
	"net/http"
	"kiroku/middlewares"
	"kiroku/models"
	"kiroku/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserAPI 用户API
type UserAPI struct{}

// Register ...
func (u *UserAPI) Register(rg *gin.RouterGroup) {
	// 需要验证
	rg.PUT("/users/:id", middlewares.JWT(), u.updateone)

	// 无需验证
	rg.POST("/users/register", u.register)
	rg.POST("/users/login", u.login)
	rg.GET("/users/:id", u.getone)
}

func (u *UserAPI) getone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param.",
		})
		return
	}

	user := models.User{}
	if err = db.Where("id = ?", id).Find(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data := models.User{}
	data.Avatar = user.Avatar
	data.Username = user.Username
	data.Signature = user.Signature
	data.Email = user.Email
	data.Phone = user.Phone

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    data,
	})
}

func (u *UserAPI) updateone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid query param.",
		})
		return
	}

	user := models.User{}
	if err = db.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	currentUserID := middlewares.GetUserID()
	defer middlewares.ResetUserID()
	if currentUserID != user.ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "怎么能擅自修改别人的个人资料呢！",
		})
		return
	}

	newUser := models.User{}
	if err = c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 根据 `struct` 更新属性，只会更新非零值的字段
	db.Model(&user).Updates(&newUser)

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
	// 对密码进行哈希加密
	user.Password = utils.EncryptPassword([]byte(user.Password))

	if err := db.Create(&user).Error; err != nil {
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
	if err := db.Where("username = ?", loginForm.Username).Find(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("user->", user.Username)
	if user.Username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "用户名不存在，请重试。",
		})
		fmt.Println("用户名不存在，请重试。")
		return
	}

	if !utils.ComparePassword([]byte(user.Password), []byte(loginForm.Password)) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "密码错误，请重试。",
		})
		fmt.Println("密码错误，请重试。")
		return
	}
	token, err := middlewares.NewToken(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "生成凭证失败。",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   token,
	})
}
