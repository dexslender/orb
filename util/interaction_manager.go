package util

import (
	"github.com/charmbracelet/log"
	"github.com/dexslender/orb/orb"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type Imanager struct {
	*orb.Orb
	Logger        *log.Logger
	Config        *orb.Config
	autocompletes []Autocomplete
	components    []Component
	interactions  []Command
	modals        []Modal
}

func (m *Imanager) OnInteraction(data *events.InteractionCreate) {
	switch i := data.Interaction.(type) {
	case discord.ApplicationCommandInteraction:
		for _, cmd := range m.interactions {
			if cmd.CommandName() == i.Data.CommandName() {
				executor := func(
					man *Imanager,
					data *events.InteractionCreate,
					i discord.ApplicationCommandInteraction,
					com Command,
				) {
					ctx := &CommandContext{
						man.Orb,
						events.ApplicationCommandInteractionCreate{
							GenericEvent:                  data.GenericEvent,
							Respond:                       data.Respond,
							ApplicationCommandInteraction: i,
						},
						man.Logger,
					}
					err := com.Run(ctx)
					if err != nil {
						com.Error(ctx, err)
					}
				}
				go executor(m, data, i, cmd)
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
	case discord.AutocompleteInteraction:
		for _, ac := range m.autocompletes {
			if ac.Command.CommandName() == i.Data.CommandName {
				executor := func(
					data *events.InteractionCreate,
					i discord.AutocompleteInteraction,
					m *Imanager,
					ac Autocomplete,
				) {
					ctx := &AutocompleteContext{
						events.AutocompleteInteractionCreate{GenericEvent: data.GenericEvent, AutocompleteInteraction: i, Respond: data.Respond},
						m.Logger,
					}
					data.Respond(
						discord.InteractionResponseTypeAutocompleteResult,
						discord.AutocompleteResult{Choices: ac.Run(ctx)},
					)
				}
				go executor(data, i, m, ac)
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
				m.Logger.Info("command setup", "total", len(ok), "guild", m.Config.Bot.GuildID, "global", m.Config.Bot.GlobalCommands)
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
