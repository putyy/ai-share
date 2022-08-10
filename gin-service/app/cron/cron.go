package cron

import (
	"context"
	"github.com/putyy/ai-share/app/library"
	"github.com/robfig/cron/v3"
	"time"
)

var Services *cron.Cron

func init() {
	Services = cron.New(cron.WithSeconds())
	_, _ = Services.AddFunc("59 59 23 * * *", func() {
		// 创建明天的hash key (限制用户每日抓取)
		ctx := context.TODO()
		tomorrow := time.Now().AddDate(0, 0, 1).Format("20060102")
		cacheKey := library.BuildRdsKv("video_parse_user").GetKey(tomorrow)
		library.Redis().HMSet(ctx, cacheKey, "tomorrow_"+tomorrow, tomorrow).Result()
		library.Redis().Expire(ctx, cacheKey, 86400*time.Second)
	})
	Services.Start()
}
