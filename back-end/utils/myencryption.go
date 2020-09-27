package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// MD5FileName 对文件名进行hash
func MD5FileName(name string) string {
	m := md5.New()
	m.Write([]byte(name))

	return hex.EncodeToString(m.Sum(nil))
}

// FileHash 计算文件hash
func FileHash(file *os.File) (hashValue string) {
	hash := sha256.New()
	_, err := io.Copy(hash, file)
	if err != nil {
		panic(err.Error())
	} else {
		sum := hash.Sum(nil)
		hashValue = hex.EncodeToString(sum)
	}
	return
}

// EncryptPassword 密码加密
func EncryptPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePassword 比对密码
func ComparePassword(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
