package util

import (
	"github.com/disgoorg/disgo/discord"
)

var (
	HiddenOpt = discord.ApplicationCommandOptionBool{
		Name:        "hidden",
		Description: "makes message with ephemeral flag",
	}
)

// Colors!
const (
	DARK     = 0x2B2D31
	BLURPLE  = 0x5865F2
	LOWBLUE  = 0x3C4270
	HIGHBLUE = 0x00A8FC
)
