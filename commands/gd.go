package commands

import (
	"github.com/dexslender/orb/commands/gd"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var gdCmd = discord.SlashCommandCreate{
	Name: "gd",
	Description: "Utilities to interact with Geometry Dash servers.",
	Options: []discord.ApplicationCommandOption{
		gd.GetUserCommand,
		gd.GetDailyCommand,
	},
}

func runGD(ctx *handler.CommandEvent) error {
	subcmd := ctx.SlashCommandInteractionData().SubCommandName
	switch *subcmd {
	case "get-user":
		return gd.GetUserRun(ctx)
	case "get-daily":
		return gd.GetDailyRun(ctx)
	default:
		return nil
	}
}
