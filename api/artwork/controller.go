package artwork

import (
	"strings"

	"github.com/gin-gonic/gin"

	"tdp-aiart/model/artwork"
	"tdp-aiart/model/user"
	"tdp-aiart/module/painter"
)

// 创作列表

func list(c *gin.Context) {

	var rq *artwork.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	// 非管理员只能查看自己或公开作品
	if c.GetUint("UserLevel") != 1 {
		if rq.UserId == 0 || rq.UserId != c.GetUint("UserId") {
			rq.Status = "public"
		}
	}

	if lst, err := artwork.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Items": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取创作

func detail(c *gin.Context) {

	var rq *artwork.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if res, err := artwork.Fetch(rq); err == nil {
		c.Set("Payload", gin.H{"Item": res})
	} else {
		c.Set("Error", err)
	}

}

// 添加创作

func create(c *gin.Context) {

	// 构造参数

	param := &painter.ReqeustParam{}

	if err := c.ShouldBindJSON(param); err != nil {
		c.Set("Error", err)
		return
	}

	// 验证配额

	userId := c.GetUint("UserId")
	ur, err := user.Fetch(&user.FetchParam{
		Id: userId,
	})

	if err != nil || ur.ArtworkQuota <= 0 {
		c.Set("Error", "可用配额不足")
		return
	}

	user.UpdateQuota(&user.UpdateQuotaParam{Id: userId, ArtworkQuota: -1})

	// 请求接口

	res, err := painter.Create(param)

	if err != nil {
		user.UpdateQuota(&user.UpdateQuotaParam{Id: userId, ArtworkQuota: 1})
		c.Set("Error", err)
		return
	}

	// 存储数据

	rq := &artwork.CreateParam{
		UserId:         userId,
		Subject:        param.Subject,
		Prompt:         param.Prompt,
		NegativePrompt: param.NegativePrompt,
		Styles:         strings.Join(param.Styles, ","),
		Strength:       param.Strength,
		InputFile:      res.InputFile,
		OutputFile:     res.OutputFile,
		Status:         param.Status,
	}

	if id, err := artwork.Create(rq); err == nil {
		c.Set("Payload", gin.H{"Id": id, "OutputFile": res.OutputFile})
		c.Set("Message", "添加成功")
	} else {
		user.UpdateQuota(&user.UpdateQuotaParam{Id: userId, ArtworkQuota: 1})
		c.Set("Error", err)
	}

}

// 修改创作

func update(c *gin.Context) {

	var rq *artwork.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := artwork.Update(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除创作

func delete(c *gin.Context) {

	var rq *artwork.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := artwork.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
