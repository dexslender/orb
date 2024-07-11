package util

import (
	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/rest"
	"github.com/disgoorg/snowflake/v2"
)

// ----Component
type Component struct {
	CustomId string
	Run      ComponentHandle
}
type ComponentContext struct {
	events.ComponentInteractionCreate
	Logger *log.Logger
}

func (e *ComponentContext) GetInteractionResponse(opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().GetInteractionResponse(e.ApplicationID(), e.Token(), opts...)
}

func (e *ComponentContext) UpdateInteractionResponse(messageUpdate discord.MessageUpdate, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().UpdateInteractionResponse(e.ApplicationID(), e.Token(), messageUpdate, opts...)
}

func (e *ComponentContext) DeleteInteractionResponse(opts ...rest.RequestOpt) error {
	return e.Client().Rest().DeleteInteractionResponse(e.ApplicationID(), e.Token(), opts...)
}

func (e *ComponentContext) GetFollowupMessage(messageID snowflake.ID, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().GetFollowupMessage(e.ApplicationID(), e.Token(), messageID, opts...)
}

func (e *ComponentContext) CreateFollowupMessage(messageCreate discord.MessageCreate, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().CreateFollowupMessage(e.ApplicationID(), e.Token(), messageCreate, opts...)
}

func (e *ComponentContext) UpdateFollowupMessage(messageID snowflake.ID, messageUpdate discord.MessageUpdate, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().UpdateFollowupMessage(e.ApplicationID(), e.Token(), messageID, messageUpdate, opts...)
}

func (e *ComponentContext) DeleteFollowupMessage(messageID snowflake.ID, opts ...rest.RequestOpt) error {
	return e.Client().Rest().DeleteFollowupMessage(e.ApplicationID(), e.Token(), messageID, opts...)
}

type ComponentHandle func(*ComponentContext) error

// ----Autocomplete
type Autocomplete struct {
	Command string
	Run     AutocompleteHandle
}
type AutocompleteContext struct {
	events.AutocompleteInteractionCreate
	Logger *log.Logger
}
type AutocompleteHandle func(*AutocompleteContext) discord.AutocompleteResult

// ----Modal
type Modal struct {
	CustomID string
	Run      ModalHandle
}
type ModalContext struct {
	events.ModalSubmitInteractionCreate
	Logger *log.Logger
}
type ModalHandle func(*ModalContext) error
