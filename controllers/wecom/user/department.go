package user

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
	"strconv"
)

const defaultDepartmentId = 3000001

// APIDepartmentCreate 创建部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90205
func APIDepartmentCreate(c *gin.Context) {
	name := c.DefaultQuery("name", "IT支持部")
	idStr := c.DefaultQuery("id", string(rune(defaultDepartmentId)))
	id, _ := strconv.Atoi(idStr)

	res, err := services.WeComContactApp.Department.Create(&power.HashMap{
		"name":     name,
		"parentid": 1,
		"id":       id,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// APIDepartmentUpdate 更新部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90206
func APIDepartmentUpdate(c *gin.Context) {
	name := c.DefaultQuery("name", "IT支持部1")
	idStr := c.DefaultQuery("id", string(rune(defaultDepartmentId)))
	id, _ := strconv.Atoi(idStr)
	res, err := services.WeComContactApp.Department.Update(id, &power.HashMap{
		"name":     name,
		"parentid": 1,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// APIDepartmentDelete 删除部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90207
func APIDepartmentDelete(c *gin.Context) {
	idStr := c.DefaultQuery("id", string(rune(defaultDepartmentId)))
	id, _ := strconv.Atoi(idStr)
	res, err := services.WeComContactApp.Department.Delete(id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// APIDepartmentList 获取部门列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/90208
func APIDepartmentList(c *gin.Context) {
	idStr := c.DefaultQuery("id", "0")
	id, _ := strconv.Atoi(idStr)
	// 0 表示获取公司所有部门
	res, err := services.WeComContactApp.Department.List(id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// APIDepartmentSimpleList 获取子部门ID列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/90208
func APIDepartmentSimpleList(c *gin.Context) {
	idStr := c.DefaultQuery("id", "0")
	id, _ := strconv.Atoi(idStr)
	// 0 表示获取企业所有部门
	res, err := services.WeComContactApp.Department.SimpleList(id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
