package crons

import (
	"log"

	"github.com/hhxsv5/gin-x-skeleton/app/cron/jobs"
	"github.com/hhxsv5/gin-x/cron"
)

var (
	m *cron.Manager
)

func Start() {
	m = cron.NewManager()
	m.Register(jobs.Test{})
	m.Start()
	log.Println("cron has started successfully")
}
