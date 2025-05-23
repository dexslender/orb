package about

import (
	"fmt"
	"runtime"
	"time"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/shirou/gopsutil/v4/cpu"
)

var SubBotStats = &discord.ApplicationCommandOptionSubCommand{
	Name: "botstats",
	Description: "returns some live bot stats",
	Options: []discord.ApplicationCommandOption{
		util.HiddenOpt,
	},
}

func BotStatsRun(ctx *util.CommandContext, cmdLen int) error {
	hidden := ctx.SlashCommandInteractionData().Bool("hidden")

	err := ctx.DeferCreateMessage(hidden)
	if err != nil { return err }

	cUser, err := ctx.Client().Rest().GetCurrentUser("")
	if err != nil { return err }

	counts := fmt.Sprintf("```js\n%s```", util.FormatSpacing(
		"Guilds", ctx.Orb.Caches().GuildsLen(),
		"Commands", cmdLen,
	))

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	prc, err := cpu.Percent(time.Second, true)
	var cpu float64
	if err != nil {
		ctx.Logger.Error("failed to get cpu percent usage", "err", err)
	} else {
		for _, core := range prc {
			cpu+=float64(core)
		}
		cpu/=float64(len(prc))
	}

	usage := fmt.Sprint(
		"```ansi",
		"\nGoroutines: ", runtime.NumGoroutine(),
		"\nMemory: ", util.GenUsageBar(float64(ms.Sys), float64(ms.TotalAlloc), 10),
		"\nCPU:    ", util.GenUsageBar(100, cpu, 10),
		"```",
	)

	_, err = ctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetAuthorName(cUser.Tag()).
			SetAuthorIcon(*cUser.AvatarURL()).
			AddFields(
				discord.EmbedField{Name: "Counts", Value: counts},
				discord.EmbedField{Name: "Stats", Value: usage},
			).
			SetFooterText("Version: "+ctx.Orb.Version).
			SetColor(util.HIGHBLUE).
		Build()).
	Build())
	return err
}
