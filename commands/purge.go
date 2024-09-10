package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/json"
	"github.com/disgoorg/snowflake/v2"
)

type Purge struct {
	base
	discord.SlashCommandCreate
}

func (c *Purge) Init(util.InteractionRegister) {
	c.Name = "purge"
	c.Description = "just deletes messages from current channel"
	c.DefaultMemberPermissions = json.NewNullablePtr(discord.PermissionManageMessages)
	c.Options = []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionInt{
			Name:        "amount",
			Description: "Amount of messages of delete.",
			Required:    true,
			MinValue:    json.Ptr(2),
			MaxValue:    json.Ptr(100),
		},
	}
}

const (
	autodelete = time.Second * 10
	limit      = 14 * 24 * time.Hour
)

func (c *Purge) Run(cctx *util.CommandContext) error {
	amount := cctx.SlashCommandInteractionData().Int("amount")
	err := cctx.DeferCreateMessage(false)
	if err != nil { return err }

	msg, err := cctx.GetInteractionResponse()
	if err != nil { return err }

	msgs, err := cctx.Client().Rest().
		GetMessages(cctx.Channel().ID(), 0, msg.ID, 0, amount)
	if err != nil { return err }

	var (
		deleting []snowflake.ID
		old      int
		errorMsg string
	)

	for _, m := range msgs {
		if m.CreatedAt.Before(time.Now().Add(-limit)) {
			old += 1
			continue
		}
		deleting = append(deleting, m.ID)
	}

	if len(deleting) <= 0 {
		errorMsg = "Nothing to do o.o"
		if old > 0 {
			errorMsg += fmt.Sprintf("\nskipped %d messages too old", old)
		}
	} else if len(deleting) == 1 {
		errorMsg = "1 message detected, you can delete it :)"
		if old > 0 {
			errorMsg += fmt.Sprintf("\nskipped %d messages too old", old)
		}
	}

	if errorMsg != "" {
		_, err := cctx.UpdateInteractionResponse(discord.MessageUpdate{
			Content: &errorMsg,
		})
		return err
	}

	if old > 0 {
		_, err := cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContentf("```go\ndelete %d messages?\nskipped %d too old```", len(deleting), old).
			AddActionRow(discord.NewPrimaryButton("Yes", "purge-yes"), discord.NewSecondaryButton("No", "purge-no")).
			Build())
		if err != nil { return err }

		action, close := bot.NewEventCollector[*events.ComponentInteractionCreate](
			cctx.Orb,
			func(e *events.ComponentInteractionCreate) bool { 
				return e.Message.ID == msg.ID && e.User().ID == cctx.User().ID
			 },
		)
		defer close()
		ctx, cl := context.WithTimeout(context.Background(), time.Second*10)
		defer cl()
		select {
		case btn := <-action:
			err := btn.DeferUpdateMessage()
			if err != nil { return err }
			switch btn.Data.CustomID() {
			case "purge-yes":
				cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
					SetContentf("Deleting %d messages...", len(deleting)).
					ClearContainerComponents().
					Build())
				err = cctx.Client().Rest().
					BulkDeleteMessages(cctx.Channel().ID(), deleting)
				if err != nil {
					return err
				}
				_, err = cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
					SetContentf(
						"Deleted %d messages ||destroying %s||",
						len(deleting),
						discord.NewTimestamp(
							discord.TimestampStyleRelative,
							time.Now().Add(autodelete),
						),
					).
					Build())
				go func(cctx *util.CommandContext) {
					time.Sleep(autodelete)
					cctx.DeleteInteractionResponse()
				}(cctx)
				return err
			case "purge-no":
				_, err := cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
					SetContent("Ok, cancelled :)").
					ClearContainerComponents().
					Build())
				return err
			default: return nil
			}
		case <-ctx.Done():
			_, err := cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
				SetContent("Ok, doing nothing").
				ClearContainerComponents().
				Build())
			return err
		}
	} else {
		cctx.UpdateInteractionResponse(discord.MessageUpdate{
			Content: json.Ptr(fmt.Sprintf("Deleting %d messages...", len(deleting))),
		})
		err = cctx.Client().Rest().
			BulkDeleteMessages(cctx.Channel().ID(), deleting)
		if err != nil {
			return err
		}
		_, err = cctx.UpdateInteractionResponse(discord.MessageUpdate{
			Content: json.Ptr(fmt.Sprintf("Deleted %d messages ||destroying %s||",
				len(deleting),
				discord.NewTimestamp(discord.TimestampStyleRelative, time.Now().Add(autodelete))),
			),
		})
		go func(cctx *util.CommandContext) {
			time.Sleep(autodelete)
			cctx.DeleteInteractionResponse()
		}(cctx)
		return err
	}
}
