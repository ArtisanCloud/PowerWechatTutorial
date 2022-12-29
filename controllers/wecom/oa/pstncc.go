package oa

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIPSTNCCCall(c *gin.Context) {
	options := []string{
		c.DefaultQuery("userID", "matrix-x"),
	}

	res, err := services.WeComApp.OAPSTNCC.Call(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIPSTNCCGetStates(c *gin.Context) {

	calleeUserID := c.DefaultQuery("calleeUserID", "matrix-x")
	callID := c.DefaultQuery("userID", "matrix-x")

	res, err := services.WeComApp.OAPSTNCC.GetStates(c.Request.Context(), calleeUserID, callID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
