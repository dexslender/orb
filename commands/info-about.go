package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/json"
)

type about struct {
	Name string
	Description string
	Details []discord.EmbedField
}

var abouts = map[string]about{
	"tickets": {
		Name: "Tickets", 
		Description: `Ticket system setup, the command provides this options for setup system`,
		Details: []discord.EmbedField{
			{
				Name: "🔍 channel",
				Value: `Custom channel to assign ticket system, it creates/send a message with active button.`,
				Inline: json.Ptr(true),
			},
			{
				Name: "🤖 webhook `unavailable`",
				Value: "Set a custom webhook, then bot clones it and send a message with button.",
				Inline: json.Ptr(true),
			},
		},
	},
}
