package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/models"
	"strconv"
)

type Menu struct {
	models.Menu
	MenuGroup models.MenuGroup `json:"menu_group,omitempty" gorm:"ForeignKey:gid;AssociationForeignKey:id"`
}

func MenuList(c *gin.Context) {
	data := make(map[string]interface{})
	var resList []Menu
	model := models.Db()

	var total int64

	if c.Query("keywords") != "" {
		model = model.Where("Name LIKE ?", "%"+c.Query("keywords")+"%")
	}

	if c.Query("is_lock") != "" {
		model = model.Where("is_lock = ?", c.Query("is_lock"))
	}

	if c.Query("gid") != "" {
		model = model.Where("gid = ?", c.Query("gid"))
	}

	if c.Query("total") != "" {
		total, _ = strconv.ParseInt(c.Query("total"), 10, 64)
	}

	model.Scopes(models.PaginateScope(c)).Preload("MenuGroup").Order("id desc").Find(&resList)
	if total == 0 {
		model.Model(Menu{}).Count(&total)
	}

	data["total"] = total
	data["page"] = c.Query("page")
	data["page_size"] = c.Query("page_size")
	data["list"] = resList
	data["m_mark"] = "menu"
	api.ResponseSuccess(c, data)
}

func MenuEdit(c *gin.Context) {
	var formData form.MenuAdminForm
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}

	model := models.Menu{
		Name:       formData.Name,
		Gid:        formData.Gid,
		Icon:       formData.Icon,
		UseVip:     formData.UseVip,
		ClickType:  formData.ClickType,
		ClickFunc:  formData.ClickFunc,
		Path:       formData.Path,
		AppId:      formData.AppId,
		ExtraData:  formData.ExtraData,
		EnvVersion: formData.EnvVersion,
		ShortLink:  formData.ShortLink,
		Sort:       formData.Sort,
		IsLock:     formData.IsLock,
		AdminUid:   GetAdminUid(c),
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
