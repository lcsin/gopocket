package task

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// SimpleTask 简单的定时任务，不建议在生产环境中使用
func SimpleTask(duration time.Duration, callback func()) {
	go func() {
		t := time.NewTicker(duration)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				callback()
			}
		}
	}()
}

// CronTask Cron定时任务
func CronTask(spec string, callback func()) {
	c := cron.New(cron.WithSeconds())

	if _, err := c.AddFunc(spec, callback); err != nil {
		log.Printf("add corn task error: %v", err)
		return
	}

	c.Start()
}
