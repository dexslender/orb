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

func runPing(e *handler.CommandEvent) error {
	a := time.Now()
	err := e.DeferCreateMessage(false)
	if err != nil { return err }
	rest := time.Since(a).Round(time.Millisecond)
	gate := e.Client().Gateway.Latency().Round(time.Millisecond)

	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		SetComponents(buildContainer(rest, gate)).
		AddFlags(discord.MessageFlagIsComponentsV2).
	Build())
	return err
}

func runPingRefresh(data discord.ButtonInteractionData, e *handler.ComponentEvent) error {
	a := time.Now()
	err := e.DeferUpdateMessage()
	if err != nil { return err }
	rest := time.Since(a).Round(time.Millisecond)
	gate := e.Client().Gateway.Latency().Round(time.Millisecond)
	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		SetComponents(buildContainer(rest, gate)).
	Build())
	return err
}

func buildContainer(rest, gate time.Duration) discord.ContainerComponent {
	return discord.NewContainer(
		discord.NewTextDisplayf("```yaml\n%s```", util.FormatSpacing(
			"Rest", rest,
			"Gateway", gate,
		)),
		discord.NewActionRow(
			discord.NewSecondaryButton("Refresh", "/ping/refresh"),
		),
	)
}
