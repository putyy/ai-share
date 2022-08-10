package models

// 文本库
type Content struct {
	ID        int      `gorm:"primary_key" json:"id,omitempty"`
	Scene     int      `json:"scene,omitempty"`
	Name      string   `json:"name,omitempty"`
	Content   string   `json:"content,omitempty"`
	AdminUid  int      `json:"admin_uid,omitempty"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
	UpdatedAt JsonTime `json:"updated_at,omitempty"`
}

func (Content) TableName() string {
	return "a_content"
}
