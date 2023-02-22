package routes

import (
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/controllers/open-platform"
)

func InitOpenPlatformAPIRoutes(r *gin.Engine) {

	apiRouter := r.Group("/wx/api/openplatform")
	{
		// auth callback
		apiRouter.GET("/callback", open_platform.HandleAuthorize)
		apiRouter.POST("/callback", open_platform.APIOpenPlatformCallback)
		apiRouter.POST("/callback/:appID", open_platform.APIOpenPlatformCallbackWithApp)

		// auth
		apiRouter.GET("/createPreAuthCode", open_platform.APIOpenPlatformPreAuthCode)
		apiRouter.GET("/getFastRegistrationURL", open_platform.GetFastRegistrationURL)
		apiRouter.GET("/auth", open_platform.HandleAuthorize)

		// authorizer info
		apiRouter.GET("/getAuthorizer", open_platform.GetAuthorizer)
		apiRouter.GET("/getAuthorizerList", open_platform.GetAuthorizers)
		apiRouter.GET("/getAuthorizerOption", open_platform.GetAuthorizerOption)
		apiRouter.GET("/setAuthorizerOption", open_platform.SetAuthorizerOption)

		// delegate
		apiRouter.GET("/getAuthorizerOfficialAccount", open_platform.GetAuthorizerOfficialAccount)
		apiRouter.GET("/getAuthorizerOfficialAccountUser", open_platform.GetAuthorizerOfficialAccountUser)
		apiRouter.GET("/getAuthorizerMiniProgram", open_platform.GetAuthorizerMiniProgram)
		apiRouter.GET("/authorizerAccountCreate", open_platform.AuthorizerAccountCreate)
		apiRouter.GET("/authorizerAccountBind", open_platform.AuthorizerAccountBind)
		apiRouter.GET("/authorizerAccountUnbind", open_platform.AuthorizerAccountUnbind)
		apiRouter.GET("/authorizerAccountGet", open_platform.AuthorizerAccountGet)
		apiRouter.GET("/officialAccountDemo", open_platform.OfficialAccountDemo)
		apiRouter.GET("/miniProgramDemo", open_platform.MiniProgramDemo)

	}

}
