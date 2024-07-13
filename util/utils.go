package util

import "github.com/disgoorg/disgo/discord"

var (
	HiddenOpt = discord.ApplicationCommandOptionBool{
		Name:        "hidden",
		Description: "makes message with ephemeral flag",
	}
)
