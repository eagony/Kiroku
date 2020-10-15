package middlewares

import (
	"kiroku/models"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// UserID 全局用户ID变量
var UserID uint

// UserRole 全局用户Role变量
var UserRole string

func init() {
	UserID = 0
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
}

// SecretKey ...
func SecretKey() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	}
}

// NewToken 生成Token
func NewToken(user *models.User) (tokens string, err error) {
	claim := jwt.MapClaims{
		"id":        user.ID,
		"role":      user.Role,
		"email":     user.Email,
		"phone":     user.Phone,
		"avatar":    user.Avatar,
		"username":  user.Username,
		"signature": user.Signature,
		"iat":       time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokens, err = token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return
}

// RefreshToken ...
func RefreshToken(token string) (tokens string, err error) {
	return
}

// JWT Auth middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if len(authorization) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "凭证无效，请重试。",
			})
			return
		}
		bearerToken := strings.Fields(authorization)[1]
		token, err := jwt.Parse(bearerToken, SecretKey())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "解析凭证失败，请重试。",
			})
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "cannot convert claim to mapclaim.",
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "凭证无效或被修改，请重试。",
			})
		}
		UserID = uint(claim["id"].(float64))
		UserRole = claim["role"].(string)
		c.Next()
	}
}

// GetUserID ...
func GetUserID() uint {
	return UserID
}

// ResetUserID ...
func ResetUserID() {
	UserID = 0
}

// GetUserRole ...
func GetUserRole() string {
	return UserRole
}

// ResetUserRole ...
func ResetUserRole() {
	UserRole = "user"
}
