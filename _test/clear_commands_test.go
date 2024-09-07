package test_test

import (
	"testing"

	"github.com/dexslender/orb/orb"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/discord"
	"github.com/kkyr/fig"
)

func TestClearCommands(t *testing.T) {
	var config orb.Config
	err := fig.Load(&config,
		fig.File("botconfig.yml"),
		fig.UseEnv("ORB"),
	)
	if err != nil {
		t.Fatal("config error: ", err)
	}

	client, err := disgo.New(config.Bot.Token)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Rest().SetGuildCommands(client.ApplicationID(), config.Bot.GuildID, []discord.ApplicationCommandCreate{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClearGlobal(t *testing.T) {
	var config orb.Config
	err := fig.Load(&config,
		fig.File("botconfig.yml"),
		fig.UseEnv("ORB"),
	)
	if err != nil {
		t.Fatal("config error: ", err)
	}

	client, err := disgo.New(config.Bot.Token)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Rest().SetGlobalCommands(client.ApplicationID(), []discord.ApplicationCommandCreate{})
	if err != nil {
		t.Fatal(err)
	}
}
