package about

import (
	"context"
	"fmt"
	"time"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/handler"
)

var SubGuild = &discord.ApplicationCommandOptionSubCommand{
	Name: "guild",
	Description: "get info and stats about this guild",
	Options: []discord.ApplicationCommandOption{ util.HiddenOpt },
}

func GuildRun(ctx *handler.CommandEvent) error {
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

	btnStats := discord.NewSecondaryButton("Stats", "guild-stats")
	btnFeats := discord.NewSecondaryButton("Features", "guild-features")
	
	components := discord.NewContainer(
		discord.NewSection(discord.NewTextDisplayf(
			"## **%s**\n%s",
			guild.Name, basic_info,
		)).WithAccessory(discord.NewThumbnail(*guild.IconURL())),
		discord.NewSmallSeparator(),
		discord.NewActionRow(&btnStats, &btnFeats),
	)

	err := ctx.CreateMessage(discord.NewMessageCreateBuilder().	
		AddComponents(components).
		SetIsComponentsV2(true).
	Build())
	if err != nil { return err }

	msg, err := ctx.GetInteractionResponse();
	if err != nil { return err }

	action, clos := bot.NewEventCollector(
		ctx.Client(),
		func(e *events.ComponentInteractionCreate) bool {
			return e.Message.ID == msg.ID && e.User().ID == ctx.User().ID 
		},
	)
	defer clos()
	Tctx, Tclos := context.WithTimeout(context.Background(), time.Minute*10)
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
		btnStats.Disabled = true
		btnFeats.Disabled = true
		_, err := ctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetComponents(components).
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
