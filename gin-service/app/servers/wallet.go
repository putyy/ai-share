package servers

import (
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/models"
	"github.com/putyy/ai-share/config"
)

func WalletInfo(uid int) *models.Wallet {
	wallet := &models.Wallet{}
	models.Db().Where("uid=?", uid).First(wallet)
	return wallet
}

func WalletLogList(form form.LogApiForm, uid int) (resList []models.WalletLog) {
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
