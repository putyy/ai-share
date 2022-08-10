package servers

import (
	"github.com/putyy/ai-share/app/models"
	"github.com/putyy/ai-share/config"
)

func ShareResourceList(id, tid int) (resList []models.ShareResource) {
	model := models.Db().Where("id > ?", id).Where("is_lock = ?", 1)
	if tid > 0 {
		model = model.Where("tid = ?", tid)
	}
	model.Order("sort desc").Limit(config.App.PageSize).Find(&resList)
	return
}

func ShareResourceTypeList() (resList []models.ShareResourceType) {
	models.Db().Where("is_lock = ?", 1).Order("sort desc, id desc").Find(&resList)
	return
}
