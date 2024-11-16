package scheduler

import (
	"log"

	"github.com/bitebait/curry/config"
	"github.com/jasonlvhit/gocron"
)

func Init(jobFunction func()) *gocron.Scheduler {
	scheduler := gocron.NewScheduler()
	maxAgeHours := uint64(config.GetConfig.Cache.MaxAge)

	err := scheduler.Every(maxAgeHours).From(gocron.NextTick()).Hours().Do(jobFunction)
	if err != nil {
		log.Fatal("Failed to initialize scheduler:", err)
	}
	return scheduler
}
