package validate

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/work/user/request"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/config"
	http "power-wechat-tutorial/controllers"
)

func ValidateUserCreate(context *gin.Context) {
	var form request.RequestUserDetail

	if err := context.ShouldBind(&form); err != nil {
		if err := context.ShouldBindJSON(&form); err != nil {

			apiResponse := http.NewAPIResponse(context)
			apiResponse.SetCode(
				config.API_ERR_CODE_REQUEST_PARAM_ERROR,
				config.API_RETURN_CODE_ERROR,
				"", "").SetData(object.HashMap{
				"message": err.Error(),
			}).ThrowJSONResponse(context)
		}
	}

	context.Set("params", &form)
	context.Next()
}
