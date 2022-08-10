package api

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/servers"
	"strconv"
)

func PptType(c *gin.Context) {
	data := make(map[string]interface{})
	data["list"] = servers.PpTypeList()
	ResponseSuccess(c, data)
}

func PptList(c *gin.Context) {
	formData := form.PptListApiForm{}
	_ = c.ShouldBindQuery(&formData)

	data := make(map[string]interface{})
	data["list"] = servers.PptList(formData, "id,name,img_url,desc_content")
	ResponseSuccess(c, data)
}

func PptContent(c *gin.Context) {
	data := make(map[string]interface{})
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		ResponseError(c, "id 必须", "")
		return
	}

	ppt := servers.PptInfo(id)
	if ppt.IsLock == 2 {
		ResponseError(c, "ppt信息有误", "")
		return
	}

	record := servers.AiDouLogSourceInfo(GetLoginUid(c), id)
	data["file_url"] = ""
	data["is_use"] = false
	if record.ID > 0 {
		data["file_url"] = ppt.FileUrl
		data["is_use"] = true
	}

	data["id"] = id
	data["ai_dou"] = ppt.AiDou
	data["img_url"] = ppt.ImgUrl
	data["name"] = ppt.Name
	data["desc_content"] = ppt.DescContent
	data["contents"] = servers.PptContents(id)
	ResponseSuccess(c, data)
}

func PptBuy(c *gin.Context) {
	data := make(map[string]interface{})
	id, _ := strconv.Atoi(c.PostForm("id"))
	if id <= 0 {
		ResponseError(c, "id 必须", "")
		return
	}
	ppt, err := servers.BuyPpt(id, GetLoginUid(c))
	if err != nil {
		ResponseError(c, err.Error(), "")
		return
	}
	data["file_url"] = ppt.FileUrl
	ResponseSuccess(c, data)
}
