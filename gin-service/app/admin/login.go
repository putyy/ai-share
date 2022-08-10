package admin

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/models"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var adminUser models.AdminUser
	models.Db().Where("username = ?", username).Find(&adminUser)

	if adminUser.ID == 0 {
		api.ResponseError(c, "账号不存在", "")
		return
	}

	if adminUser.Password != fmt.Sprintf("%x", md5.Sum([]byte(password))) {
		api.ResponseError(c, "密码错误", "")
		return
	}

	token, err := library.GenerateAdminToken(adminUser.ID)

	if err != nil {
		api.ResponseError(c, "token生成失败", "")
		return
	}

	data := make(map[string]interface{})
	data["id"] = adminUser.ID
	data["token"] = token
	api.ResponseSuccess(c, data)
}

func AdminUserInfo(c *gin.Context) {
	var adminUser models.AdminUser
	models.Db().Where("id = ?", GetAdminUid(c)).Select("id,username,name,avatar").Find(&adminUser)
	api.ResponseSuccess(c, adminUser)
}
