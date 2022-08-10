package servers

import (
	"errors"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/models"
	"github.com/putyy/ai-share/config"
	"gorm.io/gorm"
)

func PptList(form form.PptListApiForm, field string) (resList []models.Ppt) {
	model := models.Db().Where("is_lock = ?", 1)

	if form.Keyword != "" {
		model = model.Where("desc_content like ?", "%"+form.Keyword+"%")
	}

	if form.TypeId > 0 {
		model = model.Where("tid = ?", form.TypeId)
	}

	model.Where("id > ?", form.LastId).Select(field).Limit(config.App.PageSize).Order("sort desc").Find(&resList)
	return
}

func PpTypeList() (resList []models.PptType) {
	models.Db().Where("is_lock = ?", 1).Order("sort desc").Find(&resList)
	return
}

func PptInfo(id int) (res models.Ppt) {
	models.Db().Where("id = ?", id).Find(&res)
	return
}

func PptContents(id int) (resList []models.PptContent) {
	models.Db().Where("pid = ?", id).Find(&resList)
	return
}

func BuyPpt(pptId int, uid int) (models.Ppt, error) {
	pptInfo := PptInfo(pptId)
	if pptInfo.ID <= 0 {
		return pptInfo, errors.New("资源不存在")
	}
	aiDouLog := AiDouLogSourceInfo(uid, pptId)
	if aiDouLog.ID > 0 {
		return pptInfo, errors.New("不能重复购买")
	}
	douModel := models.AiDou{}
	models.Db().Where("uid=?", uid).Find(&douModel)
	if douModel.AiDou < pptInfo.AiDou {
		return pptInfo, errors.New("~爱豆不足,充值请联系客服~")
	}

	err := models.Db().Transaction(func(tx *gorm.DB) error {

		err1 := tx.Model(&models.AiDou{}).Where("uid =?", uid).UpdateColumn("ai_dou", gorm.Expr("ai_dou - ?", pptInfo.AiDou)).Error
		if err1 != nil {
			return err1
		}

		err2 := tx.Create(&models.AiDouLog{
			Uid:       uid,
			Source:    1,
			Sid:       pptInfo.ID,
			AiDou:     pptInfo.AiDou,
			Content:   "购买PPT《" + pptInfo.Name + "》",
			Direction: 2,
		}).Error

		if err2 != nil {
			// 返回任何错误都会回滚事务
			return err2
		}
		return nil
	})

	if err == nil {
		return pptInfo, nil
	}

	return pptInfo, err
}
