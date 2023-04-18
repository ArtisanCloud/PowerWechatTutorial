package routes

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var Router *gin.Engine

func InitializeRoutes(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			log.Println("origin: ", origin)
			return true
		},
	}))

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

	// RoBot Chat App Router
	InitRobotChatAPIRoutes(r)

	r.GET("/", func(c *gin.Context) {
		//c.String(200, "hello")
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte("Hello, PowerWechat"))

	})

	r.GET("/json/user", func(context *gin.Context) {
		obj := &power.HashMap{
			"say":       "I am",
			"something": "ArtisanCloud",
		}
		context.JSON(http.StatusOK, obj)
	})

	r.LoadHTMLGlob("templates/openplatform-auth.html")
	r.GET("/openplatform-auth", func(c *gin.Context) {
		c.HTML(http.StatusOK, "openplatform-auth.html", nil)
	})

}
