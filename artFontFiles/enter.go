package artFontFiles

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type ArtisticFontPath string

const (
	FAIL       ArtisticFontPath = "artFontFiles/fail.txt"
	GORM_DEBUG ArtisticFontPath = "artFontFiles/gorm-debug.txt"
	SUCCESS    ArtisticFontPath = "artFontFiles/success.txt"
	WELCOME    ArtisticFontPath = "artFontFiles/welcome!.txt"
	GIN_DEBUG  ArtisticFontPath = "artFontFiles/gin-debug"
)

// OutPutArtisticFont 输出txt文档内容到控制台
func OutPutArtisticFont(filePath ArtisticFontPath) {
	file, err := os.Open(string(filePath))
	if err != nil {
		logrus.Errorf("无法打开文件: %s", err)
		return
	}
	defer file.Close()
	// 创建一个新的扫描器来读取文件
	scanner := bufio.NewScanner(file)
	// 设置扫描器的分割函数为扫描整行（这是默认行为）
	scanner.Split(bufio.ScanLines)
	// 逐行读取文件内容并打印
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	// 检查是否有错误发生在扫描过程中
	if err := scanner.Err(); err != nil {
		logrus.Errorf("读取文件时发生错误: %s", err)
	}
}
