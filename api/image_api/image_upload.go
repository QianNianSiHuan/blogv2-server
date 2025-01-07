package image_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/utils/fileTool"
	"blogv2/utils/hashTool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

func (ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("fileTool")
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	//文件大小判断
	s := global.Config.Upload.Size
	if fileHeader.Size > s*1024*1024 {
		res.FailWithMsg(c, fmt.Sprintf("文件大小大于 %d MB", s))
		return
	}
	//文件格式判断
	filename := fileHeader.Filename
	suffix, err := fileTool.ImageSuffixJudge(filename)
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	byteData, _ := io.ReadAll(file)
	hashString := hashTool.Md5(byteData)

	var model models.ImageModel
	err = global.DB.Take(&model, "hashTool = ?", hashString).Error
	if err == nil {
		//找到了
		logrus.Infof("上传文件重复 %s <==> %s %s", filename, model.Filename, hashString)
		res.Success(c, "上传成功", model.WebPath())
		return
	}

	filePath := fmt.Sprintf("uploads/%s/%s.%s", global.Config.Upload.UploadDir, hashString, suffix)
	model = models.ImageModel{
		Filename: filename,
		Path:     filePath,
		Size:     fileHeader.Size,
		Hash:     hashString,
	}

	err = global.DB.Create(&model).Error
	if err != nil {
		res.FailWithError(c, err)
		return
	}

	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	res.Success(c, model.WebPath(), "图片上传成功")
}
