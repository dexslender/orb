package about

import (
	"context"
	"fmt"
	"time"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var SubGuild = &discord.ApplicationCommandOptionSubCommand{
	Name: "guild",
	Description: "get info and stats about this guild",
	Options: []discord.ApplicationCommandOption{ util.HiddenOpt },
}

func GuildRun(ctx *util.CommandContext) error {
	guild, ok := ctx.Guild()
	if !ok { 
		return ctx.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent("Not available here!").
		Build())
	}

	desc := ""
	if guild.Description != nil {
		desc = "\n```"+*guild.Description+"```"
	}
	
	basic_info := fmt.Sprint(
		"**Created**: ", discord.NewTimestamp(discord.TimestampStyleRelative, guild.CreatedAt()),
		"\n**Id**: ", guild.ID,
		"\n**Owner**: ", fmt.Sprintf("<@%s>",guild.OwnerID),
		"\n**Locale**: ", guild.PreferredLocale,
		desc,
	)

	actions := discord.NewActionRow(
		discord.NewSecondaryButton("Stats", "guild-stats"),
		discord.NewSecondaryButton("Features", "guild-features"),
		// discord.NewSecondaryButton("", "guild-level")
	)

	err := ctx.CreateMessage(discord.NewMessageCreateBuilder().
		SetEphemeral(ctx.SlashCommandInteractionData().Bool("hidden")).
		AddEmbeds(discord.NewEmbedBuilder().
			SetThumbnail(*guild.IconURL()).
			SetTitle(guild.Name).
			SetDescription(basic_info).
		Build()).
		AddContainerComponents(actions).
	Build())
	if err != nil { return err }

	msg, err := ctx.GetInteractionResponse();
	if err != nil { return err }

	action, clos := bot.NewEventCollector(
		ctx.Orb,
		func(e *events.ComponentInteractionCreate) bool {
			return e.Message.ID == msg.ID && e.User().ID == ctx.User().ID 
		},
	)
	defer clos()
	Tctx, Tclos := context.WithTimeout(context.Background(), time.Minute*5)
	defer Tclos()
	select {
	case button := <-action:
		err := button.DeferUpdateMessage()
		if err != nil { return err }
		switch button.Data.CustomID() {
		case "guild-stats":
			return UpdateWithGuildStats(button, guild)
		default:
			return button.DeferUpdateMessage()
		}
	case <-Tctx.Done():
		_, err := ctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			ClearContainerComponents().
			Build())
		return err
	}
}

func UpdateWithGuildStats(button *events.ComponentInteractionCreate, guild discord.Guild) error {
	return nil
}

func UpdateWithGuildFeatures() error {
	return nil
}
