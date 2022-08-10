package models

type Menu struct {
	ID         int      `json:"id" gorm:"primary_key"`
	Gid        int      `json:"gid"`
	Name       string   `json:"name"`
	Icon       string   `json:"icon"`
	UseVip     int      `json:"use_vip"`
	ClickType  int      `json:"click_type"`
	ClickFunc  string   `json:"click_func"`
	Path       string   `json:"path"`
	AppId      string   `json:"app_id"`
	ExtraData  string   `json:"extra_data"`
	EnvVersion string   `json:"env_version"`
	ShortLink  string   `json:"short_link"`
	Sort       int      `json:"sort"`
	IsLock     int      `json:"is_lock,omitempty" gorm:"default:1"`
	AdminUid   int      `json:"admin_uid,omitempty"`
	CreatedAt  JsonTime `json:"created_at,omitempty"`
	UpdatedAt  JsonTime `json:"updated_at,omitempty"`
}

func (Menu) TableName() string {
	return "a_menu"
}
