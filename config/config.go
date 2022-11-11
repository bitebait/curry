package config

import (
	"log"

	"github.com/kkyr/fig"
)

var GetConfig = getConfig()

type config struct {
	App struct {
		AsciiName string `fig:"ascii_name" default:"\n ██████╗██╗   ██╗██████╗ ██████╗ ██╗   ██╗\n██╔════╝██║   ██║██╔══██╗██╔══██╗╚██╗ ██╔╝\n██║     ██║   ██║██████╔╝██████╔╝ ╚████╔╝ \n██║     ██║   ██║██╔══██╗██╔══██╗  ╚██╔╝  \n╚██████╗╚██████╔╝██║  ██║██║  ██║   ██║   \n ╚═════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   "`
	}

	Api struct {
		Endpoint string `fig:"endpoint" default:"/api"`
	}

	DB struct {
		Name string `fig:"file_name" default:"database.db"`
	}

	Crawler struct {
		ClientTimeout         int `fig:"client_timeout" default:"60"`          // Time in seconds. Default value: 60 Seconds
		ResponseHeaderTimeout int `fig:"response_header_timeout" default:"60"` // Time in seconds. Default value: 60 Seconds

	}

	Cache struct {
		MaxAge int `fig:"max_age" default:"12"` // Time in hours. Default value: 12 Hours
	}

	Currency struct {
		Currency string `fig:"symbol" default:"BRL"`
	}
}

func getConfig() *config {
	var cfg config
	if err := fig.Load(&cfg, fig.File("config.yml")); err != nil {
		log.Panic(err)
	}
	return &cfg
}
