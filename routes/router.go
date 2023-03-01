package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

func InitializeRoutes(r *gin.Engine) {

	// Payment App Router
	InitPaymentAPIRoutes(r)

	// MiniProgram App Router
	InitMiniProgramAPIRoutes(r)

	// WeCom App Router
	InitWecomAPIRoutes(r)

	// Official App Router
	InitOfficialAPIRoutes(r)

	// OpenPlatform App Router
	InitOpenPlatformAPIRoutes(r)

	r.GET("/", func(c *gin.Context) {
		//c.String(200, "hello")
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte("Hello, PowerWechat"))

	})

	r.LoadHTMLGlob("templates/openplatform-auth.html")
	r.GET("/openplatform-auth", func(c *gin.Context) {
		c.HTML(http.StatusOK, "openplatform-auth.html", nil)
	})

}
