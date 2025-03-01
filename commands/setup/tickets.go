package setup

import (
	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/json"
)

var TicketsCommand = discord.ApplicationCommandOptionSubCommand{
	Name:        "tickets",
	Description: "just setup/configure ticket system, by default with 'commands'",
	// Options: []discord.ApplicationCommandOption{
	// 	discord.ApplicationCommandOptionInt{
	// 		Name: "type",
	// 		Description: "Type of ticket system",
	// 		Choices: []discord.ApplicationCommandOptionChoiceInt{
	// 			{
	// 				Name: "Channel",
	// 				Value: 2,
	// 			},
	// 			{
	// 				Name: "Command",
	// 				Value: 1,
	// 			},
	// 		},
	// 	},
	// },
}

var ticketChannelDefaultConfig discord.GuildChannelCreate = discord.GuildTextChannelCreate{
	Name: "tickets",
	Topic: "Create a new ticket. Contact server mods and support team from here!",
	// PermissionOverwrites: [], // TODO: needed to prevent messages from unauthorized users
}

func RunTickets(cctx *util.CommandContext) error {
	stickets := discord.NewEmbedBuilder().
		SetTitle("Ticket System Setup Wizard").
		SetDescription("Select type of ticket system\nUsers wishing to open a ticket must:").
		AddFields(
			discord.EmbedField{
				Name: "Press a `Button`",
				Value: "This option makes the bot send a message with a button to a designated channel, where members can click it to create a ticket.",
				Inline: json.Ptr(true),
			},
			discord.EmbedField{
				Name: "Execute a `Command`",
				Value: "Users should use the `/ticket` command (provided by the bot) to create a ticket.",
				Inline: json.Ptr(true),
			},
		).
		Build();
	actions := discord.NewActionRow().
		AddComponents(
			discord.NewSecondaryButton("Button", "1"),
			discord.NewSecondaryButton("Command", "2"),
		)

	err := cctx.CreateMessage(discord.NewMessageCreateBuilder().
	AddEmbeds(stickets).
	SetEphemeral(true).
	AddContainerComponents(actions).
	Build())
	if err != nil { return err }
	return nil
}
