package commands

import (
	"errors"

	"github.com/dexslender/orb/commands/about"
	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
)

type About struct { base; discord.SlashCommandCreate }

func (c *About) Init(util.InteractionRegister) {
	c.Name = "about"
	c.Description = "just to get info about things"

	c.Options = []discord.ApplicationCommandOption{
		about.SubBotStats, about.SubGuild,
	}
}

func (c *About) Run(ctx *util.CommandContext) error {
	subc := ctx.SlashCommandInteractionData().SubCommandName
	switch *subc {
		case about.SubBotStats.Name:
			return about.BotStatsRun(ctx, len(Commands))
		case about.SubGuild.Name:
			return about.GuildRun(ctx)
		default:
			return errors.New("Unknown 'about' subcommand: "+*subc)
		}
}
