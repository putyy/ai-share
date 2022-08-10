package form

type ShareResourceAdminForm struct {
	Id      int    `form:"id"`
	Tid     int    `form:"tid" binding:"required"`
	ImgUrl  string `form:"img_url" binding:"required"`
	Content string `form:"content" binding:"required"`
	Sort    int    `form:"sort"`
	IsLock  int    `form:"is_lock" binding:"required"`
}
