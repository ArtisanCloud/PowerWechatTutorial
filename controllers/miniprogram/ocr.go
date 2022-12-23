package miniprogram

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"power-wechat-tutorial/services"
)

// 本接口提供基于小程序的银行卡 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.bankcard.html
func APIOCRBankCardByURL(c *gin.Context) {
	url, exist := c.GetQuery("url")
	if !exist {
		panic("parameter url expected")
	}

	rs, err := services.MiniProgramApp.OCR.Bankcard(url, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的银行卡 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.bankcard.html
func APIOCRBankCardByData(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)
	rs, err := services.MiniProgramApp.OCR.Bankcard("", &power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的营业执照 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.businessLicense.html
func APIOCRBusinessLicenseByURL(c *gin.Context) {
	url, exist := c.GetQuery("url")
	if !exist {
		panic("parameter url expected")
	}

	rs, err := services.MiniProgramApp.OCR.BusinessLicense(url, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的营业执照 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.businessLicense.html
func APIOCRBusinessLicenseByData(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)
	rs, err := services.MiniProgramApp.OCR.BusinessLicense("", &power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的驾驶证 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.driverLicense.html
func APIOCRDriverLicenseByURL(c *gin.Context) {
	url, exist := c.GetQuery("url")
	if !exist {
		panic("parameter url expected")
	}

	rs, err := services.MiniProgramApp.OCR.DriverLicense(url, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的驾驶证 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.driverLicense.html
func APIOCRDriverLicenseByData(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)
	rs, err := services.MiniProgramApp.OCR.DriverLicense("", &power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的身份证 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.idcard.html
func APIOCRIDCardByURL(c *gin.Context) {
	url, exist := c.GetQuery("url")
	if !exist {
		panic("parameter url expected")
	}

	rs, err := services.MiniProgramApp.OCR.IDCard(url, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的身份证 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.idcard.html
func APIOCRIDCardByData(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)
	rs, err := services.MiniProgramApp.OCR.IDCard("", &power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的通用印刷体 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.printedText.html
func APIOCRPrintedTextByURL(c *gin.Context) {
	url, exist := c.GetQuery("url")
	if !exist {
		panic("parameter url expected")
	}

	rs, err := services.MiniProgramApp.OCR.PrintedText(url, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的通用印刷体 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.printedText.html
func APIOCRPrintedTextByData(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)
	rs, err := services.MiniProgramApp.OCR.PrintedText("", &power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的行驶证 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.vehicleLicense.html
func APIOCRVehicleLicenseByURL(c *gin.Context) {
	url, exist := c.GetQuery("url")
	if !exist {
		panic("parameter url expected")
	}

	rs, err := services.MiniProgramApp.OCR.VehicleLicense("photo", url, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的行驶证 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.vehicleLicense.html
func APIOCRVehicleLicenseByData(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)
	rs, err := services.MiniProgramApp.OCR.VehicleLicense("photo", "", &power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
