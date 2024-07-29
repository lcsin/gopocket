package task

import (
	"log"
	"testing"
	"time"
)

func TestSimpleTask(t *testing.T) {
	SimpleTask(time.Second, func() {
		log.Println("task exec ...")
	})

	time.Sleep(time.Second * 10)
}

func TestCornTask(t *testing.T) {
	CronTask("* * * * * ?", func() {
		log.Println("task exec ...")
	})

	time.Sleep(time.Second * 10)
}
