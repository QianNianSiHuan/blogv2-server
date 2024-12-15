package captcha_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
	"image/color"
)

type CaptchaApi struct {
}
type CaptchaResponse struct {
	CaptchaID string `json:"captchaID"`
	Captcha   string `json:"captcha"`
}

func (CaptchaApi) CaptchaView(c *gin.Context) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		// 验证码图片的高度，以像素为单位。
		Height: 60,
		//验证码图片的宽度，以像素为单位。
		Width: 200,
		//验证码图片中随机噪点的数量。在这个例子中，值为0表示没有噪点。
		NoiseCount: 5,
		// 控制显示在验证码图片中的线条的选项。在这个例子中，1: 直线  2: 曲线4: 点线8: 虚线16: 中空直线32: 中空曲线
		ShowLineOptions: 2 | 4,
		//验证码的长度，即验证码中字符的数量。
		Length: 4,
		//验证码的字符源，用于生成验证码的字符。在这个例子中，使用数字和小写字母作为字符源。
		Source: "1234567890",
		BgColor: &color.RGBA{
			//验证码图片的背景颜色。在这个例子中，使用RGBA颜色模型，R表示红色分量，G表示绿色分量，B表示蓝色分量，A表示透明度。
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		//用于绘制验证码文本的字体文件。在这个例子中，使用名为"wqy-microhei.ttc"的字体文件。
		//Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	//将driverString中指定的字体文件转换为驱动程序所需的字体格式，并将结果赋值给driver变量。这个步骤是为了将字体文件转换为正确的格式，以便在生成验证码时使用正确的字体。
	driver = driverString.ConvertFonts()
	//使用driver和stores参数创建一个新的验证码实例，并将其赋值给captcha变量。这里的stores参数表示验证码存储器，用于存储和验证验证码。
	captcha := base64Captcha.NewCaptcha(driver, global.CaptchaStore)
	//调用captcha实例的Generate方法生成验证码。lid是生成的验证码的唯一标识符，lb64s是生成的验证码图片的Base64编码字符串，lerr是生成过程中的任何错误。
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		logrus.Error(err.Error())
		res.FailWithError(c, err)
		return
	}
	res.SuccessWithData(c, CaptchaResponse{
		CaptchaID: id,
		Captcha:   b64s,
	})
}
