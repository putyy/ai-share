package models

// 后台用户
type AdminUser struct {
	ID        int      `gorm:"id" json:"id,omitempty"`
	Username  string   `json:"username,omitempty"`
	Password  string   `json:"password,omitempty"`
	Name      string   `json:"name"`
	Avatar    string   `json:"avatar"`
	IsLock    int      `json:"is_lock,omitempty" gorm:"default:1"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
	UpdatedAt JsonTime `json:"updated_at,omitempty"`
}

func (AdminUser) TableName() string {
	return "a_admin_user"
}
