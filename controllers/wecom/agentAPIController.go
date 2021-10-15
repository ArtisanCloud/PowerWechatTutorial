package wecom

import (
  "github.com/ArtisanCloud/power-wechat/src/work/agent/request"
  request2 "github.com/ArtisanCloud/power-wechat/src/work/menu/request"
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

  options := &request.RequestAgentSet{
    AgentID:            agentID,
    ReportLocationFlag: 0,
    LogoMediaID:        "j5Y8X5yocspvBHcgXMSS6z1Cn9RQKREEJr4ecgLHi4YHOYP-plvom-yD9zNI0vEl",
    Name:               "财经助手",
    Description:        "内部财经服务平台",
    RedirectDomain:     "open.work.weixin.qq.com",
    IsReportEnter:      0,
    HomeUrl:            "https://open.work.weixin.qq.com",
  }

  res, err := services.WeComApp.Agent.Set(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentMenuCreate(c *gin.Context) {
  options := &request2.RequestMenuSet{
    Button: []request2.RequestMenuSetButton{
      {
        Type: "click",
        Name: "今日歌曲",
        Key:  "V1001_TODAY_MUSIC",
      },
      {
        Name: "菜单",
        SubButton: []request2.RequestMenuSetButton{
          {
            Type: "view",
            Name: "搜索",
            Url:  "http://www.soso.com/",
          },
          {
            Type: "click",
            Name: "赞一下我们",
            Key:  "V1001_GOOD",
          },
        },
      },
    },
  }
  res, err := services.WeComApp.Menu.Create(options)

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
  agentId := c.DefaultQuery("agentId", "1000005")
  agentID, _ := strconv.Atoi(agentId)

  options := &request.RequestSetWorkbenchTemplate{
    AgentID: agentID,
    Type:    "image",
    Image: request.WorkBenchImage{
      Url:      "xxxx",
      JumpUrl:  "http://www.qq.com",
      PagePath: "pages/index",
    },
    ReplaceUserData: true,
  }
  res, err := services.WeComApp.AgentWorkbench.SetWorkbenchTemplate(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentGetWorkbenchTemplate(c *gin.Context) {

  agentId := c.DefaultQuery("agentId", "1000005")
  agentID, _ := strconv.Atoi(agentId)

  res, err := services.WeComApp.AgentWorkbench.GetWorkbenchTemplate(agentID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAgentSetWorkbenchData(c *gin.Context) {
  agentId := c.DefaultQuery("agentId", "1000005")
  agentID, _ := strconv.Atoi(agentId)

  options := &request.RequestSetWorkBenchData{
    AgentID: agentID,
    UserID:  "test",
    Type:    "keydata",
    KeyData: request.WorkBenchKeyData{
      Items: []request.WorkBenchKeyDataItem{
        {
          Key:      "待审批",
          Data:     "2",
          JumpUrl:  "http://www.qq.com",
          PagePath: "pages/index",
        },
        {
          Key:      "带批阅作业",
          Data:     "4",
          JumpUrl:  "http://www.qq.com",
          PagePath: "pages/index",
        },
        {
          Key:      "成绩录入",
          Data:     "45",
          JumpUrl:  "http://www.qq.com",
          PagePath: "pages/index",
        },
        {
          Key:      "综合评价",
          Data:     "98",
          JumpUrl:  "http://www.qq.com",
          PagePath: "pages/index",
        },
      },
    },
  }
  res, err := services.WeComApp.AgentWorkbench.SetWorkbenchData(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
