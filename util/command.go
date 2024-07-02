package util

import "github.com/disgoorg/disgo/discord"

type Command interface {
	discord.ApplicationCommandCreate
	Init()
	Run() error
	Error()
}

type Ccontext struct {
	discord.ApplicationCommandInteraction
}
