package commands

import (
	"errors"

	"github.com/dexslender/orb/commands/setup"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/omit"
)

var setupCmd = discord.SlashCommandCreate{
	Name: "setup",
	Description: "just setup/configure bot features",
	DefaultMemberPermissions: omit.NewPtr(discord.PermissionAdministrator),
	Options: []discord.ApplicationCommandOption{setup.TicketsCommand},
}

func runSetup(ctx *handler.CommandEvent) error {
	switch ctx.Vars["sub"] {
	case "tickets":
		return setup.RunTickets(ctx)
	default:
		return errors.New("unknown subcommand: " + ctx.Vars["sub"])
	}
}
