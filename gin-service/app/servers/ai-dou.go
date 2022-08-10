package servers

import (
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/models"
	"github.com/putyy/ai-share/config"
)

func AiDouInfo(uid int) (res models.AiDou) {
	models.Db().Where("uid=?", uid).Find(&res)
	return
}

func AiDouLogSourceInfo(uid int, sid int) (res models.AiDouLog) {
	models.Db().Where("uid=? and sid=?", uid, sid).Find(&res)
	return
}

func AiDouLogList(form form.LogApiForm, uid int) (resList []models.AiDouLog) {
	model := models.Db().Where("uid = ?", uid).
		Limit(config.App.PageSize).
		Order("id desc")
	if form.LastId > 0 {
		model = model.Where("id < ?", form.LastId)
	}
	if form.Direction != 0 {
		model = model.Where("direction = ?", form.Direction)
	}
	if form.Keyword != "" {
		model = model.Where("content like ?", "%"+form.Keyword+"%")
	}
	model.Find(&resList)
	return
}
