package apis

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"rinterest/middlewares"
	"rinterest/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
}

// ImageAPI 图片上传API
type ImageAPI struct{}

// Register ...
func (i *ImageAPI) Register(rg *gin.RouterGroup) {
	rg.POST("/images", middlewares.JWT(), i.upload)
}

func (i *ImageAPI) upload(c *gin.Context) {
	img, err := c.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 检查文件后缀是否为图片类型
	fileExt := strings.ToLower(path.Ext(img.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "上传失败!只允许png,jpg,gif,jpeg文件",
		})
		return
	}
	// 加盐算出图片新的文件名
	imgHashName := utils.MD5FileName(fmt.Sprintf("%s%s", img.Filename, time.Now().String()))
	fmt.Println(imgHashName)
	imgDir := os.Getenv("BASE_IMAGE_DIR")
	if !utils.IsDirExists(imgDir) {
		os.MkdirAll(imgDir, os.ModePerm)
	}
	imgPath := fmt.Sprintf("%s/%s%s", imgDir, imgHashName, fileExt)
	fmt.Println(imgPath)
	c.SaveUploadedFile(img, imgPath)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "图片上传成功。",
		"uri":     imgPath,
	})
}
