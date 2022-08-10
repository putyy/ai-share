package form

type MenuAdminForm struct {
	Id         int    `form:"id"`
	Gid        int    `form:"gid" binding:"required"`
	Name       string `form:"name" binding:"required"`
	Icon       string `form:"icon"`
	UseVip     int    `form:"use_vip"`
	ClickType  int    `form:"click_type"`
	ClickFunc  string `form:"click_func"`
	Path       string `form:"path"`
	AppId      string `form:"app_id"`
	ExtraData  string `form:"extra_data"`
	EnvVersion string `form:"env_version"`
	ShortLink  string `form:"short_link"`
	Sort       int    `form:"sort"`
	IsLock     int    `form:"is_lock" binding:"required"`
}
