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
			Required: true,
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
	if user, err := util.DecodeGDUserData(data); err != nil {
		return err
	} else {
		return ctx.CreateMessage(discord.NewMessageCreateBuilder().
			AddEmbeds(discord.NewEmbedBuilder().
				SetFooterText("Testing").
				SetDescriptionf(
					`Name: %s
ID: %s
Stars: %s
SecretCoins: %s
Moons: %s
`,
				user[1], user[2], user[3], user[13], user[52]).
			Build()).
		Build())
	}
}
