package about

import (
	"fmt"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
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

	btnStats := discord.NewSecondaryButton("Stats", fmt.Sprintf("/about/guild/%d", guild.ID))
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
	return err

	// msg, err := ctx.GetInteractionResponse();
	// if err != nil { return err }

	// action, clos := bot.NewEventCollector(
	// 	ctx.Client(),
	// 	func(e *events.ComponentInteractionCreate) bool {
	// 		return e.Message.ID == msg.ID && e.User().ID == ctx.User().ID 
	// 	},
	// )
	// defer clos()
	// Tctx, Tclos := context.WithTimeout(context.Background(), time.Minute*10)
	// defer Tclos()
	// select {
	// case button := <-action:
	// 	err := button.DeferUpdateMessage()
	// 	if err != nil { return err }
	// 	switch button.Data.CustomID() {
	// 	case "guild-stats":
	// 		return UpdateWithGuildStats(button, guild)
	// 	default:
	// 		return button.DeferUpdateMessage()
	// 	}
	// case <-Tctx.Done():
	// 	btnStats.Disabled = true
	// 	btnFeats.Disabled = true
	// 	_, err := ctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
	// 		SetComponents(components).
	// 	Build())
	// 	return err
	// }
}

func UpdateWithGuildStats(data discord.ButtonInteractionData, e *handler.ComponentEvent) error {
	guild, ok := e.Guild()
	if !ok { return e.DeferUpdateMessage() }

	err := e.DeferCreateMessage(true)
	if err != nil { return err }

	

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetIsComponentsV2(true).
		AddComponents(discord.NewTextDisplay(*guild.Banner)).
	Build())
}

func UpdateWithGuildFeatures() error {
	return nil
}
