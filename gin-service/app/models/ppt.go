package models

import (
	"github.com/putyy/ai-share/app/library"
	"gorm.io/gorm"
)

// ppt
type Ppt struct {
	ID          int      `gorm:"primary_key" json:"id,omitempty"`
	Tid         int      `json:"tid,omitempty"`
	Name        string   `json:"name,omitempty"`
	ImgUrl      string   `json:"img_url"`
	DescContent string   `json:"desc_content,omitempty"`
	Sort        int      `json:"sort"`
	FileUrl     string   `json:"file_url"`
	AiDou       int      `json:"ai_dou,omitempty"`
	IsLock      int      `json:"is_lock,omitempty" gorm:"default:1"`
	AdminUid    int      `json:"admin_uid,omitempty"`
	CreatedAt   JsonTime `json:"created_at,omitempty"`
	UpdatedAt   JsonTime `json:"updated_at,omitempty"`
}

func (Ppt) TableName() string {
	return "a_ppt"
}

func (t *Ppt) AfterFind(tx *gorm.DB) (err error) {
	t.ImgUrl = library.GetImgUrl(t.ImgUrl)
	t.FileUrl = library.GetFileUrl(t.FileUrl)
	if t.AiDou == 0 {
		t.AiDou = 100
	}

	return
}
