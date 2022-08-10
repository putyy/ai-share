package form

type PptCreateAdminForm struct {
	Id          int    `form:"id"`
	Name        string `form:"name" binding:"required"`
	Tid         int    `form:"tid" binding:"required"`
	Sort        int    `form:"sort" binding:"required"`
	AiDou       int    `form:"ai_dou" binding:"required"`
	DescContent string `form:"desc_content" binding:"required"`
	ImgUrl      string `form:"img_url" binding:"required"`
	FileUrl     string `form:"file_url" binding:"required"`
}

type PptListApiForm struct {
	CommonForm
	TypeId int `form:"type_id"`
}

type PptContentCreateAdminForm struct {
	Id       int      `form:"id"`
	ImageArr []string `form:"images[]"`
	Contents []string `form:"contents[]"`
}
