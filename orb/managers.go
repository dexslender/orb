package orb

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
)

type ActivityManager interface {
	SetupActivity() gateway.ConfigOpt
	StartActivityUpdater(bot.Client)
}

type InteractionManager interface {
	OnInteraction(data *events.InteractionCreate)
	SetupCommands(bot.Client)
}
