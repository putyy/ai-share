package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/library"
)

func SystemInfoDetail(c *gin.Context) {
	systemInfoHKey := library.BuildRdsKv("system_info").GetHashKey()
	systemLockData, err1 := library.Redis().HGetAll(c, systemInfoHKey).Result()
	if err1 != nil || len(systemLockData) == 0 {
		api.ResponseSuccess(c, map[string]string{
			"mini_check":           "1", // 1正常  2小程序审核中
			"system_close":         "1", // 1系统正常 2关站中
			"system_close_content": "",  // 关站文案
			"customer_service_wx":  "",  // 客服微信
		})
		return
	}
	systemLockData["customer_service_wx"] = library.GetImgUrl(systemLockData["customer_service_wx"])
	api.ResponseSuccess(c, systemLockData)
}

func SystemInfoEdit(c *gin.Context) {
	systemInfoHKey := library.BuildRdsKv("system_info").GetHashKey()
	library.Redis().HSet(c, systemInfoHKey, map[string]interface{}{
		"mini_check":           c.PostForm("mini_check"),
		"system_close":         c.PostForm("system_close"),
		"system_close_content": c.PostForm("system_close_content"),
		"customer_service_wx":  library.GetSaveUrl(c.PostForm("customer_service_wx")),
		"admin_uid":            GetAdminUid(c),
	})
	api.ResponseSuccess(c, "")
}
