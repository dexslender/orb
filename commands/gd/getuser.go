package gd

import (
	"io"

	"github.com/dexslender/orb/util/gd"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/k0kubun/pp/v3"
)

var GetUserCommand = discord.ApplicationCommandOptionSubCommand{
	Name: "get-user",
	Description: "Get info about some gd user.",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name: "query",
			Description: "Your query request.",
			Required: true,
		},
	},
}

func GetUserRun(ctx* handler.CommandEvent) error {
	query := ctx.SlashCommandInteractionData().String("query")
	res, err := gd.GDClient.Request(
		gd.Users,
		gd.UsersParams{
			Query: query,
			Secret: gd.COMMON_KEY,
		},
	)
	if err != nil { return err }
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil { return err }
	if user := gd.DecodeGDData[gd.PartialUser](string(data)); user != nil {
		fuser := pp.Sprint(user) 
		return ctx.CreateMessage(discord.NewMessageCreateBuilder().
			SetIsComponentsV2(true).
			AddComponents(discord.NewTextDisplayf("```ansi\n%s```", fuser)).
		Build())
	} else {
		return ctx.CreateMessage(discord.NewMessageCreateBuilder().
			SetIsComponentsV2(true).
			AddComponents(discord.NewContainer(
				discord.NewTextDisplay(":no_entry_sign: _Not Found_"),
			)).
		Build())
	}
}
