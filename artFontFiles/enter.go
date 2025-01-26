package artFontFiles

import (
	"embed"
	"fmt"
	"github.com/sirupsen/logrus"
)

//go:embed artFont
var artFontFiles embed.FS

type ArtisticFontPath string

const (
	FAIL       ArtisticFontPath = "artFont/fail.txt"
	GORM_DEBUG ArtisticFontPath = "artFont/gorm-debug.txt"
	SUCCESS    ArtisticFontPath = "artFont/success.txt"
	WELCOME    ArtisticFontPath = "artFont/welcome!.txt"
	GIN_DEBUG  ArtisticFontPath = "artFontFiles/gin-debug"
)

// OutPutArtisticFont 输出txt文档内容到控制台
func OutPutArtisticFont(filePath ArtisticFontPath) {
	file, err := artFontFiles.ReadFile(string(filePath))
	if err != nil {
		logrus.Errorf("无法打开文件: %s", err)
		return
	}
	fmt.Println(string(file))
	//defer file.Close()
	//// 创建一个新的扫描器来读取文件
	//scanner := bufio.NewScanner(string(file))
	//// 设置扫描器的分割函数为扫描整行（这是默认行为）
	//scanner.Split(bufio.ScanLines)
	//// 逐行读取文件内容并打印
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	fmt.Println(line)
	//}
	////检查是否有错误发生在扫描过程中
	//if err := scanner.Err(); err != nil {
	//	logrus.Errorf("读取文件时发生错误: %s", err)
	//}
}
