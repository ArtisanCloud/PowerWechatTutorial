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

	r.GET("/", func(c *gin.Context) {
		//c.String(200, "hello")
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte("Hello, PowerWechat"))
	})

}
