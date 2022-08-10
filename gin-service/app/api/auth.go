package api

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/servers"
	config2 "github.com/putyy/ai-share/config"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"log"
	"strconv"
)

type auth struct {
	Username string `valid:"Required; MaxSize(18); MinSize(8);"`
	Password string `valid:"Required; MaxSize(18); MinSize(6);"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{username, password}
	ok, _ := valid.Valid(a)

	data := make(map[string]interface{})
	if !ok {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.Message: %s", err.Key, err.Message)
		}
		ResponseError(c, "缺少参数", "")
		return
	}

	maps := make(map[string]interface{})
	maps["user_name"] = username
	user := servers.UserInfo(maps, "*")

	if user.ID <= 0 {
		ResponseError(c, "账号不存在", "")
		return
	}

	if user.IsLock == 2 {
		Response(c, library.UserIsLock, "")
		return
	}

	if user.DeletedAt.Valid == true {
		Response(c, library.UserIsDel, "")
		return
	}

	if fmt.Sprintf("%x", md5.Sum([]byte(password))) != user.Password {
		ResponseError(c, "密码错误", "")
		return
	}

	token, err := library.GenerateApiToken(library.ApiClaims{
		Uid:  user.ID,
		Type: 1,
		Vip:  user.Vip,
	})

	if err != nil {
		ResponseError(c, "token生成失败", "")
		return
	}
	data["uid"] = user.ID
	data["token"] = token
	ResponseSuccess(c, data)
	return
}

// 微信小程序登录
func WxMiNiLogin(c *gin.Context) {
	redisOpts := &cache.RedisOpts{
		Host:        config2.Redis.Host,
		Database:    config2.Redis.Database,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60, //second
	}

	redisCache := cache.NewRedis(c, redisOpts)

	cfg := &config.Config{
		AppID:     config2.Wechat.AppID,
		AppSecret: config2.Wechat.AppSecret,
	}

	wc := wechat.NewWechat()
	wc.SetCache(redisCache)
	mini := wc.GetMiniProgram(cfg)
	wxUser, err := mini.GetAuth().Code2Session(c.Query("code"))
	if err != nil || wxUser.ErrCode != 0 {
		ResponseError(c, "登录失败", "")
		return
	}

	maps := make(map[string]interface{})
	maps["unionid"] = wxUser.UnionID
	user := servers.UserInfo(maps, "*")
	if user.ID <= 0 {
		fromUid, _ := strconv.Atoi(c.DefaultQuery("from_uid", "1"))
		uid, err := servers.UserRegister(wxUser, fromUid)
		if err != nil {
			ResponseError(c, "注册失败", "")
			return
		}
		user.ID = uid
	}

	if user.IsLock == 2 {
		Response(c, library.UserIsLock, "")
		return
	}

	if user.DeletedAt.Valid == true {
		Response(c, library.UserIsDel, "")
		return
	}

	token, err := library.GenerateApiToken(library.ApiClaims{
		Uid:  user.ID,
		Type: 1,
		Vip:  user.Vip,
	})
	if err != nil {
		ResponseError(c, "token生成失败", "")
		return
	}

	data := make(map[string]interface{})
	data["uid"] = user.ID
	data["token"] = token
	ResponseSuccess(c, data)
	return
}
