package message

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func APIAppChatCreate(c *gin.Context) {

  res, err := services.WeComApp.MessageAppChat.Create(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)

}

func APIAppChatUpdate(c *gin.Context) {

}

func APIAppChatGet(c *gin.Context) {

}



func APIAppChatSendText(c *gin.Context) {
  
}

func APIAppChatSendImage(c *gin.Context) {
  
}

func APIAppChatSendVoice(c *gin.Context) {
  
}

func APIAppChatSendVideo(c *gin.Context) {
  
}

func APIAppChatSendFile(c *gin.Context) {
  
}

func APIAppChatSendTextcard(c *gin.Context) {
  
}

func APIAppChatSendNews(c *gin.Context) {
  
}

func APIAppChatSendMpnews(c *gin.Context) {
  
}

func APIAppChatSendMarkdown(c *gin.Context) {

  
}


