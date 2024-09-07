package util

import (
	"context"

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
	Command
	// Opt string
	Run AutocompleteHandle
}
type AutocompleteContext struct {
	events.AutocompleteInteractionCreate
	Logger *log.Logger
}
type AutocompleteHandle func(*AutocompleteContext) []discord.AutocompleteChoice

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

type InteractionPayload[I discord.Interaction] struct {
	*events.GenericEvent
	events.InteractionResponderFunc
	Interaction I
}

type Task interface {
	Deleteable() bool
	OnInteraction(*events.InteractionCreate)
}

type InteractionTask[I discord.Interaction] struct {
	filter     func(InteractionPayload[I]) bool
	c          chan<- InteractionPayload[I]
	deleteable bool
}

func (i *InteractionTask[I]) OnInteraction(data *events.InteractionCreate) {
	if _, ok := data.Interaction.(I); !ok { return }
	p := InteractionPayload[I]{
		data.GenericEvent,
		data.Respond,
		data.Interaction.(I),
	}
	if i.filter(p) {
		i.c <- p
	}
}

func (i InteractionTask[I]) Deleteable() bool {
	return i.deleteable
}

func MakeInteractionTask[I discord.Interaction](ctx context.Context, filter func(InteractionPayload[I]) bool, c chan InteractionPayload[I]) *InteractionTask[I] {
	it := &InteractionTask[I]{
		filter: filter,
		c:      c,
	}
	go func(ctx context.Context, it *InteractionTask[I]) {
		if v := ctx.Done(); v != nil {
			<-v
			it.deleteable = true
		}
	}(ctx, it)
	return it
}
