package api

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/servers"
	"strconv"
)

func ResourceType(c *gin.Context) {
	data := make(map[string]interface{})
	data["list"] = servers.ShareResourceTypeList()
	ResponseSuccess(c, data)
}

func ResourceList(c *gin.Context) {
	lastId, _ := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	tid, _ := strconv.Atoi(c.DefaultQuery("tid", "0"))
	data := make(map[string]interface{})
	data["list"] = servers.ShareResourceList(lastId, tid)
	ResponseSuccess(c, data)
}
