package util

import (
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"github.com/kkyr/fig"
)

type Config struct {
	Bot struct {
		Token          string `validate:"required"`
		ClientSecret   string `fig:"client_secret"`
		GuildId        snowflake.ID
		SetupCommands  bool
		GlobalCommands bool
		MobileOs       bool
		DebugLog       bool
	}
	PresenceUpdater struct {
		Enabled   bool
		Delay     time.Duration `default:"10s"`
		Presences []Presence
	}
}

type Presence struct {
	Status                 discord.OnlineStatus
	Name, Type, URL, State string
}

func LoadConfig(filename, env string) (config Config, err error) {
	err = fig.Load(&config,
		fig.File(filename),
		fig.UseEnv(env),
	)
	return
}
