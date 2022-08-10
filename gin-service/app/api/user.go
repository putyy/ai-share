package api

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/servers"
	"github.com/putyy/ai-share/config"
	"time"
)

func CenterInfo(c *gin.Context) {
	systemInfoHKey := library.BuildRdsKv("system_info").GetHashKey()
	customerServiceWx, _ := library.Redis().HGet(c, systemInfoHKey, "customer_service_wx").Result()
	ResponseSuccess(c, map[string]interface{}{
		"list":                servers.MenuList(),
		"customer_service_wx": library.GetImgUrl(customerServiceWx),
		"vip_list":            config.Vip,
	})
}

func User(c *gin.Context) {
	user := servers.UserInfo(map[string]interface{}{"id": GetLoginUid(c)}, "id,user_name,nick_name,head_img,vip,created_at")
	data := make(map[string]interface{})
	data["user"] = user
	data["more"] = map[string]interface{}{
		"vip_name":     config.Vip[user.Vip].Name,
		"register_day": (time.Now().Unix() - user.CreatedAt.Unix()) / 86400,
		"wallet":       servers.WalletInfo(GetLoginUid(c)),
		"friends":      servers.UserFriendCount(GetLoginUid(c)),
		"dou":          servers.AiDouInfo(GetLoginUid(c)),
		"bg_img":       "https://www.putyy.com/uploads/article/20220322/PD6fjBJIiYRuW1183JFMcbJfsstzAvmNTaKTJsNc.png",
	}
	ResponseSuccess(c, data)
}

func UserEdit(c *gin.Context) {
	formData := form.UserApiForm{}
	if err1 := c.ShouldBind(&formData); err1 != nil {
		ResponseError(c, "参数错误", err1.Error())
		return
	}
	formData.Uid = GetLoginUid(c)
	servers.UserEdit(formData)
	ResponseSuccess(c, "")
}

func AiDouLogList(c *gin.Context) {
	formData := form.LogApiForm{}
	_ = c.ShouldBindQuery(&formData)
	data := make(map[string]interface{})
	data["list"] = servers.AiDouLogList(formData, GetLoginUid(c))
	ResponseSuccess(c, data)
}

func WalletLogList(c *gin.Context) {
	formData := form.LogApiForm{}
	_ = c.ShouldBindQuery(&formData)
	data := make(map[string]interface{})
	data["list"] = servers.WalletLogList(formData, GetLoginUid(c))
	ResponseSuccess(c, data)
}

func FansList(c *gin.Context) {
	formData := form.FansListApiForm{}
	_ = c.ShouldBindQuery(&formData)
	data := make(map[string]interface{})
	data["list"] = servers.UserFriendList(formData, GetLoginUid(c))
	ResponseSuccess(c, data)
}

func FeedbackList(c *gin.Context) {
	formData := form.CommonForm{}
	_ = c.ShouldBindQuery(&formData)
	data := make(map[string]interface{})
	data["list"] = servers.FeedbackList(formData, GetLoginUid(c))
	ResponseSuccess(c, data)
}

func FeedbackCreate(c *gin.Context) {
	uid := GetLoginUid(c)
	content := c.PostForm("content")
	if content == "" {
		ResponseError(c, "反馈内容不能为空", "")
		return
	}
	count := servers.FeedbackUserNoCheckCount(uid)
	if count >= 3 {
		ResponseError(c, "您反馈的问题未处理数太多，请联系客服稍后再试！", "")
		return
	}
	servers.FeedbackCreate(uid, c.PostForm("content"))
	ResponseSuccess(c, "")
}
