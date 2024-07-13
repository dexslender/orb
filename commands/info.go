package commands

import (
	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
)

type Info struct {
	base
	discord.SlashCommandCreate
}

func (i *Info) Init(util.InteractionRegister) {
	i.Name = "info"
	i.Description = "just returns bot's info and states"
	i.Options = []discord.ApplicationCommandOption{util.HiddenOpt}
}

func (i *Info) Run(ctx *util.CommandContext) error {
	hidden := ctx.SlashCommandInteractionData().Bool("hidden")
	err := ctx.DeferCreateMessage(hidden)
	if err != nil {
		return err
	}
	cu, err := ctx.Client().Rest().GetCurrentUser("")
	if err != nil {
		return err
	}
	info := discord.NewEmbedBuilder().
		SetAuthorName(cu.Tag()).
		SetAuthorIcon(*cu.AvatarURL()).
		SetFooterTextf("Version: %s", ctx.Orb.Version).
		Build()

	_, err = ctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		AddEmbeds(info).
		Build())
	return err
}
