package utils

import "os"

// IsDirExists 检查文件或者目录是否存在
func IsDirExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
