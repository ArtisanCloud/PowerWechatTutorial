package wecom

import (
  "github.com/ArtisanCloud/power-wechat/src/work/invoice/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 查询电子发票
// https://open.work.weixin.qq.com/api/doc/90000/90135/90284
func APIInvoiceReimburseGetInvoiceInfo(c *gin.Context) {
  cardID := c.DefaultQuery("cardID", "CARDID")
  encryptCode := c.DefaultQuery("encryptCode", "ENCRYPTCODE")

  res, err := services.WeComApp.Invoice.GetInvoiceInfo(cardID, encryptCode)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 更新发票状态
// https://open.work.weixin.qq.com/api/doc/90000/90135/90285
func APIInvoiceReimburseUpdateInvoiceStatus(c *gin.Context) {
  cardID := c.DefaultQuery("cardID", "CARDID")
  encryptCode := c.DefaultQuery("encryptCode", "ENCRYPTCODE")
  reimburseStatus := c.DefaultQuery("status", "INVOICE_REIMBURSE_INIT")

  res, err := services.WeComApp.Invoice.UpdateInvoiceStatus(cardID, encryptCode, reimburseStatus)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 批量更新发票状态
// https://open.work.weixin.qq.com/api/doc/90000/90135/90286
func APIInvoiceReimburseUpdateStatusBatch(c *gin.Context) {
  openID := c.DefaultQuery("openID", "OPENID")
  reimburseStatus := c.DefaultQuery("status", "INVOICE_REIMBURSE_INIT")
  invoiceList := []*request.RequestCardInvoice{
    {CardID: "cardid_1", EncryptCode: "encrypt_code_1"},
    {CardID: "cardid_2", EncryptCode: "encrypt_code_2"},
  }

  res, err := services.WeComApp.Invoice.UpdateStatusBatch(openID, reimburseStatus, invoiceList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 批量查询电子发票
// https://open.work.weixin.qq.com/api/doc/90000/90135/90287
func APIInvoiceReimburseGetInvoiceInfoBatch(c *gin.Context) {

  invoiceList := []*request.RequestCardInvoice{
    {CardID: "CARDID1", EncryptCode: "ENCRYPTCODE1"},
    {CardID: "CARDID2", EncryptCode: "ENCRYPTCODE2"},
  }
  res, err := services.WeComApp.Invoice.GetInvoiceInfoBatch(invoiceList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
