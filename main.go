package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/dexslender/orb/commands"
	"github.com/dexslender/orb/orb"
	"github.com/dexslender/orb/util"
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
	f, diags := parser.ParseHCLFile("bot.config")
	if diags.HasErrors() { log.Fatal("bot.config HCL error: ", diags.Error()) }
	mdiags := gohcl.DecodeBody(f.Body, orb.HCLctx, &config)
	if mdiags.HasErrors() { log.Fatal("bot.config HCL error: ", mdiags.Error()) }
	// -----logger
	logger.SetLevel(log.Level(config.Bot.LogLevel))
	logger.Debug("config loaded")
	// -----bot
	bot := orb.New(version, logger, &config)
	bot.SetActivityManager(&util.Amanager{Logger: logger, Config: bot.Config})
	manager := &util.Imanager{Logger: logger, Config: bot.Config, Orb: bot}
	manager.AddCommandsFromPackage(commands.Commands)
	bot.SetCommandManager(manager)
	bot.Setup()
}
