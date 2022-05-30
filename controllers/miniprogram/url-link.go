package miniprogram

import (
  "github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/urlLink/request"
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "net/http"
  "power-wechat-tutorial/services"
)

// 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
func APIURLLinkGenerate(c *gin.Context) {

  path := c.DefaultQuery("path", "/pages/publishHomework/publishHomework")

  rs, err := services.MiniProgramApp.URLLink.Generate(&request.URLSchemeGenerate{
    JumpWxa: &request.JumpWxa{
      Path:  path,
      Query: "",
    },
    IsExpire:   true,
    ExpireTime: 1606737600,
  })

  if err != nil {
    panic(err)
  }

  content, _ := ioutil.ReadAll(rs.Body)
  //fmt.Dump("content-type:",rs.Header.Get("Content-Type"))
  c.Header("Content-Type", rs.Header.Get("Content-Type"))
  c.Header("Content-Disposition", rs.Header.Get("attachment;filename=\""+rs.Header.Get("filename")+"\""))
  c.Data(http.StatusOK, rs.Header.Get("Content-Type"), content)

}
