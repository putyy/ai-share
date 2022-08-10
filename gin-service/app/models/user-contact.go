package models

// 用户关系
type UserContact struct {
	Uid         int      `gorm:"primary_key" json:"uid,omitempty"`
	Superior    int      `json:"superior,omitempty"`
	SuperiorTwo int      `json:"superior_two"`
	UpdatedAt   JsonTime `json:"updated_at"`
}

func (UserContact) TableName() string {
	return "a_user_contact"
}
