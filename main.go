package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/dexslender/orb/commands"
	"github.com/dexslender/orb/orb"
	"github.com/dexslender/orb/util"
	"github.com/kkyr/fig"
)

func main() {
	// -----logger
	logger := log.NewWithOptions(os.Stderr, log.Options{ReportTimestamp: true})
	// -----config
	var config util.Config
	err := fig.Load(&config,
		fig.File("botconfig.yml"),
		fig.UseEnv("ORB"),
	)
	if err != nil {
		logger.Fatal("when loading config", "err", err)
	}

	logger.SetLevel(log.Level(config.Bot.LogLevel))
	logger.Debug("config loaded")

	// -----bot
	bot := orb.New(logger, &config)
	bot.SetActivityManager(&util.Amanager{Logger: logger, Config: bot.Config})
	bot.SetCommandManager(&util.Imanager{Logger: logger, Config: bot.Config})
	bot.AddCommands(commands.Commands...)
	bot.Setup()
}
