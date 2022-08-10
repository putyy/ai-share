package servers

import (
	"errors"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/models"
	"github.com/putyy/ai-share/config"
	"gorm.io/gorm"
	"time"
)

func OpenVip(formData form.OpenVipAdminForm, adminUid int) (map[string]interface{}, error) {
	ordSn, err := library.MakeOrderNo(formData.Uid)
	if err != nil {
		return nil, errors.New("订单创建失败")
	}

	user := UserInfo(map[string]interface{}{"id": formData.Uid}, "id,vip,vip_end_at")
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}

	if user.Vip >= 2 && user.Vip > formData.Level {
		return nil, errors.New("用户等级大于或等于开通等级")
	}

	vipEndAt := ""
	today := time.Now()
	if user.VipEndAt.IsZero() || user.VipEndAt.Before(today) {
		user.VipEndAt = models.JsonTime{Time: today}
	}

	if config.Vip[formData.Level].Length == -1 {
		vipEndAt = user.VipEndAt.AddDate(10, 0, 0).Format("2006-01-02 15:04:05")
	} else {
		vipEndAt = user.VipEndAt.AddDate(0, 0, config.Vip[formData.Level].Length*30).Format("2006-01-02 15:04:05")
	}

	if formData.ShowPrice > config.Vip[formData.Level].Price {
		return nil, errors.New("填写金额不能大于设置金额")
	}

	if formData.ShowPrice > 0 {
		formData.PayType = 2
	}

	res := make(map[string]interface{})

	err1 := models.Db().Transaction(func(tx *gorm.DB) error {
		order := models.Order{
			Uid:          formData.Uid,
			OrderNo:      ordSn,
			Scene:        1,
			TotalPrice:   formData.Price,
			ActualAmount: formData.Price,
			PayType:      formData.PayType,
			PayStatus:    2,
			Remark:       formData.Remark + ":购买:《" + config.Vip[formData.Level].Name + "》",
			AdminUid:     adminUid,
		}
		err2 := tx.Create(&order).Error

		if err2 != nil {
			// 返回任何错误都会回滚事务
			return err2
		}

		err3 := tx.Create(&models.OrderVip{
			Oid:   order.ID,
			Level: formData.Level,
		}).Error

		if err3 != nil {
			// 返回任何错误都会回滚事务
			return err3
		}

		err4 := tx.Model(&models.User{}).Where("id = ?", formData.Uid).Updates(map[string]interface{}{
			"vip":        formData.Level,
			"vip_end_at": vipEndAt,
		}).Error
		if err4 != nil {
			// 返回任何错误都会回滚事务
			return err4
		}
		var parentUserInfo struct {
			Superior int `json:"superior,omitempty"`
			Vip      int `json:"vip,omitempty"`
		}
		// 分钱 查找上级及上级的等级
		tx.Table(models.UserContact{}.TableName()+" a").Select("a.superior,b.vip").
			Joins("left join "+models.User{}.TableName()+" b on a.superior = b.id").
			Where("a.uid = ?", formData.Uid).Find(&parentUserInfo)

		benefit := 0
		if parentUserInfo.Vip > 0 && config.Vip[parentUserInfo.Vip].Profit > 0 && parentUserInfo.Vip > formData.Level {
			benefit = int(float64(config.Vip[parentUserInfo.Vip].Profit) / 100 * float64(formData.Price))
		}

		if benefit > 0 {
			err5 := tx.Model(&models.Wallet{}).Where("uid =?", parentUserInfo.Superior).Updates(map[string]interface{}{
				"balance":       gorm.Expr("balance + ?", benefit),
				"total_balance": gorm.Expr("total_balance + ?", benefit),
			}).Error
			if err5 != nil {
				return err5
			}

			err6 := tx.Create(&models.WalletLog{
				Uid:       parentUserInfo.Superior,
				Oid:       order.ID,
				Source:    1,
				Balance:   benefit,
				Content:   "《下级开通会员获益》",
				Direction: 1,
			}).Error

			if err6 != nil {
				// 返回任何错误都会回滚事务
				return err6
			}
		}

		res["vip"] = formData.Level
		res["vip_end_at"] = vipEndAt
		return nil
	})

	return res, err1
}

func RechargeAiDou(formData form.RechargeAiDouAdminForm, adminUid int) error {
	ordSn, err := library.MakeOrderNo(formData.Uid)
	if err != nil {
		return errors.New("订单创建失败")
	}

	user := UserInfo(map[string]interface{}{"id": formData.Uid}, "id")
	if user.ID == 0 {
		return errors.New("用户不存在")
	}

	if formData.Price > 0 {
		formData.PayType = 2
	}

	err1 := models.Db().Transaction(func(tx *gorm.DB) error {
		order := models.Order{
			Uid:          formData.Uid,
			OrderNo:      ordSn,
			Scene:        2,
			TotalPrice:   formData.Price,
			ActualAmount: formData.Price,
			PayType:      formData.PayType,
			PayStatus:    2,
			Remark:       formData.Remark + ":购买:《充值爱享豆》",
			AdminUid:     adminUid,
		}
		err2 := tx.Create(&order).Error

		if err2 != nil {
			return err2
		}

		err22 := tx.Create(&models.OrderAiDou{
			Oid:   order.ID,
			AiDou: formData.AiDou,
		}).Error

		if err22 != nil {
			return err22
		}

		err3 := tx.Model(&models.AiDou{}).Where("uid =?", formData.Uid).Updates(map[string]interface{}{
			"ai_dou":       gorm.Expr("ai_dou + ?", formData.AiDou),
			"total_ai_dou": gorm.Expr("total_ai_dou + ?", formData.AiDou),
		}).Error
		if err3 != nil {
			return err3
		}

		err4 := tx.Create(&models.AiDouLog{
			Uid:       formData.Uid,
			Source:    3,
			Sid:       order.ID,
			AiDou:     formData.AiDou,
			Content:   "《充值爱享豆》",
			Direction: 1,
		}).Error

		if err4 != nil {
			return err4
		}
		return nil
	})
	return err1
}
