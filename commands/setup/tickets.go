package setup

import (
	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
)

var TicketsCommand = discord.ApplicationCommandOptionSubCommand{
	Name:        "tickets",
	Description: "just setup/configure ticket system (/info about:tickets)",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionChannel{
			Name:         "channel",
			Description:  "channel to assign ticket system, will created automatically if not assigned.",
			Required:     false,
			ChannelTypes: []discord.ChannelType{discord.ChannelTypeGuildText},
		},
		discord.ApplicationCommandOptionRole{ // To database
			Name:        "role",
			Description: "role for support ticket system (moderators, specials)",
			Required:    false,
		},
		discord.ApplicationCommandOptionChannel{ // To database
			Name:        "category",
			Description: "custom category to assign ticket system, will created automatically if not assigned.",
			Required:    false,
		},
		// discord.ApplicationCommandOptionString{      // NOTE: Buttons aren't supported on custom webhooks :(
		// 	Name:        "webhook",                 // NOTE: Maybe creating identical webhook but by the client
		// 	Description: "custom webhook",
		// 	Required:    false,
		// },
	},
}

var ticketChannelDefaultConfig discord.GuildChannelCreate = discord.GuildTextChannelCreate{
	Name: "tickets",
	Topic: "Create a new ticket. Contact server mods and support team from here!",
	// PermissionOverwrites: [], // TODO: needed to prevent messages from unauthorized users
}

func RunTickets(cctx *util.CommandContext) error {
	err := cctx.DeferCreateMessage(false)
	if err != nil { return err }
	var action int
	if _, ok := cctx.SlashCommandInteractionData().OptChannel("channel"); ok { action = 2 } else 
	if _, ok := cctx.SlashCommandInteractionData().OptString("webhook"); ok { action = 3 } else 
	{ action = 1 }
	return HandleAction(cctx, action)
}

func HandleAction(cctx *util.CommandContext, action int) error {
	return nil
}

func OnClickTicket(cctx *util.ComponentContext) error {
	cctx.Logger.Info("received interaction request from ticket button")
	return nil
}

/*
if channel, ok := cctx.SlashCommandInteractionData().OptChannel("channel"); ok {
		err := cctx.Client().Rest().SendTyping(channel.ID)
		if err != nil {
			_, err := cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
				SetContentf("<:signerror:1071603595067265064> Error\n```go\n%s```", err).
				Build(),
			)
			return err
		}
		msg, _ := cctx.Client().Rest().CreateMessage(
			channel.ID,
			discord.NewMessageCreateBuilder().
				AddEmbeds(discord.NewEmbedBuilder().
					SetAuthorName("Ticket System").
					SetAuthorIcon("https://media.discordapp.net/attachments/930845545138909195/1261429159163461632/emoji.png?ex=6692ecf3&is=66919b73&hm=855c1992eb1c215dbbf0b533c0fa816dfd6694231075710c98efe7ba11de51dd&=&format=webp&quality=lossless&width=417&height=417").
					SetTitle("Click to open new ticket!").
					SetColor(0x2B2D31).
					SetFooterText("please read the rules").
					Build()).
				AddActionRow(discord.NewSecondaryButton("Open", "ticket-open").
					WithEmoji(discord.ComponentEmoji{1071606184693481502, "signadd", false})).
				Build())
		_, err = cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContentf("```yaml\nTicket System\nCreated Message: %d```", msg.ID).
			Build())
		return err
	} else if wurl, ok := cctx.SlashCommandInteractionData().OptString("webhook-url"); ok {
		wh, err := webhook.NewWithURL(wurl)
		if err != nil {
			_, err := cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
				SetContentf("<:signerror:1071603595067265064> Error\n```go\n%s```", err).
				Build(),
			)
			return err

		}
		msg, err := cctx.Client().Rest().CreateWebhookMessage(wh.ID(), wh.Token(), discord.NewWebhookMessageCreateBuilder().
			AddEmbeds(discord.NewEmbedBuilder().
				SetAuthorName("Ticket System").
				SetAuthorIcon("https://media.discordapp.net/attachments/930845545138909195/1261429159163461632/emoji.png?ex=6692ecf3&is=66919b73&hm=855c1992eb1c215dbbf0b533c0fa816dfd6694231075710c98efe7ba11de51dd&=&format=webp&quality=lossless&width=417&height=417").
				SetTitle("Click to open new ticket!").
				SetColor(0x2B2D31).
				SetFooterText("please read the rules").
				Build()).
			AddActionRow(discord.NewSecondaryButton("Open", "ticket-open").
				WithEmoji(discord.ComponentEmoji{1071606184693481502, "signadd", false})).
			Build(),
			true,
			0)
		if err != nil {
			return err
		}
		_, err = cctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContentf("```yaml\nTicket System\nCreated via Webhook Message: %d```", msg.ID).
			Build())
		return err
	} else {
		return nil
	}
*/