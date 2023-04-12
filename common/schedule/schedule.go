package schedule

import (
	rank_schedule "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/schedule/rank"
	"github.com/robfig/cron"
)

func InitCron() {
	c := cron.New()
	_ = c.AddFunc("@hourly", rank_schedule.InitiateUpdateScore)
	c.Start()
}
