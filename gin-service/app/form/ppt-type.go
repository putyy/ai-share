package form

type ContentAdminForm struct {
	Id      int    `form:"id"`
	Scene   int    `form:"scene" binding:"required"`
	Name    string `form:"name" binding:"required"`
	Content string `form:"content" binding:"required"`
}
