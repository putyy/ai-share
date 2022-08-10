package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/config"
)

func VipConfig(c *gin.Context) {
	api.ResponseSuccess(c, config.Vip)
}
