package gd

import (
	"io"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
)

var GetDailyCommand = discord.ApplicationCommandOptionSubCommand{
	Name: "get-daily",
	Description: "Get daily or weekly, by default request current daily",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionBool{
			Name: "weekly",
			Description: "Get weekly?",
			Required: false,
		},
	},
}

func GetDailyRun(ctx *util.CommandContext) error {
	//weekly := ctx.SlashCommandInteractionData().Bool("weekly")
	
	res, err := util.GDClient.Request(
		util.Daily,
		util.DailyParams{
			Secret: util.COMMON_KEY,
			Weekly: 1,
		},
	)
	if err != nil { return err }
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil { return err }
	return ctx.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent(string(data)).	
	Build())
}

