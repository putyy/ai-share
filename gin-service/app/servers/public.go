package servers

import (
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/models"
	"github.com/putyy/ai-share/config"
	"gorm.io/gorm"
)

func MenuList() (resList []models.MenuGroup) {
	models.Db().Where("is_lock = ?", 1).Order("sort desc").Preload("Menus", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort desc")
	}).Find(&resList)
	return
}

func FeedbackUserNoCheckCount(uid int) (res int64) {
	models.Db().Model(&models.Feedback{}).Where("uid = ? and admin_uid=0", uid).Count(&res)
	return
}

func FeedbackCreate(uid int, content string) {
	models.Db().Create(&models.Feedback{
		Uid:     uid,
		Content: content,
	})
}

func FeedbackList(form form.CommonForm, uid int) (resList []models.Feedback) {
	model := models.Db().Where("uid = ?", uid).
		Select("id,content,remark,created_at").
		Limit(config.App.PageSize).
		Order("id desc")
	if form.LastId > 0 {
		model = model.Where("id < ?", form.LastId)
	}
	model.Find(&resList)
	return
}

func GetContent(scene int) (res models.Content) {
	models.Db().Where("scene = ?", scene).Select("name,content").First(&res)
	return
}
