package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/dexslender/orb/orb"
	"github.com/dexslender/orb/util"
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
	})
	config, err := util.LoadConfig("botconfig.yml", "ORB")
	if err != nil {
		logger.Fatal("config failed", "err", err)
	}
	if config.Bot.DebugLog {
		logger.SetLevel(log.DebugLevel)
	}

	bot := orb.New(logger, config)

	bot.Setup()

	bot.StartNLock()
}
