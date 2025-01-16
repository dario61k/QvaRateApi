package cron

import (
	"os"
	"qvarate_api/cron/jobs"

	"github.com/robfig/cron/v3"
)

func SetupCron() {
	c := cron.New()

	c.AddFunc(os.Getenv("CRON_EXECUTION"), jobs.GetData)

	c.Start()
}
