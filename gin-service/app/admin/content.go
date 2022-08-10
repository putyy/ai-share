package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/models"
)

func ContentList(c *gin.Context) {
	data := make(map[string]interface{})
	var resList []models.Content
	model := models.Db()

	var total int64

	model.Scopes(models.PaginateScope(c)).Order("id desc").Find(&resList)
	if total == 0 {
		model.Model(models.Content{}).Count(&total)
	}

	data["total"] = total
	data["page"] = c.Query("page")
	data["page_size"] = c.Query("page_size")
	data["list"] = resList
	data["scene"] = library.TextScenes
	api.ResponseSuccess(c, data)
}

func ContentEdit(c *gin.Context) {
	var formData form.ContentAdminForm
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}

	model := models.Content{
		Scene:   formData.Scene,
		Name:    formData.Name,
		Content: formData.Content,
	}

	var err error
	if formData.Id > 0 {
		err = models.Db().Model(&model).Where("id = ?", formData.Id).Updates(model).Error
	} else {
		err = models.Db().Create(&model).Error
	}
	if err != nil {
		api.ResponseError(c, "创建失败", err.Error())
		return
	}

	api.ResponseSuccess(c, model)
}
