package rank

import "github.com/robfig/cron"

func InitCron() {
	c := cron.New()
	_ = c.AddFunc("@hourly", initiateUpdateScore)
	c.Start()
}

func initiateUpdateScore() {
	
}
