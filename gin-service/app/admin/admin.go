package admin

import "github.com/gin-gonic/gin"

func GetAdminUid(c *gin.Context) int {
	uid, _ := c.Get("adminUid")
	return uid.(int)
}
