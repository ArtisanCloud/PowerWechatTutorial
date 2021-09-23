package wecom

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
  "strconv"
)

func APIAgentGet(c *gin.Context) {

  agentId := c.DefaultQuery("agentId", "AGENTID")
  agentID, _ := strconv.Atoi(agentId)
  res, err := services.WeComApp.Agent.Get(agentID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentList(c *gin.Context) {
  res, err := services.WeComApp.Agent.List()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentSet(c *gin.Context) {
  agentId := c.DefaultQuery("agentId", "AGENTID")
  agentID, _ := strconv.Atoi(agentId)

  data := &power.HashMap{
    "agentid":              agentID,
    "report_location_flag": 0,
    "logo_mediaid":         "j5Y8X5yocspvBHcgXMSS6z1Cn9RQKREEJr4ecgLHi4YHOYP-plvom-yD9zNI0vEl",
    "name":                 "财经助手",
    "description":          "内部财经服务平台",
    "redirect_domain":      "open.work.weixin.qq.com",
    "isreportenter":        0,
    "home_url":             "https://open.work.weixin.qq.com",
  }

  res, err := services.WeComApp.Agent.Set(agentID, data)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentMenuCreate(c *gin.Context) {
  data := &power.HashMap{
    "button": []power.HashMap{
      {
        "type": "click",
        "name": "今日歌曲",
        "key":  "V1001_TODAY_MUSIC",
      },
      {
        "name": "菜单",
        "sub_button": []power.HashMap{
          {
            "type": "view",
            "name": "搜索",
            "url":  "http://www.soso.com/",
          },
          {
            "type": "click",
            "name": "赞一下我们",
            "key":  "V1001_GOOD",
          },
        },
      },
    },
  }
  res, err := services.WeComApp.Menu.Create(data)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentMenuGet(c *gin.Context) {
  res, err := services.WeComApp.Menu.Get()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentMenuDelete(c *gin.Context) {

  agentId := c.DefaultQuery("agentId", "AGENTID")
  agentID, _ := strconv.Atoi(agentId)

  res, err := services.WeComApp.Menu.Delete(agentID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentSetWorkbenchTemplate(c *gin.Context) {

  data := &power.HashMap{

    "items": []power.HashMap{
      {
        "key":      "待审批",
        "data":     "2",
        "jump_url": "http://www.qq.com",
        "pagepath": "pages/index",
      },
      {
        "key":      "带批阅作业",
        "data":     "4",
        "jump_url": "http://www.qq.com",
        "pagepath": "pages/index",
      },
      {
        "key":      "成绩录入",
        "data":     "45",
        "jump_url": "http://www.qq.com",
        "pagepath": "pages/index",
      },
      {
        "key":      "综合评价",
        "data":     "98",
        "jump_url": "http://www.qq.com",
        "pagepath": "pages/index",
      },
    },
  }

  res, err := services.WeComApp.AgentWorkbench.SetWorkbenchTemplate(data)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentGetWorkbenchTemplate(c *gin.Context) {

  agentId := c.DefaultQuery("agentId", "AGENTID")
  agentID, _ := strconv.Atoi(agentId)

  res, err := services.WeComApp.AgentWorkbench.GetWorkbenchTemplate(agentID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentSetWorkbenchData(c *gin.Context) {
  agentId := c.DefaultQuery("agentId", "AGENTID")
  agentID, _ := strconv.Atoi(agentId)

  data := &power.HashMap{
    "agentid": agentID,
    "type":    "image",
    "image": power.HashMap{
      "url":      "xxxx",
      "jump_url": "http://www.qq.com",
      "pagepath": "pages/index",
    },
    "replace_user_data": true,
  }
  res, err := services.WeComApp.AgentWorkbench.SetWorkbenchData(data)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
