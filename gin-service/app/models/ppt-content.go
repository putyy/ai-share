package models

import (
	"github.com/putyy/ai-share/app/library"
	"gorm.io/gorm"
)

// ppt内容
type PptContent struct {
	ID      int    `gorm:"primary_key" json:"id,omitempty"`
	Pid     int    `json:"pid,omitempty"`
	ImgUrl  string `json:"img_url"`
	Content string `json:"content"`
}

func (PptContent) TableName() string {
	return "a_ppt_content"
}

func (t *PptContent) AfterFind(tx *gorm.DB) (err error) {
	t.ImgUrl = library.GetImgUrl(t.ImgUrl)
	return
}
