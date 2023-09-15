package scheduler

import (
	"log"

	"github.com/bitebait/curry/config"
	"github.com/jasonlvhit/gocron"
)

func Init(f func()) *gocron.Scheduler {
	s := gocron.NewScheduler()
	err := s.Every(uint64(config.GetConfig.Cache.MaxAge)).From(gocron.NextTick()).Hours().Do(f)
	if err != nil {
		log.Fatal("Failed to initialize scheduler:", err)
	}
	return s
}
