package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/models"
	"gorm.io/gorm"
	"strconv"
)

type Ppt struct {
	models.Ppt
	PptType models.PptType `json:"ppt_type,omitempty" gorm:"ForeignKey:tid;AssociationForeignKey:id"`
}

func PptList(c *gin.Context) {
	data := make(map[string]interface{})
	var resList []Ppt
	model := models.Db()

	var total int64

	if c.Query("keywords") != "" {
		model = model.Where("desc_content LIKE ?", "%"+c.Query("keywords")+"%")
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

	model.Scopes(models.PaginateScope(c)).Preload("PptType").Order("id desc").Find(&resList)
	if total == 0 {
		model.Model(Ppt{}).Count(&total)
	}

	data["total"] = total
	data["page"] = c.Query("page")
	data["page_size"] = c.Query("page_size")
	data["list"] = resList
	data["m_mark"] = "ppt"
	api.ResponseSuccess(c, data)
}

func PptEdit(c *gin.Context) {
	var formData form.PptCreateAdminForm
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}

	model := models.Ppt{
		Tid:         formData.Tid,
		Name:        formData.Name,
		ImgUrl:      library.GetSaveUrl(formData.ImgUrl),
		DescContent: formData.DescContent,
		Sort:        formData.Sort,
		FileUrl:     library.GetSaveUrl(formData.FileUrl),
		AiDou:       formData.AiDou,
		AdminUid:    GetAdminUid(c),
	}

	err2 := models.Db().Transaction(func(tx *gorm.DB) error {
		if formData.Id > 0 {
			err11 := tx.Model(&model).Where("id = ?", formData.Id).Updates(model).Error
			model.ID = formData.Id
			if err11 != nil {
				// 返回任何错误都会回滚事务
				return err11
			}
		} else {
			err2 := tx.Create(&model).Error
			if err2 != nil {
				// 返回任何错误都会回滚事务
				return err2
			}
		}
		return nil
	})

	if err2 != nil {
		api.ResponseError(c, "创建失败", err2.Error())
		return
	}

	api.ResponseSuccess(c, model)
}

func PptContent(c *gin.Context) {
	var ppContent []models.PptContent
	models.Db().Where("pid = ?", c.Query("id")).Find(&ppContent)
	api.ResponseSuccess(c, map[string]interface{}{
		"content": ppContent,
	})
}

func PptContentEdit(c *gin.Context) {
	var formData form.PptContentCreateAdminForm
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}

	var pptContent []models.PptContent
	for i, v := range formData.ImageArr {
		pptContent = append(pptContent, models.PptContent{
			Pid:     formData.Id,
			ImgUrl:  library.GetSaveUrl(v),
			Content: formData.Contents[i],
		})
	}

	err2 := models.Db().Transaction(func(tx *gorm.DB) error {
		err1 := tx.Where("pid = ?", formData.Id).Delete(models.PptContent{}).Error
		if err1 != nil {
			return err1
		}

		err2 := tx.Create(&pptContent).Error
		if err2 != nil {
			return err2
		}
		return nil
	})

	if err2 != nil {
		api.ResponseError(c, "操作失败", err2.Error())
		return
	}

	api.ResponseSuccess(c, "")
}
