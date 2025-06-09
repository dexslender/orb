package gd

import (
	"io"

	"github.com/dexslender/orb/util"
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
	res, err := util.GDClient.Request(
		util.Users,
		util.UsersParams{
			Query: query,
			Secret: util.COMMON_KEY,
		},
	)
	if err != nil { return err }
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil { return err }
	if user := util.DecodeGDData[util.PartialUser](string(data)); user != nil {
		fuser := pp.Sprint(user) 
		return ctx.CreateMessage(discord.NewMessageCreateBuilder().
			SetContentf("```ansi\n%s```", fuser).
		Build())
	}
	return ctx.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent("```go\nNot found```").
	Build())
}
// xd
