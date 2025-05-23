package util

import (
	"fmt"
	"strings"

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

func FormatSpacing(keyvals ...any) (s string) {
	max := 0
	for i, v := range keyvals {
		if i&1 == 0 {
			if len(v.(string)) > max {
				max = len(v.(string))
			}
		}
	}
	for i, v := range keyvals {
		if i&1 == 0 {
			addings := max - len(v.(string)) + 1
			s += fmt.Sprintf("%v:%s", v, strings.Repeat(" ", addings))
		} else {
			s += fmt.Sprintf("%v\n", v)
		}
	}
	return
}
// https://go.dev/play/p/VdsqNKtr9GX
func GenUsageBar(total, used float64, size int) string {
	color := 32
	fill, empty := "\u25b0", "\u25b1"
	percent := float64((100*used) / total)
	if percent >= 66.666666 { color = 31 } else
	if percent >= 33.333333 { color = 33 }
	amount := int(percent / (100 / float64(size)))
	return fmt.Sprintf("\x1b[%dm%s%s\x1b[0m %.2f%%",
		color,
		strings.Repeat(fill, amount),
		strings.Repeat(empty, size-amount),
		percent,
	)	
}
