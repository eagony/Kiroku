package middlewares

import (
	"net/http"
	"rinterest/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// NewToken 生成Token
func NewToken(user *models.User) (tokens string, err error) {
	claim := jwt.MapClaims{
		"id":        user.ID,
		"Role":      user.Role,
		"Email":     user.Email,
		"Phone":     user.Phone,
		"Avatar":    user.Avatar,
		"Username":  user.Username,
		"Signature": user.Signature,
		"iat":       time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokens, err = token.SignedString([]byte("you-will-never-guess"))
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
				"message": "Bad authorization.",
			})
			return
		}
		bearerToken := strings.Fields(authorization)[1]
		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("you-will-never-guess"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "parse token failed.",
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token.",
			})
		}
		c.Next()
	}
}
