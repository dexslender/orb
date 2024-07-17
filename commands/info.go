package commands

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/json"
	"github.com/shirou/gopsutil/v4/cpu"
)

type Info struct {
	base
	discord.SlashCommandCreate
}

func (i *Info) Init(util.InteractionRegister) {
	i.Name = "info"
	i.Description = "just returns bot's info and states"
	i.Options = []discord.ApplicationCommandOption{util.HiddenOpt}
}

func (i *Info) Run(ctx *util.CommandContext) error {
	hidden := ctx.SlashCommandInteractionData().Bool("hidden")
	err := ctx.DeferCreateMessage(hidden)
	if err != nil {
		return err
	}

	cu, err := ctx.Client().Rest().GetCurrentUser("")
	if err != nil {
		return err
	}

	counts := fmt.Sprintf("```js\n%s```", FormatSpacing(
		"Guilds", ctx.Orb.Caches().GuildsLen(),
		"Commands", len(Commands),
	))

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	prc, err := cpu.Percent(time.Second, true)
	var cpu string
	if err != nil {
		ctx.Logger.Error("falied to get cpu percent usage", "err", err)
	} else {
		for i, v := range prc {
			cpu += fmt.Sprintf("\tcore %d: %.2f%%\n", i, v)
		}
	}

	usage := fmt.Sprintf("```js\n%s```", FormatSpacing(
		"Goroutines", runtime.NumGoroutine(),
		"Memory", ms.TotalAlloc,
		"CPU", "\n"+cpu,
	))

	info := discord.NewEmbedBuilder().
		SetAuthorName(cu.Tag()).
		SetAuthorIcon(*cu.AvatarURL()).
		AddFields(
			discord.EmbedField{Name: "Counts", Value: counts, Inline: json.Ptr(true)},
			discord.EmbedField{Name: "Stats", Value: usage, Inline: json.Ptr(true)},
		).
		SetFooterTextf("Version: %s", ctx.Orb.Version).
		Build()

	_, err = ctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		AddEmbeds(info).
		Build())
	return err
}

func FormatSpacing(keyvals ...any) (s string) {
	max := 0
	for i, v := range keyvals {
		if i%2 != 0 {
			continue
		}
		if len(v.(string)) > max {
			max = len(v.(string))
		}
	}
	for i, v := range keyvals {
		if i%2 == 0 {
			addings := max - len(v.(string)) + 1
			s += fmt.Sprintf("%v:%s", v, strings.Repeat(" ", addings))
		} else {
			s += fmt.Sprintf("%v\n", v)
		}
	}
	return
}
