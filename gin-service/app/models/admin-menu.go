package models

// 后台用户
type AdminMenu struct {
	ID        int      `gorm:"id" json:"id,omitempty"`
	Pid       int      `json:"pid"`
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	Avatar    string   `json:"avatar"`
	Sort      int      `json:"sort" gorm:"default:1"`
	Icon      string   `json:"icon"`
	IsShow    int      `json:"is_show" gorm:"default:1"`
	IsLock    int      `json:"is_lock" gorm:"default:1"`
	AdminUid  int      `json:"admin_uid,omitempty"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
	UpdatedAt JsonTime `json:"updated_at,omitempty"`
	DeletedAt JsonTime `json:"deleted_at,omitempty"`
}

func (AdminMenu) TableName() string {
	return "a_admin_menu"
}
