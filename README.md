# PowerWeChat 接口example路由表

-- ./routes/

* apiMiniProgram.go
* apiOfficial.go
* apiPayment.go
* apiWecom.go

寻找企业微信的成员功能接口

```go
// ./routes/apiWecom.go

wecomRouter := r.Group("/wecom")
{
  // Handle user route
  wecomRouter.POST("/user/create", validate.ValidateUserCreate, user.APIUserCreate)
  wecomRouter.GET("/user/get", user.APIUserGet)
  wecomRouter.PUT("/user/update", user.APIUserUpdate)
  wecomRouter.DELETE("/user/delete", user.APIUserDelete)
  wecomRouter.DELETE("/user/batch", user.APIUserBatchDelete)

  ...

}
```

example: 获取成员信息，使用Insomian或postman等接口调试工具，提交json数据,获取成员。

```go

// ./routes/controllers/wecom/user/userAPIController.go
// UserGet 读取成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90196
func APIUserGet(c *gin.Context) {
  userId := c.DefaultQuery("userId", "TestUserId")
  res, err := services.WeComContactApp.User.Get(userId)
  if err != nil {
	  panic(err)
  }

  c.JSON(http.StatusOK, res)
}
...

```

