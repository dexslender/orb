package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/dexslender/orb/commands"
	"github.com/dexslender/orb/orb"
	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/snowflake/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

var version string = "dev"

func main() {
	// -----logger
	logger := log.NewWithOptions(os.Stderr, log.Options{ReportTimestamp: true})
	// -----config
	var config orb.Config
	parser := hclparse.NewParser()
	f, diags := parser.ParseHCLFile("bot.hcl")
	if diags.HasErrors() {
		log.Fatal("bot.config HCL parser error: ", diags.Error())
	}
	mdiags := gohcl.DecodeBody(f.Body, orb.HCLctx(logger), &config)
	if mdiags.HasErrors() {
		log.Fatal("bot.config HCL decoder error: ", mdiags.Error())
	}
	// -----logger
	logger.SetLevel(log.Level(config.Bot.LogLevel))
	logger.Debug("config loaded")
	// -----bot
	bot := orb.New(version, logger, &config)
	bot.SetActivityManager(&util.Amanager{Logger: logger, Config: bot.Config})
	router := handler.New()
	commands.Setup(bot, router)
	if config.Bot.SetupCommands {
		handler.SyncCommands(bot.Client, commands.List, []snowflake.ID{})
	}
	bot.Setup()
	bot.AddEventListeners(router)

	// Connect to discord
	bot.StartNLock()
}
