package util

import (
	"github.com/charmbracelet/log"
	"github.com/disgoorg/snowflake/v2"
)

type Config struct {
	Bot struct {
		Token          string `validate:"required"`
		GuildID        snowflake.ID
		SetupCommands  bool
		GlobalCommands bool
		LogLevel       log.Level
	}
}
