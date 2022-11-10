package scheduler

import (
	"log"

	"github.com/bitebait/curry/config"
	"github.com/jasonlvhit/gocron"
)

func Init(f func()) *gocron.Scheduler {
	cfg := config.GetConfig()
	s := gocron.NewScheduler()
	err := s.Every(uint64(cfg.Cache.MaxAge)).From(gocron.NextTick()).Hours().Do(f)
	if err != nil {
		log.Panic("Failed to init scheduler.")
	}
	return s
}
