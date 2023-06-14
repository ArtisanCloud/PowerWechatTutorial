package merchant

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchant/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path"
	"power-wechat-tutorial/services"
)

func APIUploadImg(c *gin.Context) {

	// 读取图片文件
	homePath, _ := os.UserHomeDir()
	imagePath := path.Join(homePath, "Desktop/641.png")
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println("无法读取图片文件:", err)
		os.Exit(1)
	}

	hash := sha256.Sum256(imageData)
	hashString := hex.EncodeToString(hash[:])

	para := &request.RequestMediaUpload{
		File: imagePath,
		Meta: &request.Meta{
			Filename: "641.png",
			Sha256:   hashString,
		},
	}

	rs, err := services.PaymentApp.Merchant.UploadImg(c.Request.Context(), para)
	if err != nil {
		log.Println("出错了： ", err)
		c.String(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, rs)

}
