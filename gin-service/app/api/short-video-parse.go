package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/servers/handleVideo"
	"github.com/putyy/ai-share/config"
	"strconv"
	"strings"
	"time"
)

const (
	// pc端
	pc_ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"
	//mac chrome
	mac_ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"
	// 移动端
	phone_ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
)

func ShortVideoParse(c *gin.Context) {
	path := ""
	type platform struct {
		Name     string `json:"name"`
		Platform string `json:"platform"`
	}

	type resStruct struct {
		PlatformInfo *platform `json:"platformInfo"`
		Path         string    `json:"path"`
	}

	platformInfo := &platform{}

	keyWords := c.PostForm("content")
	if keyWords == "" {
		ResponseError(c, "缺少关键词", "")
		return
	}

	rUrl, err := library.ParseGetUrl(keyWords)
	if err != nil {
		ResponseError(c, err.Error(), "")
		return
	}

	rUrlMd5 := fmt.Sprintf("%x", md5.New().Sum([]byte(rUrl)))
	rKey := library.BuildRdsKv("video_parse_cache").GetKey(rUrlMd5)
	cacheData, error1 := library.Redis().Get(c, rKey).Result()
	if error1 == nil && cacheData != "" {
		res := &resStruct{}
		_ = json.Unmarshal([]byte(cacheData), &res)
		ResponseSuccess(c, res)
		return
	}

	todayCacheKey := library.BuildRdsKv("video_parse_user").GetKey(time.Now().Format("20060102"))
	todayCacheField := "count_" + strconv.Itoa(GetLoginUid(c))
	if config.Vip[GetLoginUserVip(c)].VideoParse > 0 {
		// 有配置限制
		userTodayParseCount, err11 := library.Redis().HMGet(c, todayCacheKey, todayCacheField).Result()
		if err11 == nil && userTodayParseCount[0] != nil {
			if count, _ := strconv.Atoi(userTodayParseCount[0].(string)); count >= 3 {
				ResponseError(c, "抓取次数超过今日限制", "")
				return
			}
		}
	}

	if strings.Contains(rUrl, "douyin.com") { //抖音
		platformInfo.Name = "抖音"
		platformInfo.Platform = "douyin"
		path, err = handleVideo.DouYin(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "huoshan.com") { //火山
		platformInfo.Name = "火山"
		platformInfo.Platform = "huoshan"
		path, err = handleVideo.HuoShan(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "kuaishou.com") { //快手
		platformInfo.Name = "快手"
		platformInfo.Platform = "kuaishou"
		path, err = handleVideo.KuaiShou(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "weishi.qq.com") { //微视
		platformInfo.Name = "微视"
		platformInfo.Platform = "weishi"
		path, err = handleVideo.WeiShi(rUrl, phone_ua)
	} else {
		platformInfo.Name = ""
		platformInfo.Platform = ""
		ResponseError(c, "暂不支持该平台！！！", "")
		return
	}
	if err != nil {
		ResponseError(c, err.Error(), "")
		return
	}

	res := &resStruct{
		PlatformInfo: platformInfo,
		Path:         path,
	}

	jsonData, err1 := json.Marshal(res)
	if err1 == nil {
		library.Redis().Set(c, rKey, jsonData, 600*time.Second)
	}
	library.Redis().HIncrBy(c, todayCacheKey, todayCacheField, 1)
	ResponseSuccess(c, res)
}
