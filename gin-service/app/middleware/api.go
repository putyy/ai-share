package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/config"
	"strconv"
	"time"
)

func Api() gin.HandlerFunc {
	return func(c *gin.Context) {
		systemInfoKey := library.BuildRdsKv("system_info").GetHashKey()
		systemInfo, err1 := library.Redis().HGetAll(c, systemInfoKey).Result()
		if err1 == nil && len(systemInfo) > 0 {
			// 小程序审核中
			if systemInfo["mini_check"] == "2" {
				api.Response(c, library.SystemWxMiniIsCheck, "")
				c.Abort()
				return
			}

			if systemInfo["system_close"] == "2" {
				api.Response(c, library.SystemIsClose, map[string]string{"content": systemInfo["system_close_content"]})
				c.Abort()
				return
			}
		}

		token := c.GetHeader("x-token")
		if token == "" {
			api.Response(c, library.ErrorAuthCheckTokenFail, "")
			c.Abort()
			return
		}

		claims, err := library.ParseApiToken(token)
		if err != nil {
			api.Response(c, library.ErrorAuthCheckTokenFail, "")
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			api.Response(c, library.ErrorAuthCheckTokenFail, "")
			c.Abort()
			return
		}

		if claims.Version != config.App.Version {
			api.Response(c, library.SystemVersionUpdate, "")
			c.Abort()
			return
		}

		c.Set("ApiClaims", claims)

		// todo 是否被锁定删除
		userCacheKey := library.BuildRdsKv("user_info").GetHashKey(strconv.Itoa(claims.Uid))
		userCacheData, err1 := library.Redis().HGetAll(c, userCacheKey).Result()
		if err1 == nil && len(userCacheData) > 0 {
			if userCacheData["deleted_at"] != "0" {
				api.Response(c, library.UserIsDel, map[string]string{"content": library.MsgFlags[library.UserIsDel]})
				c.Abort()
				return
			}
			if userCacheData["is_lock"] != "1" {
				api.Response(c, library.UserIsLock, map[string]string{"content": library.MsgFlags[library.UserIsLock]})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
