package models

import (
	"github.com/putyy/ai-share/app/library"
	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"primary_key" json:"id,omitempty"`
	UserName  string         `json:"user_name"`
	Password  string         `json:"password,omitempty"`
	OpenID    string         `gorm:"column:openid" json:"openid,omitempty"`
	UnionID   string         `gorm:"column:unionid" json:"unionid,omitempty"`
	NickName  string         `json:"nick_name"`
	HeadImg   string         `json:"head_img"`
	Vip       int            `json:"vip" gorm:"default:0"`
	IsLock    int            `json:"is_lock,omitempty" gorm:"default:1"`
	VipEndAt  JsonTime       `json:"vip_end_at,omitempty"`
	CreatedAt JsonTime       `json:"created_at,omitempty"`
	UpdatedAt JsonTime       `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (User) TableName() string {
	return "a_user"
}

func (t *User) AfterFind(tx *gorm.DB) (err error) {
	t.HeadImg = library.GetImgUrl(t.HeadImg)
	return
}
