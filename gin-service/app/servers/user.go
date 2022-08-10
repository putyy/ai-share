package servers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/models"
	"github.com/putyy/ai-share/config"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	"gorm.io/gorm"
	"time"
)

func UserInfo(m map[string]interface{}, field string) (res models.User) {
	models.Db().Unscoped().Where(m).Select(field).Find(&res)
	return
}

func UserEdit(formData form.UserApiForm) {
	u := models.User{
		NickName: formData.NickName,
		UserName: formData.UserName,
		HeadImg:  library.GetSaveUrl(formData.HeadImg),
	}
	models.Db().Model(&u).Where("id = ?", formData.Uid).Updates(u)
}

func UserRegister(wxUser auth.ResCode2Session, fromUid int) (int, error) {
	u := models.User{
		OpenID:  wxUser.OpenID,
		UnionID: wxUser.UnionID,
	}
	err := models.Db().Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&u).Error
		if err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		SuperiorTwo := 1
		if fromUid > 1 {
			uc := &models.UserContact{}
			tx.Where("uid = ?", fromUid).Find(uc)
			if uc.SuperiorTwo > 0 {
				SuperiorTwo = uc.SuperiorTwo
			}
		}

		err1 := tx.Create(&models.UserContact{
			Uid:         u.ID,
			Superior:    fromUid,
			SuperiorTwo: SuperiorTwo,
		}).Error

		err3 := tx.Create(&models.AiDou{
			Uid: u.ID,
		}).Error

		err4 := tx.Create(&models.Wallet{
			Uid: u.ID,
		}).Error

		if err1 != nil || err3 != nil || err4 != nil {
			// 返回任何错误都会回滚事务
			return errors.New("注册失败")
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return u.ID, nil
}

func UserFriendCount(uid int) (res int64) {
	models.Db().Model(&models.UserContact{}).Where("superior = ?", uid).Count(&res)
	return
}

type userResult struct {
	ID        int             `json:"id,omitempty"`
	UserName  string          `json:"user_name"`
	NickName  string          `json:"nick_name"`
	HeadImg   string          `json:"head_img"`
	Vip       int             `json:"vip"`
	VipName   string          `json:"vip_name"`
	CreatedAt models.JsonTime `json:"created_at,omitempty"`
}

func UserFriendList(form form.FansListApiForm, uid int) (resList []userResult) {
	user := models.User{}
	userContact := models.UserContact{}

	model := models.Db().Table(userContact.TableName()+" a").Select("b.id,b.user_name,b.nick_name,b.head_img,b.vip,b.created_at").
		Joins("left join "+user.TableName()+" b on a.uid = b.id").
		Where("a.superior = ?", uid)

	if form.Vip > -1 {
		model = model.Where("vip = ?", form.Vip)
	}

	if form.LastId > 0 {
		model = model.Where("uid > ?", form.LastId)
	}

	if form.Keyword != "" {
		model = model.Where("(nick_name like ? or user_name like ?)", "%"+form.Keyword+"%", "%"+form.Keyword+"%")
	}

	model.Scan(&resList)
	for k, v := range resList {
		resList[k].HeadImg = library.GetImgUrl(v.HeadImg)
		resList[k].VipName = config.Vip[v.Vip].Name
	}
	return
}

func UpdateUserCache(c *gin.Context, id string) {
	user := models.User{}
	models.Db().Unscoped().Where("id = ?", id).Select("id,is_lock,deleted_at").Find(&user)
	var deletedAt int64
	if user.DeletedAt.Valid == true {
		deletedAt = user.DeletedAt.Time.Unix()
	}
	rKey := library.BuildRdsKv("user_info").GetHashKey(id)
	library.Redis().HSet(c, rKey, map[string]interface{}{
		"is_lock":    user.IsLock,
		"deleted_at": deletedAt,
	})

	library.Redis().Expire(c, rKey, 172800*time.Second)
}
