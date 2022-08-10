package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/models"
)

func MenuGroupList(c *gin.Context) {
	data := make(map[string]interface{})
	var resList []models.MenuGroup
	model := models.Db()

	model.Order("id desc").Find(&resList)

	data["list"] = resList
	data["m_mark"] = "menu-group"
	api.ResponseSuccess(c, data)
}

func MenuGroupEdit(c *gin.Context) {
	var formData form.MenuGroupAdminForm
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}

	model := models.MenuGroup{
		Name:     formData.Name,
		Sort:     formData.Sort,
		IsLock:   formData.IsLock,
		AdminUid: GetAdminUid(c),
	}

	var err error
	if formData.Id > 0 {
		err = models.Db().Model(&model).Where("id = ?", formData.Id).Updates(model).Error
		model.ID = formData.Id
	} else {
		err = models.Db().Create(&model).Error
	}
	if err != nil {
		api.ResponseError(c, "创建失败", err.Error())
		return
	}

	api.ResponseSuccess(c, model)
}
