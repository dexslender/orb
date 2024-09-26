package gd

import (
	"io"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
)

var GetUserCommand = discord.ApplicationCommandOptionSubCommand{
	Name: "get-user",
	Description: "Get info about some gd user.",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name: "query",
			Description: "Your query request.",
		},
	},
}

func GetUserRun(ctx* util.CommandContext) error {
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
	return ctx.CreateMessage(discord.NewMessageCreateBuilder().
		SetContentf("```go\n%s```", string(data)).
	Build())
}
