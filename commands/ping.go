package commands

import (
	"time"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
)

type Ping struct { base; discord.SlashCommandCreate }

func (c *Ping) Init(add util.InteractionRegister) {
	c.Name = "ping"
	c.Description = "just returns pong"
	// TODO: Make hidden by default and fix problems (ephemeral update message only once time)
	// c.Options = []discord.ApplicationCommandOption{ util.HiddenOpt }

	add.Component("refresh-ping", refresh)
}

func (c *Ping) Run(cctx *util.CommandContext) error {
	s := time.Now()
	err := cctx.DeferCreateMessage(false)
	if err != nil {
		return err
	}
	rest := time.Since(s).Round(time.Millisecond)
	GW := cctx.Client().Gateway().
		Latency().Round(time.Millisecond)

	refresh := discord.NewSecondaryButton("Refresh", "refresh-ping")

	_, err = cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
	SetContentf("```yaml\n%s```", util.FormatSpacing(
			"Rest", rest,
			"Gateway", GW,
		)).
		AddActionRow(refresh).
		Build(),
	)
	return err
}

func refresh(cctx *util.ComponentContext) error {
	s := time.Now()
	err := cctx.DeferUpdateMessage()
	if err != nil {
		return err
	}
	rest := time.Since(s).Round(time.Millisecond)
	GW := cctx.Client().Gateway().Latency().Round(time.Millisecond)
	_, err = cctx.Client().Rest().UpdateMessage(
		cctx.Message.ChannelID,
		cctx.Message.ID,
		discord.NewMessageUpdateBuilder().
			SetContentf("```yaml\n%s```", util.FormatSpacing(
				"Rest", rest,
				"Gateway", GW,
			)).
			Build(),
	)
	return err
}
