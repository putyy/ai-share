package models

// 朋友圈素材分类
type ShareResourceType struct {
	ID        int      `gorm:"primary_key" json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Sort      int      `json:"sort"`
	IsLock    int      `json:"is_lock,omitempty" gorm:"default:1"`
	AdminUid  int      `json:"admin_uid,omitempty"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
	UpdatedAt JsonTime `json:"updated_at,omitempty"`
}

func (ShareResourceType) TableName() string {
	return "a_share_resource_type"
}
