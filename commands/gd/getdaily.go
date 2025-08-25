package gd

import (
	"io"

	"github.com/dexslender/orb/util/gd"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
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

func GetDailyRun(ctx *handler.CommandEvent) error {
	//weekly := ctx.SlashCommandInteractionData().Bool("weekly")
	
	res, err := gd.GDClient.Request(
		gd.Daily,
		gd.DailyParams{
			Secret: gd.COMMON_KEY,
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

