package commands

import (
	"github.com/dexslender/orb/commands/gd"
	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
)

type GD struct {
	base
	discord.SlashCommandCreate
}

func (c *GD) Init(util.InteractionRegister) {
	c.Name = "gd"
	c.Description = "Utilities to interact with Geometry Dash servers."
	c.Options = []discord.ApplicationCommandOption{gd.GetUserCommand}
}

func (c *GD) Run(ctx *util.CommandContext) error {
	subcmd := ctx.SlashCommandInteractionData().SubCommandName
	switch *subcmd {
	case "get-user": return gd.GetUserRun(ctx)
	default: return nil
	}
}
