package official_account

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/goods/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// GoodsAdd 导入商品
func GoodsAdd(ctx *gin.Context) {
	products := &request.RequestProduct{
		Product: []*request.Product{
			{
				PID: "pid001",
				ImageInfo: &request.ImageInfo{
					MainImageList: []*request.MainImage{
						{URL: "http://www.google.com/a.jpg"},
					},
				},
				CategoryInfo: &request.CategoryInfo{
					CategoryItem: []*request.CategoryItem{
						{CategoryName: "图书"},
						{CategoryName: "少儿图书"},
					},
				},
				OfficialCategoryInfo: &request.OfficialCategoryInfo{
					CategoryItem: []*request.CategoryItem{
						{CategoryName: "图书"},
					},
				},
				LinkInfo: &request.LinkInfo{
					URL:      "pages/index/index",
					WxaAppID: "wxa0x01adx3423566",
					LinkType: "wxa",
				},
				Title:    "test title",
				SubTitle: "test sub_title",
				Brand:    "test brand",
				ShopInfo: &request.ShopInfo{
					Source: 2,
				},
				Desc: "test desc",
				PriceInfo: &request.PriceInfo{
					MinPrice:    250,
					MaxPrice:    250.22,
					MinOriPrice: 300.1,
					MaxOriPrice: 320.15,
				},
				SaleInfo: &request.SaleInfo{
					SaleStatus: "on",
					Stock:      1000,
				},
				CustomInfo: &request.CustomInfo{CustomList: []*request.Custom{
					{
						Key:   "book_desc",
						Value: "“熊猫先生”通过4个富有节律性、带有因果关系、幽默风趣，又有正反对比的故事，让小朋友明白礼仪的重要性。",
					},
					{
						Key:   "author",
						Value: "史蒂夫•安东尼",
					},
					{
						Key:   "publisher",
						Value: "中信出版社",
					},
				}},
				SKUInfo: &request.SKUInfo{
					SKUItem: []*request.SKUItem{
						{
							SKUID:       "sku001",
							BarcodeType: "ean13",
							Barcode:     2018032105140,
							ImageInfo: &request.ImageInfo{
								MainImageList: []*request.MainImage{
									{URL: "http://www.baidu.com/c.jpg"},
									{URL: "http://www.baidu.com/d.jpg"},
								},
							},
							LinkUrl: "pages/index/index?a=b",
							PriceInfo: &request.PriceInfo{
								MinPrice:    250,
								MaxPrice:    250.22,
								MinOriPrice: 300.1,
								MaxOriPrice: 320.15,
							},
							SaleInfo: &request.SaleInfo{
								SaleStatus: "on",
								Stock:      500,
							},
							ShopInfo: &request.ShopInfo{
								Source: 2,
							},
							AttributeInfo: &request.AttributeInfo{
								AttributeItem: []*request.AttributeItem{
									{
										AttributeKey:   "酒店房型",
										AttributeValue: "海景大床房",
									},
								},
							},
						},
					},
				},
				PartialUpdate: 0,
				TagInfo:       nil,
			},
		},
	}
	data, err := services.OfficialAccountApp.Goods.Add(products)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GoodsUpdate 更新商品
func GoodsUpdate(ctx *gin.Context) {
	products := &request.RequestProduct{
		Product: []*request.Product{
			{
				PID: "pid001",
				ImageInfo: &request.ImageInfo{
					MainImageList: []*request.MainImage{
						{URL: "http://www.google.com/a.jpg"},
					},
				},
				CategoryInfo: &request.CategoryInfo{
					CategoryItem: []*request.CategoryItem{
						{CategoryName: "图书"},
						{CategoryName: "少儿图书"},
					},
				},
				OfficialCategoryInfo: &request.OfficialCategoryInfo{
					CategoryItem: []*request.CategoryItem{
						{CategoryName: "图书"},
					},
				},
				LinkInfo: &request.LinkInfo{
					URL:      "pages/index/index",
					WxaAppID: "wxa0x01adx3423566",
					LinkType: "wxa",
				},
				Title:    "test title",
				SubTitle: "test sub_title",
				Brand:    "test brand",
				ShopInfo: &request.ShopInfo{
					Source: 2,
				},
				Desc: "test desc",
				PriceInfo: &request.PriceInfo{
					MinPrice:    250,
					MaxPrice:    250.22,
					MinOriPrice: 300.1,
					MaxOriPrice: 320.15,
				},
				SaleInfo: &request.SaleInfo{
					SaleStatus: "on",
					Stock:      1000,
				},
				CustomInfo: &request.CustomInfo{CustomList: []request.Custom{
					{
						Key:   "book_desc",
						Value: "“熊猫先生”通过4个富有节律性、带有因果关系、幽默风趣，又有正反对比的故事，让小朋友明白礼仪的重要性。",
					},
					{
						Key:   "author",
						Value: "史蒂夫•安东尼",
					},
					{
						Key:   "publisher",
						Value: "中信出版社",
					},
				}},
				SKUInfo: &request.SKUInfo{
					SKUItem: []*request.SKUItem{
						{
							SKUID:       "sku001",
							BarcodeType: "ean13",
							Barcode:     2018032105140,
							ImageInfo: &request.ImageInfo{
								MainImageList: []*request.MainImage{
									{URL: "http://www.baidu.com/c.jpg"},
									{URL: "http://www.baidu.com/d.jpg"},
								},
							},
							LinkUrl: "pages/index/index?a=b",
							PriceInfo: &request.PriceInfo{
								MinPrice:    250,
								MaxPrice:    250.22,
								MinOriPrice: 300.1,
								MaxOriPrice: 320.15,
							},
							SaleInfo: &request.SaleInfo{
								SaleStatus: "on",
								Stock:      500,
							},
							ShopInfo: &request.ShopInfo{
								Source: 2,
							},
							AttributeInfo: &request.AttributeInfo{
								AttributeItem: []*request.AttributeItem{
									{
										AttributeKey:   "酒店房型",
										AttributeValue: "海景大床房",
									},
								},
							},
						},
					},
				},
				PartialUpdate: 0,
				TagInfo:       nil,
			},
		},
	}
	data, err := services.OfficialAccountApp.Goods.Update(products)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GoodsStatus 查询导入/更新商品状态
func GoodsStatus(c *gin.Context) {
	statusTicket := c.DefaultQuery("statusTicket", "115141102647330200")
	res, err := services.OfficialAccountApp.Goods.Status(statusTicket)

	if err != nil {
		panic(err)
	}
	c.JSON(200, res)
}

// GoodsGet 获取单个商品信息
func GoodsGet(c *gin.Context) {
	pid := c.DefaultQuery("pid", "pid001")
	res, err := services.OfficialAccountApp.Goods.Get(&request.RequestProductGet{
		Product: &request.ProductID{
			PID: pid,
		},
	})

	if err != nil {
		panic(err)
	}
	c.JSON(200, res)
}

// GoodsList 分页获取商品信息
func GoodsList(c *gin.Context) {

	context := c.DefaultQuery("context", "")
	page := 1
	size := 10
	res, err := services.OfficialAccountApp.Goods.List(context, page, size)

	if err != nil {
		panic(err)
	}
	c.JSON(200, res)
}
