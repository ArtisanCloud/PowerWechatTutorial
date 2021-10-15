package oa

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/ArtisanCloud/power-wechat/src/work/oa/webdrive/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 新建空间
// https://work.weixin.qq.com/api/doc/90000/90135/93655
func APIWebDriveSpaceCreate(c *gin.Context) {
  options := &request.RequestWebDriveSpaceCreate{
    UserID:    c.DefaultQuery("userID", "USERID"),
    SpaceName: c.DefaultQuery("spaceName", "SPACE_NAME"),
    AuthInfo: []*power.HashMap{
      &power.HashMap{
        "type":   1,
        "userid": "USERID",
        "auth":   2,
      },
      &power.HashMap{
        "type":         2,
        "departmentid": "DEPARTMENTID",
        "auth":         1,
      },
    },
  }

  res, err := services.WeComApp.OAWebDrive.SpaceCreate(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 添加成员/部门
// https://work.weixin.qq.com/api/doc/90000/90135/93656
func APIWebDriveSpaceAcAdd(c *gin.Context) {

  options := &request.RequestWebDriveSpaceACLAdd{
    UserID:  c.DefaultQuery("userID", "USERID"),
    SpaceID: c.DefaultQuery("spaceName", "SPACE_NAME"),
    AuthInfo: []*power.HashMap{
      &power.HashMap{
        "type":   1,
        "userid": "USERID",
        "auth":   2,
      },
      &power.HashMap{
        "type":         2,
        "departmentid": "DEPARTMENTID",
        "auth":         1,
      },
    },
  }
  res, err := services.WeComApp.OAWebDrive.SpaceACLAdd(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取文件列表
// https://work.weixin.qq.com/api/doc/90000/90135/93657
func APIWebDriveFileList(c *gin.Context) {

  options := &request.RequestWebDriveFileList{
    UserID:   c.DefaultQuery("userID", "USERID"),
    SpaceID:  c.DefaultQuery("spaceID", "SPACEID"),
    FatherID: c.DefaultQuery("fatherID", "FATHERID"),
    SortType: 1,
    Start:    1,
    Limit:    20,
  }

  res, err := services.WeComApp.OAWebDrive.FileList(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 新增指定人
// https://work.weixin.qq.com/api/doc/90000/90135/93658
func APIWebDriveFileAclAdd(c *gin.Context) {

  options := &request.RequestWebDriveFileACLAdd{
    UserID: c.DefaultQuery("userID", "USERID"),
    FileID: c.DefaultQuery("fileID", "FILEID"),
    AuthInfo: []*power.HashMap{
      &power.HashMap{
        "type":   1,
        "userid": "USERID",
        "auth":   2,
      },
      &power.HashMap{
        "type":         2,
        "departmentid": "DEPARTMENTID",
        "auth":         1,
      },
    },
  }

  res, err := services.WeComApp.OAWebDrive.FileACLAdd(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
