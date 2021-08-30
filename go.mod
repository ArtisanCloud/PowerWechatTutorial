module wechat-pay-test

go 1.16

replace github.com/ArtisanCloud/power-wechat => ../power-wechat

replace github.com/ArtisanCloud/go-libs => ../go-libs

require (
	github.com/ArtisanCloud/go-libs v1.1.2
	github.com/ArtisanCloud/power-wechat v1.1.0
	github.com/gin-gonic/gin v1.7.2
)
