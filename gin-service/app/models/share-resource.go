package models

import (
	"github.com/putyy/ai-share/app/library"
	"gorm.io/gorm"
)

// 朋友圈素材分类
type ShareResource struct {
	ID        int      `gorm:"primary_key" json:"id,omitempty"`
	Tid       int      `json:"tid,omitempty"`
	ImgUrl    string   `json:"img_url"`
	Content   string   `json:"content,omitempty"`
	Sort      int      `json:"sort"`
	IsLock    int      `json:"is_lock,omitempty" gorm:"default:1"`
	AdminUid  int      `json:"admin_uid,omitempty"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
	UpdatedAt JsonTime `json:"updated_at,omitempty"`
}

func (ShareResource) TableName() string {
	return "a_share_resource"
}

func (t *ShareResource) AfterFind(tx *gorm.DB) (err error) {
	t.ImgUrl = library.GetImgUrl(t.ImgUrl)
	return
}
