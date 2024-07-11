package util

import (
	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type InteractionManager interface {
	OnInteraction(data *events.InteractionCreate)
	AddCommands(...Command)
	SetupCommands(bot.Client)
}

type InteractionRegister interface {
	// Add Command handler to Manager
	Command(Command)
	Component(string, ComponentHandle)
	Autocomplete(string, AutocompleteHandle)
	Modal(string, ModalHandle)
	// TODO: Add specific as Button()
}

var _ InteractionManager = (*Imanager)(nil)
var _ InteractionRegister = (*Imanager)(nil)

type Imanager struct {
	Logger        *log.Logger
	Config        *Config
	interactions  []Command
	components    []Component
	autocompletes []Autocomplete
	modals        []Modal
}

func (m *Imanager) OnInteraction(data *events.InteractionCreate) {
	switch i := data.Interaction.(type) {
	case discord.ApplicationCommandInteraction:
		for _, cmd := range m.interactions {
			if cmd.CommandName() == i.Data.CommandName() {
				ctx := &CommandContext{
					events.ApplicationCommandInteractionCreate{
						GenericEvent:                  data.GenericEvent,
						Respond:                       data.Respond,
						ApplicationCommandInteraction: i,
					},
					m.Logger,
				}
				err := cmd.Run(ctx)
				if err != nil {
					cmd.Error(ctx, err)
				}
			}
		}
	case discord.ComponentInteraction:
		for _, comp := range m.components {
			if comp.CustomId == i.Data.CustomID() {
				ctx := &ComponentContext{
					events.ComponentInteractionCreate{
						GenericEvent:         data.GenericEvent,
						ComponentInteraction: i,
						Respond:              data.Respond,
					},
					m.Logger,
				}
				err := comp.Run(ctx)
				if err != nil {
					m.Logger.Error("component fail", "customId", i.Data.CustomID(), "err", err)
				}
			}
		}
	default:
		m.Logger.Warn("unhandled interaction", "type", i.Type())
	}
}

func (m *Imanager) SetupCommands(c bot.Client) {
	var up []discord.ApplicationCommandCreate
	for _, cmd := range m.interactions {
		up = append(up, cmd)
	}
	if m.Config.Bot.SetupCommands {
		var (
			err error
			ok  []discord.ApplicationCommand
		)
		if m.Config.Bot.GlobalCommands {
			ok, err = c.Rest().SetGlobalCommands(
				c.ApplicationID(),
				up,
			)
		} else if m.Config.Bot.GuildID != 0 {
			ok, err = c.Rest().SetGuildCommands(
				c.ApplicationID(),
				m.Config.Bot.GuildID,
				up,
			)

		}
		if m.Config.Bot.SetupCommands || m.Config.Bot.GuildID != 0 {
			if err != nil {
				m.Logger.Error("command setup fail", "err", err)
			} else {
				m.Logger.Info("command setup", "total", len(ok), "guild", m.Config.Bot.GuildID, "global", m.Config.Bot.SetupCommands)
			}
		}
	}
}

func (m *Imanager) AddCommands(cs ...Command) {
	for _, cmd := range cs {
		m.Command(cmd)
	}
}

func (m *Imanager) GetCommand(query string) Command {
	for _, v := range m.interactions {
		if v.CommandName() == query {
			return v
		}
	}
	return nil
}

// Add Command handler to Manager
func (m *Imanager) Command(cmd Command) {
	defer func() {
		if err := recover(); err != nil {
			m.Logger.Error("command failed to Init()", "command", cmd, "err", err)
		} else {
			m.interactions = append(m.interactions, cmd)
		}

	}()
	cmd.Init(m)
}

func (m *Imanager) Component(customId string, handle ComponentHandle) {
	if customId != "" {
		m.components = append(m.components, Component{customId, handle})
	}
}

func (m *Imanager) Autocomplete(_ string, _ AutocompleteHandle) {
	panic("not implemented") // TODO: Implement
}

func (m *Imanager) Modal(_ string, _ ModalHandle) {
	panic("not implemented") // TODO: Implement
}
