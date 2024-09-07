package commands

import (
	"errors"

	"github.com/dexslender/orb/commands/setup"
	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/json"
)

type Setup struct {
	base
	discord.SlashCommandCreate
}

func (c *Setup) Init(add util.InteractionRegister) {
	c.Name = "setup"
	c.Description = "just setup/configure bot features"
	c.DefaultMemberPermissions = json.NewNullablePtr(discord.PermissionAdministrator)
	c.Options = []discord.ApplicationCommandOption{setup.TicketsCommand}

	add.Component("ticket-open", setup.OnClickTicket)
}

func (c *Setup) Run(cctx *util.CommandContext) error {
	switch *cctx.SlashCommandInteractionData().SubCommandName {
	case "tickets":
		return setup.RunTickets(cctx)
	default:
		return errors.New("unknown subcommand: " + *cctx.SlashCommandInteractionData().SubCommandName)
	}
}
