package main

import (
	"fmt"

	"kevinsudut.com/paper-id-take-home-test/app"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/config"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/database"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/log"
)

func main() {
	cfg, err := config.ReadConfig("./config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Failed read config: %s", err.Error()))
	}

	logger, err := log.InitLog(cfg.Log)
	if err != nil {
		panic(fmt.Sprintf("Failed init logger: %s", err.Error()))
	}

	db, err := database.InitDatabase()
	if err != nil {
		logger.FatalTrace(err, "Failed init database")
	}

	err = app.InitApp(cfg.Http, logger, db)
	if err != nil {
		logger.FatalTrace(err, "Failed init app")
	}
}
