package models

// 意见反馈
type Feedback struct {
	ID        int      `gorm:"primary_key" json:"id,omitempty"`
	Uid       int      `json:"uid,omitempty"`
	Content   string   `json:"content,omitempty"`
	Remark    string   `json:"remark"`
	AdminUid  int      `json:"admin_uid,omitempty" gorm:"default:0"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
	UpdatedAt JsonTime `json:"updated_at"`
}

func (Feedback) TableName() string {
	return "a_feedback"
}
