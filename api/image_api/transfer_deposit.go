package image_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/utils/hashTool"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

// 图片转存
type TransferDepositRequest struct {
	Url string `json:"url" binding:"required"`
}

func (ImageApi) TransferDepositView(c *gin.Context) {
	var cr TransferDepositRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg(c, "图片转存请求失败")
		return
	}
	response, err := http.Get(cr.Url)
	byteData, _ := io.ReadAll(response.Body)
	hashString := hashTool.Md5(byteData)
	suffixData := strings.Split(response.Header.Get("Content-Type"), "/")
	suffix := suffixData[len(suffixData)-1]
	filePath := fmt.Sprintf("uploads/%s/%s.%s", global.Config.Upload.UploadDir, hashString, suffix)
	err = os.WriteFile(filePath, byteData, 0666)
	if err != nil {
		res.FailWithMsg(c, "图片保存失败")
		return
	}
	res.SuccessWithData(c, "/"+filePath)
}
