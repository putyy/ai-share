package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/library"
	"time"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		token := c.GetHeader("x-token")
		if token == "" {
			api.Response(c, library.ErrorAuthCheckTokenFail, data)
			c.Abort()
			return
		}

		claims, err := library.ParseAdminToken(token)
		if err != nil {
			api.Response(c, library.ErrorAuthCheckTokenFail, data)
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			api.Response(c, library.ErrorAuthCheckTokenFail, data)
			c.Abort()
			return
		}
		c.Set("adminUid", claims.Uid)
		c.Next()
	}
}
