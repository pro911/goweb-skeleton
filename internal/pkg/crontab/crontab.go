package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewCronTab() {
	//初始化定时任务引擎
	c := cron.New()

	id, err := c.AddFunc("* * * * *", HealthInspect)
	fmt.Println(time.Now(), id, err)

	//启动任务
	c.Start()

	// 监听退出信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	//等待退出
	<-sigCh

	c.Stop()
}
