package routes

import (
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/controllers/open-platform"
)

func InitOpenPlatformAPIRoutes(r *gin.Engine) {

	apiRouter := r.Group("/open-platform")
	{
		apiRouter.POST("/callback/verify/ticket", open_platform.APIOpenPlatformCallbackVerifyTicket)
		apiRouter.POST("/callback", open_platform.APIOpenPlatformCallback)

		apiRouter.GET("/getFastRegistrationURL", open_platform.GetFastRegistrationURL)
		apiRouter.GET("/auth", open_platform.HandleAuthorize)

		apiRouter.GET("/getAuthorizer", open_platform.GetAuthorizer)
		apiRouter.GET("/getAuthorizerList", open_platform.GetAuthorizers)
		apiRouter.GET("/getAuthorizerOption", open_platform.GetAuthorizerOption)
		apiRouter.GET("/setAuthorizerOption", open_platform.SetAuthorizerOption)

	}

}
