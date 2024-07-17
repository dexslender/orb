package commands

import (
	"fmt"
	"time"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
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

func (c *Purge) Run(cctx *util.CommandContext) error {
	amount := cctx.SlashCommandInteractionData().Int("amount")
	err := cctx.DeferCreateMessage(false)
	if err != nil {
		return err
	}
	msg, err := cctx.GetInteractionResponse()
	if err != nil {
		return err
	}
	msgs, err := cctx.Client().Rest().GetMessages(
		cctx.Channel().ID(),
		0,
		msg.ID,
		0,
		amount,
	)
	if err != nil {
		return err
	}
	if len(msgs) <= 1 {
		_, err = cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContentf("Only 1 message detected :(\n||U can delete it :)||").
			Build(),
		)
		return err
	}
	var (
		toDel []snowflake.ID
		no    int
		other string
	)
	days := 14 * 24 * time.Hour
	for _, msg := range msgs {
		if msg.CreatedAt.Before(time.Now().Add(-days)) {
			no += 1
			continue
		}
		toDel = append(toDel, msg.ID)
	}
	if no > 0 {
		other = fmt.Sprintf("\nskiped %d messages very old", no)
	}
	if len(toDel) <= 0 {
		_, err := cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContentf("Nothing to do :(%s", other).
			Build())
		return err
	}

	cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		SetContentf("Deleting %d messages...", len(toDel)).
		Build(),
	)
	err = cctx.Client().Rest().BulkDeleteMessages(
		cctx.Channel().ID(),
		toDel,
	)
	if err != nil {
		_, err := cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContentf("<:signerror:1071603595067265064> Error\n```go\n%s```", err).
			Build(),
		)
		return err
	}

	autodelete := time.Second * 10
	_, err = cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		SetContentf("Deleted %d messages...%s\nAutodelete %s",
			len(toDel),
			other,
			discord.NewTimestamp(
				discord.TimestampStyleRelative,
				time.Now().Add(autodelete))).
		Build(),
	)
	go func(task *util.CommandContext) {
		// TODO: maybe in local little bit slow to send delete req
		time.Sleep(autodelete - time.Second)
		cctx.DeleteInteractionResponse()
	}(cctx)

	return err
}
