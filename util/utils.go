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

func Btoi(b bool) int {
	if b { return 1 }
	return 0
}

// Colors!
const (
	DARK     = 0x2B2D31
	BLURPLE  = 0x5865F2
	LOWBLUE  = 0x3C4270
	HIGHBLUE = 0x00A8FC
)
