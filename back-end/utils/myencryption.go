package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

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
	return hashValue
}
