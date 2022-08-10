package router

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/admin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/middleware"
	"github.com/putyy/ai-share/config"
)

func InitRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger()).Use(gin.Recovery())
	gin.SetMode(config.App.RunMode)

	r.GET("/api/auth", api.Login)
	r.GET("/api/wx-auth", api.WxMiNiLogin)

	r.GET("/api/share-resource/type", api.ResourceType)
	r.GET("/api/share-resource/list", api.ResourceList)

	apiGroup := r.Group("/api").Use(middleware.RateLimiter()).Use(middleware.Api())
	{
		apiGroup.GET("member-center/center-info", api.CenterInfo)
		apiGroup.GET("member-center/user", api.User)
		apiGroup.POST("member-center/user-edit", api.UserEdit)
		apiGroup.GET("member-center/ai-dou-log", api.AiDouLogList)
		apiGroup.GET("member-center/wallet-log", api.WalletLogList)
		apiGroup.GET("member-center/fans-list", api.FansList)
		apiGroup.POST("member-center/feedback-create", api.FeedbackCreate)
		apiGroup.GET("member-center/feedback-list", api.FeedbackList)

		apiGroup.GET("ppt/type", api.PptType)
		apiGroup.GET("ppt/list", api.PptList)
		apiGroup.POST("ppt/buy", api.PptBuy)
		apiGroup.GET("ppt/detail", api.PptContent)

		apiGroup.GET("public/scene", api.Scene)
		apiGroup.GET("public/qiniu-token", api.GetQiNiuWebToken)
		apiGroup.GET("public/content", api.GetContent)

		apiGroup.POST("video-parse/short", api.ShortVideoParse)
	}

	r.GET("/admin/login", admin.Login)
	adminRoute := r.Group("/admin").Use(middleware.Admin())
	{
		adminRoute.GET("admin-user-info", admin.AdminUserInfo)

		adminRoute.POST("__lock", admin.LockModel)
		adminRoute.POST("__delete", admin.DeleteModel)

		adminRoute.GET("public/vip-config", admin.VipConfig)
		adminRoute.GET("public/scene", api.Scene)
		adminRoute.GET("public/qiniu-token", api.GetQiNiuWebToken)

		adminRoute.GET("ppt/list", admin.PptList)
		adminRoute.POST("ppt/edit", admin.PptEdit)
		adminRoute.GET("ppt/content", admin.PptContent)
		adminRoute.POST("ppt/content-edit", admin.PptContentEdit)

		adminRoute.GET("ppt-type/list", admin.PptTypeList)
		adminRoute.POST("ppt-type/edit", admin.PptTypeEdit)

		adminRoute.GET("share-resource/list", admin.ShareResourceList)
		adminRoute.POST("share-resource/edit", admin.ShareResourceEdit)

		adminRoute.GET("share-resource-type/list", admin.ShareResourceTypeList)
		adminRoute.POST("share-resource-type/edit", admin.ShareResourceTypeEdit)

		adminRoute.GET("feedback/list", admin.FeedbackList)
		adminRoute.POST("feedback/remark", admin.FeedbackRemark)

		adminRoute.GET("content/list", admin.ContentList)
		adminRoute.POST("content/edit", admin.ContentEdit)

		adminRoute.GET("menu/list", admin.MenuList)
		adminRoute.POST("menu/edit", admin.MenuEdit)

		adminRoute.GET("menu-group/list", admin.MenuGroupList)
		adminRoute.POST("menu-group/edit", admin.MenuGroupEdit)

		adminRoute.GET("user/list", admin.UserList)
		adminRoute.POST("user/lock", admin.UserLocked)
		adminRoute.POST("user/deleted", admin.UserDeleted)
		adminRoute.POST("user/open-vip", admin.UserOpenVip)
		adminRoute.POST("user/recharge-ai-dou", admin.RechargeAiDou)

		adminRoute.GET("system-info/detail", admin.SystemInfoDetail)
		adminRoute.POST("system-info/edit", admin.SystemInfoEdit)
	}

	return r
}
