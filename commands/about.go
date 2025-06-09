package commands

import (
	"errors"

	"github.com/dexslender/orb/commands/about"
	"github.com/dexslender/orb/orb"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var aboutCmd = discord.SlashCommandCreate{
	Name: "about",
	Description: "just to get info about things",
	Options: []discord.ApplicationCommandOption{
		about.SubBotStats, about.SubGuild,
	},
}

func runAbout(o *orb.Orb) handler.CommandHandler {
	return func(ctx *handler.CommandEvent) error {
		subc := ctx.Vars["sub"]
		switch subc {
			case about.SubBotStats.Name:
				return about.BotStatsRun(ctx, o, len(List))
			case about.SubGuild.Name:
				return about.GuildRun(ctx)
			default:
				return errors.New("Unknown 'about' subcommand: "+subc)
			}
	}
}
