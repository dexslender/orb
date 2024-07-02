package util

import (
	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type CommandManager interface {
	OnInteraction(data *events.InteractionCreate)
	AddCommands(...Command)
}

var _ CommandManager = (*Cmanager)(nil)

type Cmanager struct {
	Logger *log.Logger
}

func (m *Cmanager) OnInteraction(data *events.InteractionCreate) {
	switch i := data.Interaction.(type) {
	case discord.ApplicationCommandInteraction:
		m.Logger.Info("received app command interaction")
	default:
		m.Logger.Warn("unhandled interaction", "type", i.Type())
	}
}

func (m *Cmanager) AddCommands(cs ...Command) {

}
