package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/models"
	"github.com/putyy/ai-share/app/servers"
	"strconv"
)

type User struct {
	models.User
	UserContact models.UserContact `json:"user_contact,omitempty" gorm:"ForeignKey:uid;AssociationForeignKey:id"`
	Wallet      models.Wallet      `json:"wallet,omitempty" gorm:"ForeignKey:uid;AssociationForeignKey:id"`
	AiDou       models.AiDou       `json:"ai_dou,omitempty" gorm:"ForeignKey:uid;AssociationForeignKey:id"`
}

func UserList(c *gin.Context) {
	data := make(map[string]interface{})
	var resList []User
	model := models.Db()

	var total int64

	if c.Query("id") != "" {
		model = model.Where("id = ?", c.Query("id"))
	}

	if c.Query("user_name") != "" {
		model = model.Where("user_name = ?", c.Query("user_name"))
	}

	if c.Query("nick_name") != "" {
		model = model.Where("nick_name LIKE ?", "%"+c.Query("nick_name")+"%")
	}

	if c.Query("is_lock") != "" {
		model = model.Where("is_lock = ?", c.Query("is_lock"))
	}

	if c.Query("created_at_start") != "" {
		model = model.Where("created_at >= ?", c.Query("created_at_start"))
	}

	if c.Query("created_at_end") != "" {
		model = model.Where("created_at < ?", c.Query("created_at_end"))
	}

	if c.Query("total") != "" {
		total, _ = strconv.ParseInt(c.Query("total"), 10, 64)
	}

	model.Scopes(models.PaginateScope(c)).Order("id desc").
		Preload("UserContact").
		Preload("Wallet").
		Preload("AiDou").
		Find(&resList)

	if total == 0 {
		model.Model(models.User{}).Count(&total)
	}

	data["total"] = total
	data["page"] = c.Query("page")
	data["page_size"] = c.Query("page_size")
	data["list"] = resList
	api.ResponseSuccess(c, data)
}

func UserLocked(c *gin.Context) {
	id := c.PostForm("id")
	isLock := c.PostForm("is_lock")
	if id == "" {
		api.ResponseError(c, "参数错误", "")
		return
	}

	err := models.Db().Table(User{}.TableName()).Where("id = ?", id).Update("is_lock", isLock).Error
	if err != nil {
		api.ResponseError(c, err.Error(), "")
		return
	}
	servers.UpdateUserCache(c, id)
	api.ResponseSuccess(c, "")
}

func UserDeleted(c *gin.Context) {
	id := c.PostForm("id")

	if id == "" {
		api.ResponseError(c, "参数错误", "")
		return
	}

	err := models.Db().Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		api.ResponseError(c, err.Error(), "")
		return
	}

	servers.UpdateUserCache(c, id)
	api.ResponseSuccess(c, "")
}

func UserOpenVip(c *gin.Context) {
	var formData form.OpenVipAdminForm
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}

	res, err2 := servers.OpenVip(formData, GetAdminUid(c))
	if err2 != nil {
		api.ResponseError(c, "参数错误", err2.Error())
		return
	}
	api.ResponseSuccess(c, res)
}

func RechargeAiDou(c *gin.Context) {
	var formData form.RechargeAiDouAdminForm
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}
	err2 := servers.RechargeAiDou(formData, GetAdminUid(c))
	if err2 != nil {
		api.ResponseError(c, "参数错误", err2.Error())
		return
	}
	api.ResponseSuccess(c, "")
}
