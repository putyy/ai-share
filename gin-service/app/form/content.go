package form

type PptTypeAdminForm struct {
	Id     int    `form:"id"`
	Name   string `form:"name" binding:"required"`
	Sort   int    `form:"sort"`
	IsLock int    `form:"is_lock" binding:"required"`
}
