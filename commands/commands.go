package commands

import (
	"errors"

	"github.com/dexslender/orb/util"
)

var Commands = []util.Command{
	new(Ping),
	new(Purge),
	new(Setup),
	new(Info),
	new(GD),
}

type base struct {
	util.Command
}

func (c *base) Run(*util.CommandContext) error { return errors.New("missing run function :(") }

func (c *base) Error(cctx *util.CommandContext, err error) {
	command := cctx.Data.CommandName()
	if sub := cctx.SlashCommandInteractionData().SubCommandName; sub != nil {
		command+="/"+*sub
	}
	cctx.Logger.Error("command returned error",
		"command", command,
		"error", err,
	)
}
