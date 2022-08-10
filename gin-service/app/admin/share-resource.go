package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/models"
	"strconv"
)

type ShareResource struct {
	models.ShareResource
	ShareResourceType models.ShareResourceType `json:"share_resource_type,omitempty" gorm:"ForeignKey:tid;AssociationForeignKey:id"`
}

func ShareResourceList(c *gin.Context) {
	data := make(map[string]interface{})
	var resList []ShareResource
	model := models.Db()

	var total int64

	if c.Query("keywords") != "" {
		model = model.Where("content LIKE ?", "%"+c.Query("keywords")+"%")
	}

	if c.Query("is_lock") != "" {
		model = model.Where("is_lock = ?", c.Query("is_lock"))
	}

	if c.Query("tid") != "" {
		model = model.Where("tid = ?", c.Query("tid"))
	}

	if c.Query("total") != "" {
		total, _ = strconv.ParseInt(c.Query("total"), 10, 64)
	}

	model.Scopes(models.PaginateScope(c)).Preload("ShareResourceType").Order("id desc").Find(&resList)
	if total == 0 {
		model.Model(ShareResource{}).Count(&total)
	}

	data["total"] = total
	data["page"] = c.Query("page")
	data["page_size"] = c.Query("page_size")
	data["list"] = resList
	data["m_mark"] = "share-resource"
	api.ResponseSuccess(c, data)
}

func ShareResourceEdit(c *gin.Context) {
	var formData form.ShareResourceAdminForm
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}

	model := models.ShareResource{
		Tid:      formData.Tid,
		ImgUrl:   library.GetSaveUrl(formData.ImgUrl),
		Content:  formData.Content,
		Sort:     formData.Sort,
		IsLock:   formData.IsLock,
		AdminUid: GetAdminUid(c),
	}

	var err error
	if formData.Id > 0 {
		err = models.Db().Model(&model).Where("id = ?", formData.Id).Updates(model).Error
		model.ID = formData.Id
	} else {
		err = models.Db().Create(&model).Error
	}
	if err != nil {
		api.ResponseError(c, "创建失败", err.Error())
		return
	}

	api.ResponseSuccess(c, model)
}
