package main

import (
	"fmt"
	"github.com/putyy/ai-share/app/cron"
	"github.com/putyy/ai-share/config"
	"github.com/putyy/ai-share/router"
	"net/http"
	"time"
)

func main() {
	defer cron.Services.Stop()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.App.HTTPPort),
		Handler:        router.InitRouter(),
		ReadTimeout:    time.Duration(config.App.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.App.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}