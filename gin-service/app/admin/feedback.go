package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/models"
	"strconv"
)

func FeedbackList(c *gin.Context) {
	data := make(map[string]interface{})
	var resList []models.Feedback
	model := models.Db()

	var total int64

	if c.Query("keywords") != "" {
		model = model.Where("content LIKE ?", "%"+c.Query("keywords")+"%")
	}

	if c.Query("total") != "" {
		total, _ = strconv.ParseInt(c.Query("total"), 10, 64)
	}

	model.Scopes(models.PaginateScope(c)).Order("id desc").Find(&resList)
	if total == 0 {
		model.Model(models.Feedback{}).Count(&total)
	}

	data["total"] = total
	data["page"] = c.Query("page")
	data["page_size"] = c.Query("page_size")
	data["list"] = resList
	api.ResponseSuccess(c, data)
}

func FeedbackRemark(c *gin.Context) {
	id := c.PostForm("id")
	remark := c.PostForm("remark")
	if id == "" {
		api.ResponseError(c, "参数错误", "")
		return
	}
	m := models.Feedback{
		Remark:   remark,
		AdminUid: GetAdminUid(c),
	}
	err := models.Db().Model(&m).Where("id = ?", id).Updates(m).Error

	if err != nil {
		api.ResponseError(c, "创建失败", err.Error())
		return
	}

	api.ResponseSuccess(c, "")
}
