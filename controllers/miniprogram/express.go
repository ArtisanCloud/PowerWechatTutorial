package miniprogram

import (
  "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 生成运单
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.addOrder.html
func APIExpressAddOrder(c *gin.Context) {

  options := &power.HashMap{

    "add_source":    0,
    "order_id":      "01234567890123456789",
    "openid":        "oABC123456",
    "delivery_id":   "SF",
    "biz_id":        "xyz",
    "custom_remark": "易碎物品",
    "sender": power.HashMap{
      "name":      "张三",
      "tel":       "020-88888888",
      "mobile":    "18666666666",
      "company":   "公司名",
      "post_code": "123456",
      "country":   "中国",
      "province":  "广东省",
      "city":      "广州市",
      "area":      "海珠区",
      "address":   "XX路XX号XX大厦XX栋XX",
    },
    "receiver": power.HashMap{
      "name":      "王小蒙",
      "tel":       "020-77777777",
      "mobile":    "18610000000",
      "company":   "公司名",
      "post_code": "654321",
      "country":   "中国",
      "province":  "广东省",
      "city":      "广州市",
      "area":      "天河区",
      "address":   "XX路XX号XX大厦XX栋XX",
    },
    "shop": power.HashMap{
      "wxa_path":    "/index/index?from=waybill&id=01234567890123456789",
      "img_url":     "https://mmbiz.qpic.cn/mmbiz_png/OiaFLUqewuIDNQnTiaCInIG8ibdosYHhQHPbXJUrqYSNIcBL60vo4LIjlcoNG1QPkeH5GWWEB41Ny895CokeAah8A/640",
      "goods_name":  "微信气泡狗抱枕&微信气泡狗钥匙扣",
      "goods_count": 2,
    },
    "cargo": power.HashMap{
      "count":   2,
      "weight":  5.5,
      "space_x": 30.5,
      "space_y": 20,
      "space_z": 20,
      "detail_list": []power.HashMap{
        power.HashMap{
          "name":  "微信气泡狗抱枕",
          "count": 1,
        },
        power.HashMap{
          "name":  "微信气泡狗钥匙扣",
          "count": 1,
        },
      },
      "insured": power.HashMap{
        "use_insured":   1,
        "insured_value": 10000,
      },
      "service": power.HashMap{
        "service_type": 0,
        "service_name": "标准快递",
      },
    },
  }

  rs, err := services.MiniProgramApp.Express.AddOrder(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 批量获取运单数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.batchGetOrder.html
func APIExpressBatchGetOrder(c *gin.Context) {
  options := []*power.HashMap{
    &power.HashMap{
      "order_id":    "01234567890123456789",
      "delivery_id": "SF",
      "waybill_id":  "123456789",
    },
    &power.HashMap{
      "order_id":    "01234567890123456789",
      "delivery_id": "SF",
      "waybill_id":  "123456789",
    },
  }

  rs, err := services.MiniProgramApp.Express.BatchGetOrder(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 绑定、解绑物流账号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.bindAccount.html
func APIExpressBindAccount(c *gin.Context) {
  rs, err := services.MiniProgramApp.Express.BindAccount("bind", "123456", "YUNDA", "123456789123456789")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 取消运单
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.cancelOrder.html
func APIExpressCancelOrder(c *gin.Context) {

  orderID, exist := c.GetQuery("orderID")
  if !exist {
    panic("parameter order id expected")
  }
  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  deliveryID, exist := c.GetQuery("deliveryID")
  if !exist {
    panic("parameter delivery id expected")
  }

  wayBillID, exist := c.GetQuery("wayBillID")
  if !exist {
    panic("parameter way bill id expected")
  }

  rs, err := services.MiniProgramApp.Express.CancelOrder(orderID, openID, deliveryID, wayBillID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取所有绑定的物流账号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getAllAccount.html
func APIExpressGetAllAccount(c *gin.Context) {
  rs, err := services.MiniProgramApp.Express.GetAllAccount()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取支持的快递公司列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getAllDelivery.html
func APIExpressGetAllDelivery(c *gin.Context) {
  rs, err := services.MiniProgramApp.Express.GetAllDelivery()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取运单数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getOrder.html
func APIExpressGetOrder(c *gin.Context) {

  orderID, exist := c.GetQuery("orderID")
  if !exist {
    panic("parameter order id expected")
  }
  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  deliveryID, exist := c.GetQuery("deliveryID")
  if !exist {
    panic("parameter delivery id expected")
  }

  wayBillID, exist := c.GetQuery("wayBillID")
  if !exist {
    panic("parameter way bill id expected")
  }

  rs, err := services.MiniProgramApp.Express.GetOrder(orderID, openID, deliveryID, wayBillID, 1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 查询运单轨迹
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getPath.html
func APIExpressGetPath(c *gin.Context) {
  orderID, exist := c.GetQuery("orderID")
  if !exist {
    panic("parameter order id expected")
  }
  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  deliveryID, exist := c.GetQuery("deliveryID")
  if !exist {
    panic("parameter delivery id expected")
  }

  wayBillID, exist := c.GetQuery("wayBillID")
  if !exist {
    panic("parameter way bill id expected")
  }

  rs, err := services.MiniProgramApp.Express.GetPath(orderID, openID, deliveryID, wayBillID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取打印员
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getPrinter.html
func APIExpressGetPrinter(c *gin.Context) {
  rs, err := services.MiniProgramApp.Express.GetPrinter()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取电子面单余额
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getQuota.html
func APIExpressGetQuota(c *gin.Context) {
  rs, err := services.MiniProgramApp.Express.GetQuota("YTO", "xyz")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 模拟快递公司更新订单状态, 该接口只能用户测试
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.testUpdateOrder.html
func APIExpressTestUpdateOrder(c *gin.Context) {

  rs, err := services.MiniProgramApp.Express.TestUpdateOrder("test_biz_id", "xxxxxxxxxxxx", "TEST", "xxxxxxxxxx", 123456789, 100001, "揽件阶段")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 配置面单打印员
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.updatePrinter.html
func APIExpressUpdatePrinter(c *gin.Context) {
  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  rs, err := services.MiniProgramApp.Express.UpdatePrinter(openID, "bind", "123,456")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取面单联系人信息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.getContact.html
func APIExpressGetContact(c *gin.Context) {
  token, exist := c.GetQuery("token")
  if !exist {
    panic("parameter token expected")
  }

  waybillID, exist := c.GetQuery("waybill_id")
  if !exist {
    panic("parameter waybill id expected")
  }

  rs, err := services.MiniProgramApp.Express.GetContact(token, waybillID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 预览面单模板
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.previewTemplate.html
func APIExpressPreviewTemplate(c *gin.Context) {
  orderID, exist := c.GetQuery("orderID")
  if !exist {
    panic("parameter order id expected")
  }
  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  options := &power.HashMap{
    "waybill_id":       "1234567890123",
    "waybill_data":     "##ZTO_mark##11-22-33##ZTO_bagAddr##广州##",
    "waybill_template": "PGh0bWw+dGVzdDwvaHRtbD4=",
    "custom": power.HashMap{
      "order_id":      orderID,
      "openid":        openID,
      "delivery_id":   "ZTO",
      "biz_id":        "xyz",
      "custom_remark": "易碎物品",
      "sender": power.HashMap{
        "name":      "张三",
        "tel":       "18666666666",
        "mobile":    "020-88888888",
        "company":   "公司名",
        "post_code": "123456",
        "country":   "中国",
        "province":  "广东省",
        "city":      "广州市",
        "area":      "海珠区",
        "address":   "XX路XX号XX大厦XX栋XX",
      },
      "receiver": power.HashMap{
        "name":      "王小蒙",
        "tel":       "18610000000",
        "mobile":    "020-77777777",
        "company":   "公司名",
        "post_code": "654321",
        "country":   "中国",
        "province":  "广东省",
        "city":      "广州市",
        "area":      "天河区",
        "address":   "XX路XX号XX大厦XX栋XX",
      },
      "shop": power.HashMap{
        "wxa_path":    "/index/index?from=waybill",
        "img_url":     "https://mmbiz.qpic.cn/mmbiz_png/KfrZwACMrmwbPGicysN6kibW0ibXwzmA3mtTwgSsdw4Uicabduu2pfbfwdKicQ8n0v91kRAUX6SDESQypl5tlRwHUPA/640",
        "goods_name":  "一千零一夜钻石包&爱马仕柏金钻石包",
        "goods_count": 2,
      },
      "cargo": power.HashMap{
        "count":   2,
        "weight":  5.5,
        "space_x": 30.5,
        "space_y": 20,
        "space_z": 20,
        "detail_list": []power.HashMap{
          power.HashMap{
            "name":  "一千零一夜钻石包",
            "count": 1,
          },
          power.HashMap{
            "name":  "爱马仕柏金钻石包",
            "count": 1,
          },
        },
        "insured": power.HashMap{
          "use_insured":   1,
          "insured_value": 10000,
        },
        "service": power.HashMap{
          "service_type": 0,
          "service_name": "标准快递",
        },
      },
    },
  }

  rs, err := services.MiniProgramApp.Express.PreviewTemplate(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 更新商户审核结果
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.updateBusiness.html
func APIExpressUpdateBusiness(c *gin.Context) {

  options := &power.HashMap{
    "shop_app_id": "wxABCD",
    "biz_id":      "xyz",
    "result_code": 0,
    "result_msg":  "审核通过",
  }

  rs, err := services.MiniProgramApp.Express.UpdateBusiness(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 更新运单轨迹
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.updatePath.html
func APIExpressUpdatePath(c *gin.Context) {

  token, exist := c.GetQuery("token")
  if !exist {
    panic("parameter token expected")
  }

  waybillID, exist := c.GetQuery("waybill_id")
  if !exist {
    panic("parameter waybill id expected")
  }

  options := &power.HashMap{
    "token":       token,
    "waybill_id":  waybillID,
    "action_time": 1533052800,
    "action_type": 300002,
    "action_msg":  "丽影邓丽君【18666666666】正在派件",
  }
  rs, err := services.MiniProgramApp.Express.UpdatePath(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
