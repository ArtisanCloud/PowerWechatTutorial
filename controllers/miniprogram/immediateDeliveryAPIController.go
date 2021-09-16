package miniprogram

import (
  "github.com/ArtisanCloud/go-libs/object"
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 异常件退回商家商家确认收货接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.abnormalConfirm.html
func APIImmediateDeliveryAbnormalConfirm(c *gin.Context) {

  options := &power.HashMap{
    "shopid":        "123456",
    "shop_order_id": "123456",
    "shop_no":       "shop_no_111",
    "waybill_id":    "123456",
    "remark":        "remark",
    "delivery_sign": "123456",
  }

  rs, err := services.MiniProgramApp.Delivery.AbnormalConfirm(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 订单增加小费
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addOrder.html
func APIImmediateDeliveryAddOrder(c *gin.Context) {

  options := &power.HashMap{
    "cargo": &object.HashMap{
      "cargo_first_class":  "美食夜宵",
      "cargo_second_class": "零食小吃",
      "goods_detail": object.HashMap{
        "goods": []object.HashMap{
          object.HashMap{
            "good_count": 1,
            "good_name":  "水果",
            "good_price": 10,
            "good_unit":  "元",
          },
          object.HashMap{
            "good_count": 2,
            "good_name":  "蔬菜",
            "good_price": 20,
            "good_unit":  "元",
          },
        },
      },
      "goods_height": 1,
      "goods_length": 3,
      "goods_value":  5,
      "goods_weight": 1,
      "goods_width":  2,
    },
    "delivery_id":   "SFTC",
    "delivery_sign": "01234567890123456789",
    "openid":        "oABC123456",
    "order_info": object.HashMap{
      "delivery_service_code":  "",
      "expected_delivery_time": 0,
      "is_direct_delivery":     0,
      "is_finish_code_needed":  1,
      "is_insured":             0,
      "is_pickup_code_needed":  1,
      "note":                   "test_note",
      "order_time":             1555220757,
      "order_type":             0,
      "poi_seq":                "1111",
      "tips":                   0,
    },
    "receiver": object.HashMap{
      "address":         "xxx地铁站",
      "address_detail":  "2号楼202",
      "city":            "北京市",
      "coordinate_type": 0,
      "lat":             40.1529600000,
      "lng":             116.5060300000,
      "name":            "老王",
      "phone":           "18512345678",
    },
    "sender": object.HashMap{
      "address":         "xx大厦",
      "address_detail":  "1号楼101",
      "city":            "北京市",
      "coordinate_type": 0,
      "lat":             40.4486120000,
      "lng":             116.3830750000,
      "name":            "刘一",
      "phone":           "13712345678",
    },
    "shop": object.HashMap{
      "goods_count": 2,
      "goods_name":  "宝贝",
      "img_url":     "https://mmbiz.qpic.cn/mmbiz_png/xxxxxxxxx/0?wx_fmt=png",
      "wxa_path":    "/page/index/index",
    },
    "shop_no":        "12345678",
    "sub_biz_id":     "sub_biz_id_1",
    "shop_order_id":  "SFTC_001",
    "shopid":         "122222222",
    "delivery_token": "xxxxxxxx",
  }

  rs, err := services.MiniProgramApp.Delivery.AddOrder(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)

}

// 订单增加小费
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addTip.html
func APIImmediateDeliveryAddTip(c *gin.Context) {
  options := &power.HashMap{
    "shopid":        "123456",
    "shop_order_id": "123456",
    "waybill_id":    "123456",
    "tips":          5,
    "remark":        "gogogo",
    "delivery_sign": "123456",
    "shop_no":       "shop_no_111",
  }

  rs, err := services.MiniProgramApp.Delivery.AddTips(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 第三方代商户发起绑定配送公司帐号的请求
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.bindAccount.html
func APIImmediateDeliveryBindAccount(c *gin.Context) {

  deliveryID, exist := c.GetQuery("deliveryID")
  if !exist {
    panic("parameter delivery id expected")
  }

  rs, err := services.MiniProgramApp.Delivery.BindAccount(deliveryID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 取消配送单接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.cancelOrder.html
func APIImmediateDeliveryCancelOrder(c *gin.Context) {
  options := &power.HashMap{
    "shopid":           "123456",
    "shop_order_id":    "123456",
    "waybill_id":       "123456",
    "delivery_id":      "123456",
    "cancel_reason_id": 1,
    "cancel_reason":    "",
    "delivery_sign":    "123456",
    "shop_no":          "shop_no_111",
  }

  rs, err := services.MiniProgramApp.Delivery.CancelOrder(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取已支持的配送公司列表接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getAllImmeDelivery.html
func APIImmediateDeliveryGetAllImmeDelivery(c *gin.Context) {

  rs, err := services.MiniProgramApp.Delivery.GetAllImmeDelivery()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 拉取已绑定账号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getBindAccount.html
func APIImmediateDeliveryGetBindAccount(c *gin.Context) {
  rs, err := services.MiniProgramApp.Delivery.GetBindAccount()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 拉取配送单信息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getOrder.html
func APIImmediateDeliveryGetOrder(c *gin.Context) {
  rs, err := services.MiniProgramApp.Delivery.GetOrder("", "", "", "")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 模拟配送公司更新配送单状态
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.mockUpdateOrder.html
func APIImmediateDeliveryMockUpdateOrder(c *gin.Context) {

  options := &power.HashMap{
    "shopid":        "test_shop_id",
    "shop_order_id": "xxxxxxxxxxx",
    "waybill_id":    "xxxxxxxxxxxxx",
    "action_time":   12345678,
    "order_status":  101,
    "action_msg":    "",
  }

  rs, err := services.MiniProgramApp.Delivery.MockUpdateOrder(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 第三方代商户发起开通即时配送权限
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.openDelivery.html
func APIImmediateDeliveryOpenDelivery(c *gin.Context) {
  rs, err := services.MiniProgramApp.Delivery.OpenDelivery()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 预下配送单接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.preAddOrder.html
func APIImmediateDeliveryPreAddOrder(c *gin.Context) {
  options := &power.HashMap{
    "cargo": object.HashMap{
      "cargo_first_class":  "美食夜宵",
      "cargo_second_class": "零食小吃",
      "goods_detail": object.HashMap{
        "goods": []object.HashMap{
          {
            "good_count": 1,
            "good_name":  "水果",
            "good_price": 10,
            "good_unit":  "元",
          },
          {
            "good_count": 2,
            "good_name":  "蔬菜",
            "good_price": 20,
            "good_unit":  "元",
          },
        },
        "goods_height": 1,
        "goods_length": 3,
        "goods_value":  5,
        "goods_weight": 1,
        "goods_width":  2,
      },
      "delivery_id":   "SFTC",
      "delivery_sign": "01234567890123456789",
      "openid":        "oABC123456",
      "order_info": object.HashMap{
        "delivery_service_code":  "",
        "expected_delivery_time": 0,
        "is_direct_delivery":     0,
        "is_finish_code_needed":  1,
        "is_insured":             0,
        "is_pickup_code_needed":  1,
        "note":                   "test_note",
        "order_time":             1555220757,
        "order_type":             0,
        "poi_seq":                "1111",
        "tips":                   0,
      },
      "receiver": object.HashMap{
        "address":         "xxx地铁站",
        "address_detail":  "2号楼202",
        "city":            "北京市",
        "coordinate_type": 0,
        "lat":             40.1529600000,
        "lng":             116.5060300000,
        "name":            "老王",
        "phone":           "18512345678",
      },
      "sender": object.HashMap{
        "address":         "xx大厦",
        "address_detail":  "1号楼101",
        "city":            "北京市",
        "coordinate_type": 0,
        "lat":             40.4486120000,
        "lng":             116.3830750000,
        "name":            "刘一",
        "phone":           "13712345678",
      },
      "shop": object.HashMap{
        "goods_count": 2,
        "goods_name":  "宝贝",
        "img_url":     "https://mmbiz.qpic.cn/mmbiz_png/xxxxxxxxx/0?wx_fmt=png",
        "wxa_path":    "/page/index/index",
      },
      "shop_no":       "12345678",
      "sub_biz_id":    "sub_biz_id_1",
      "shop_order_id": "SFTC_001",
      "shopid":        "122222222",
    },
  }

  rs, err := services.MiniProgramApp.Delivery.PreAddOrder(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 预取消配送单接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.preCancelOrder.html
func APIImmediateDeliveryPreCancelOrder(c *gin.Context) {
  options := &power.HashMap{
    "shopid":           "123456",
    "shop_order_id":    "123456",
    "waybill_id":       "123456",
    "delivery_id":      "123456",
    "cancel_reason_id": 1,
    "cancel_reason":    "",
    "delivery_sign":    "123456",
    "shop_no":          "shop_no_111",
  }

  rs, err := services.MiniProgramApp.Delivery.PreCancelOrder(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 模拟配送公司更新配送单状态
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.realMockUpdateOrder.html
func APIImmediateDeliveryRealMockUpdateOrder(c *gin.Context) {
  options := &power.HashMap{
    "shopid":        "xxxxxxx",
    "shop_order_id": "xxxxxxxxxxx",
    "action_time":   1584145981,
    "order_status":  101,
    "action_msg":    "",
    "delivery_sign": "xxxxxxx",
  }

  rs, err := services.MiniProgramApp.Delivery.RealMockUpdateOrder(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 重新下单
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.reOrder.html
func APIImmediateDeliveryReOrder(c *gin.Context) {
  options := &power.HashMap{
    "cargo": object.HashMap{
      "cargo_first_class":  "美食夜宵",
      "cargo_second_class": "零食小吃",
      "goods_detail": object.HashMap{
        "goods": []object.HashMap{
          object.HashMap{
            "good_count": 1,
            "good_name":  "水果",
            "good_price": 10,
            "good_unit":  "元",
          },
          object.HashMap{
            "good_count": 2,
            "good_name":  "蔬菜",
            "good_price": 20,
            "good_unit":  "元",
          },
        },
        "goods_height": 1,
        "goods_length": 3,
        "goods_value":  5,
        "goods_weight": 1,
        "goods_width":  2,
      },
      "delivery_id":   "SFTC",
      "delivery_sign": "01234567890123456789",
      "openid":        "oABC123456",
      "order_info": object.HashMap{
        "delivery_service_code":  "",
        "expected_delivery_time": 0,
        "is_direct_delivery":     0,
        "is_finish_code_needed":  1,
        "is_insured":             0,
        "is_pickup_code_needed":  1,
        "note":                   "test_note",
        "order_time":             1555220757,
        "order_type":             0,
        "poi_seq":                "1111",
        "tips":                   0,
      },
      "receiver": object.HashMap{
        "address":         "xxx地铁站",
        "address_detail":  "2号楼202",
        "city":            "北京市",
        "coordinate_type": 0,
        "lat":             40.1529600000,
        "lng":             116.5060300000,
        "name":            "老王",
        "phone":           "18512345678",
      },
      "sender": object.HashMap{
        "address":         "xx大厦",
        "address_detail":  "1号楼101",
        "city":            "北京市",
        "coordinate_type": 0,
        "lat":             40.4486120000,
        "lng":             116.3830750000,
        "name":            "刘一",
        "phone":           "13712345678",
      },
      "shop": object.HashMap{
        "goods_count": 2,
        "goods_name":  "宝贝",
        "img_url":     "https://mmbiz.qpic.cn/mmbiz_png/xxxxxxxxx/0?wx_fmt=png",
        "wxa_path":    "/page/index/index",
      },
      "shop_no":        "12345678",
      "sub_biz_id":     "sub_biz_id_1",
      "shop_order_id":  "SFTC_001",
      "shopid":         "122222222",
      "delivery_token": "xxxxxxxx",
    },
  }

  rs, err := services.MiniProgramApp.Delivery.ReOrder(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
