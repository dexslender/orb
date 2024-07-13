package orb

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo/discord"
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
	ActivityManager struct {
		Enabled     bool
		OnlineMobil bool
		Interval    time.Duration `fig:"delay" default:"10s"`
		Activities  []Activity
	}
}

type Activity struct {
	Status                 discord.OnlineStatus
	Name, Type, URL, State string
}
