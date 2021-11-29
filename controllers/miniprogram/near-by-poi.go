package miniprogram

import (
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
  "strconv"
)

// 添加地点
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.add.html
func APINearbyPoiAdd(c *gin.Context) {

  rs, err := services.MiniProgramApp.NearbyPoi.Add(&power.HashMap{
    "is_comm_nearby":     "1", //值固定
    "kf_info":            "{\"open_kf\":true,\"kf_headimg\":\"http://mmbiz.qpic.cn/mmbiz_jpg/kKMgNtnEfQzDKpLXYhgo3W3Gndl34gITqmP914zSwhajIEJzUPpx40P7R8fRe1QmicneQMhFzpZNhSLjrvU1pIA/0?wx_fmt=jpeg\",\"kf_name\":\"Harden\"}",
    "pic_list":           "{\"list\":[\"http://mmbiz.qpic.cn/mmbiz_jpg/kKMgNtnEfQzDKpLXYhgo3W3Gndl34gITqmP914zSwhajIEJzUPpx40P7R8fRe1QmicneQMhFzpZNhSLjrvU1pIA/0?wx_fmt=jpeg\",\"http://mmbiz.qpic.cn/mmbiz_jpg/kKMgNtnEfQzDKpLXYhgo3W3Gndl34gITRneE5FS9uYruXGMmrtmhsBySwddEWUGOibG8Ze2NT5E3Dyt79I0htNg/0?wx_fmt=jpeg\"]}",
    "service_infos":      "{\"service_infos\":[{\"id\":2,\"type\":1,\"name\":\"快递\",\"appid\":\"wx1373169e494e0c39\",\"path\":\"index\"},{\"id\":0,\"type\":2,\"name\":\"自定义\",\"appid\":\"wx1373169e494e0c39\",\"path\":\"index\"}]}",
    "store_name":         "羊村小马烧烤",
    "contract_phone":     "111111111",
    "hour":               "00:00-11:11",
    "company_name":       "深圳市腾讯计算机系统有限公司",
    "credential":         "156718193518281",
    "address":            "新疆维吾尔自治区克拉玛依市克拉玛依区碧水路15-1-8号(碧水云天广场)",
    "qualification_list": "3LaLzqiTrQcD20DlX_o-OV1-nlYMu7sdVAL7SV2PrxVyjZFZZmB3O6LPGaYXlZWq",
    "poi_id":             "",
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

func APINearbyPoiDelete(c *gin.Context) {

  poiID, exist := c.GetQuery("poiID")
  if !exist {
    panic("parameter poi id expected")
  }

  rs, err := services.MiniProgramApp.NearbyPoi.Delete(poiID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

func APINearbyPoiGetList(c *gin.Context) {

  strPage := c.DefaultQuery("page", "1")

  page, err := strconv.Atoi(strPage)

  strPageRows := c.DefaultQuery("pageRows", "100")

  pageRows, err := strconv.Atoi(strPageRows)

  rs, err := services.MiniProgramApp.NearbyPoi.GetList(page, pageRows)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

func APINearbySetShowStatus(c *gin.Context) {

  poiID, exist := c.GetQuery("poiID")
  if !exist {
    panic("parameter poi id expected")
  }

  rs, err := services.MiniProgramApp.NearbyPoi.SetShowStatus(poiID, 1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
