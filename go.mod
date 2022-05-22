module power-wechat-tutorial

go 1.16

replace github.com/ArtisanCloud/PowerWeChat => ../PowerWeChat
replace github.com/ArtisanCloud/PowerLibs => ../PowerLibs
replace github.com/ArtisanCloud/PowerSocialite => ../PowerSocialite

require (
	github.com/ArtisanCloud/PowerLibs v1.3.3
	github.com/ArtisanCloud/PowerWeChat v1.3.1-0.20220520132835-d88a6ef53692
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/assert/v2 v2.0.1
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-redis/redis/v8 v8.11.3 // indirect
	github.com/golang-module/carbon v1.5.5
	github.com/google/uuid v1.3.0 // indirect
	github.com/jinzhu/configor v1.2.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f // indirect
	golang.org/x/sys v0.0.0-20210906170528-6f6e22806c34 // indirect
	golang.org/x/text v0.3.7
)
