package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/dexslender/orb/commands"
	"github.com/dexslender/orb/orb"
	"github.com/dexslender/orb/util"
	"github.com/kkyr/fig"
)

var version string = "dev"

func main() {
	// -----logger
	logger := log.NewWithOptions(os.Stderr, log.Options{ReportTimestamp: true})
	// -----config
	var config orb.Config
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
	bot := orb.New(version, logger, &config)
	bot.SetActivityManager(&util.Amanager{Logger: logger, Config: bot.Config})
	manager := &util.Imanager{Logger: logger, Config: bot.Config, Orb: bot}
	manager.AddCommands(commands.Commands...)
	bot.SetCommandManager(manager)
	bot.Setup()
}
