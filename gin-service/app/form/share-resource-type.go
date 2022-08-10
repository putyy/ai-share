package form

type ShareResourceTypeAdminForm struct {
	Id     int    `form:"id"`
	Name   string `form:"name" binding:"required"`
	Sort   int    `form:"sort" binding:"required"`
	IsLock int    `form:"is_lock" binding:"required"`
}
