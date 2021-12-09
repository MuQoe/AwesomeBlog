package tools

import (
	global "AwesomeBlog/globals"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

// Encrypt 密码加密函数
func Encrypt(password string) string {
	dk := pbkdf2.Key([]byte(password), []byte(global.AppSetting.TokenSecret), 4096, 32, sha256.New)
	return base64.StdEncoding.EncodeToString(dk)
}

// GetRandomNum 随机生成多位随机数
func GetRandomNum(n int) string {
	var randomStr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	data := ""
	for i := 0; i < n; i++ {
		data += randomStr[rand.Intn(len(randomStr))]
	}
	return data
}
