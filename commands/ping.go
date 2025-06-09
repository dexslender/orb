package commands

import (
	"time"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var ping = discord.SlashCommandCreate{
	Name: "ping",
	Description: "just returns pong",
}

func runPing(cctx *handler.CommandEvent) error {
	s := time.Now()
	err := cctx.DeferCreateMessage(false)
	if err != nil {
		return err
	}
	rest := time.Since(s).Round(time.Millisecond)
	GW := cctx.Client().Gateway.
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

func runPingRefresh(cctx *handler.ComponentEvent) error {
	s := time.Now()
	err := cctx.DeferUpdateMessage()
	if err != nil {
		return err
	}
	rest := time.Since(s).Round(time.Millisecond)
	GW := cctx.Client().Gateway.Latency().Round(time.Millisecond)
	_, err = cctx.Client().Rest.UpdateMessage(
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
