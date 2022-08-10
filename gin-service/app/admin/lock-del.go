package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/models"
	"gorm.io/gorm/schema"
)

var modelMaps = map[string]schema.Tabler{}

func init() {
	modelMaps["share-resource"] = models.ShareResource{}
	modelMaps["share-resource-type"] = models.ShareResourceType{}
	modelMaps["ppt-type"] = models.PptType{}
	modelMaps["ppt"] = models.Ppt{}
	modelMaps["menu"] = models.Menu{}
	modelMaps["menu-group"] = models.MenuGroup{}
}

func DeleteModel(c *gin.Context) {
	id := c.PostForm("id")
	model, ok := modelMaps[c.PostForm("m_mark")]

	if id == "" || ok == false {
		api.ResponseError(c, "参数错误", "")
		return
	}
	err := models.Db().Where("id = ?", id).Delete(&model).Error
	if err != nil {
		api.ResponseError(c, err.Error(), "")
		return
	}
	api.ResponseSuccess(c, "")
}

func LockModel(c *gin.Context) {
	id := c.PostForm("id")
	isLock := c.PostForm("is_lock")
	model, ok := modelMaps[c.PostForm("m_mark")]

	if id == "" || ok == false {
		api.ResponseError(c, "参数错误", "")
		return
	}

	err := models.Db().Table(model.TableName()).Where("id = ?", id).Update("is_lock", isLock).Error

	if err != nil {
		api.ResponseError(c, err.Error(), "")
		return
	}
	api.ResponseSuccess(c, "")
}
