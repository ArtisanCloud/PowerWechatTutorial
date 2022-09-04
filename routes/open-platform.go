package routes

import (
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/controllers/open-platform"
)

func InitOpenPlatformAPIRoutes(r *gin.Engine) {

	apiRouter := r.Group("/open-platform")
	{
		apiRouter.POST("/callback", open_platform.APIOpenPlatformCallback)
	}

}
