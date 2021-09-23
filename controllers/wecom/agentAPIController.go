package wecom

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func APIAgentGet(c *gin.Context) {

  agentId := c.DefaultQuery("agentId", "AGENTID")
  agentID :=
  res, err := services.WeComApp.Agent.Get(agentId)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentList(c *gin.Context) {

}

func APIAgentSet(c *gin.Context) {

}



func APIAgentMenuCreate(c *gin.Context) {

}

func APIAgentMenuGet(c *gin.Context) {

}

func APIAgentMenuDelete(c *gin.Context) {

}


func APIAgentSetWorkbenchTemplate(c *gin.Context) {

}

func APIAgentGetWorkbenchTemplate(c *gin.Context) {

}

func APIAgentSetWorkbenchData(c *gin.Context) {


}