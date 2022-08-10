package models

type MenuGroup struct {
	ID        int      `json:"id" gorm:"primary_key" `
	Sort      int      `json:"sort"`
	Name      string   `json:"name"`
	IsLock    int      `json:"is_lock"`
	Menus     []Menu   `json:"menu,omitempty" gorm:"ForeignKey:gid;AssociationForeignKey:id"`
	AdminUid  int      `json:"admin_uid,omitempty"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
	UpdatedAt JsonTime `json:"updated_at,omitempty"`
}

func (MenuGroup) TableName() string {
	return "a_menu_group"
}
