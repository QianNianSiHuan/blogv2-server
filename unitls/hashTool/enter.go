package hashTool

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func Md5(data []byte) string {
	md5New := md5.New()
	md5New.Write(data)
	return hex.EncodeToString(md5New.Sum(nil))
}
func FileMd5(file string) (hash string, err error) {
	byteData, err := os.ReadFile(file)
	if err != nil {
		return
	}
	hash = Md5(byteData)
	return
}
