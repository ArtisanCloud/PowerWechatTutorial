package miniprogram

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/services"
)

// 获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func APIURLSchemeGenerate(c *gin.Context) {

  path, exist := c.GetQuery("path")
  if !exist {
    panic("parameter path expected")
  }

  rs, err := services.MiniprogramApp.URLScheme.Generate(
    &power.HashMap{
      "path":  path,
      "query": "",
    },
    true,
    1,
    1606737600,
    30,
  )

  if err != nil {
    panic(err)
  }

  //content, _ := ioutil.ReadAll(rs.Body)
  ////fmt.Dump("content-type:",rs.Header.Get("Content-Type"))
  //c.Header("Content-Type", rs.Header.Get("Content-Type"))
  //c.Header("Content-Disposition", rs.Header.Get("attachment;filename=\""+rs.Header.Get("filename")+"\""))
  //c.Data(http.StatusOK, rs.Header.Get("Content-Type"), content)
  c.JSON(200, rs)
}
