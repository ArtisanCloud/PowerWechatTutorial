package request

import (
	"github.com/ArtisanCloud/PowerWeChat/src/work/oauth/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateRequestOAuthCallbackQRCode(context *gin.Context) {
	var form request.ParaOAuthCallbackQRCode

	//logger.Info("validate make reservation", nil)
	if err := context.ShouldBind(&form); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, "Lack of parameters!")
	}

	context.Set("params", form)

	context.Next()
}
