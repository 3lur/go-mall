package config

import (
	"flag"
	"fmt"

	"github.com/3lur/go-mall/internal/common/configs"
	"github.com/joho/godotenv"
)

type Config struct {
	Database configs.Database
	App      configs.App
}

func New() *Config {
	var filename string

	// register command, usage: --config=.env
	flag.StringVar(&filename, "config", ".env", "config file, eg: --config=[.filename]")
	flag.Parse()

	if filename == "" {
		filename = ".env"
	}
	if err := godotenv.Load(filename); err != nil {
		panic(fmt.Errorf("failed to load [%s] file: %s", filename, err))
	}

	return &Config{
		Database: configs.DatabaseStore(),
		App:      configs.AppStore(),
	}
}
