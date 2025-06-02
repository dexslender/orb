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

	// cUser, err := ctx.Client().Rest.GetCurrentUser("")
	// if err != nil { return err }

	counts := fmt.Sprintf("```js\n%s```", util.FormatSpacing(
		"Guilds", ctx.Client().Caches.GuildsLen(),
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
		AddComponents(discord.NewContainer(
			discord.NewTextDisplay("**Counts**\n"+counts),
			discord.NewTextDisplay("**Stats**"+usage),
			discord.NewSmallSeparator(),
			discord.NewTextDisplayf("-# Version: %s", ctx.Orb.Version),
		)).
		AddFlags(discord.MessageFlagIsComponentsV2).
		Build())
	return err
}
