package test

import (
	"testing"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/json/v2"
)

func TestBuildMessage(t *testing.T) {
	actions := discord.NewActionRow(
		discord.NewSecondaryButton("Stats", "guild-stats"),
		discord.NewSecondaryButton("Features", "guild-features"),
		// discord.NewSecondaryButton("", "guild-level")
	)

	// msg := discord.NewMessageCreateBuilder().
	// 	AddComponents(discord.NewContainer(
	// 		discord.NewSection(discord.NewTextDisplayf(
	// 			"**%s**\n%s",
	// 			"wub wub", "nothing o.o",
	// 		)).WithAccessory(discord.NewThumbnail("Should be a url, but this is a test")),
	// 		discord.NewSmallSeparator(),
	// 		&actions,
	// 	)).
	// Build()

	// data, err := json.MarshalIndent(msg, "", "	")
	// if err != nil { t.Fatal(err) }
	// t.Logf("%s", data)

	for i := range actions.Components {
		if x, ok := actions.Components[i].(discord.ButtonComponent); ok {
			x.Disabled = true
			actions.Components[i] = x
		}
	}

	data, err := json.MarshalIndent(actions, "", "	")
	if err != nil { t.Fatal(err) }

	t.Logf("%s", data)
}
