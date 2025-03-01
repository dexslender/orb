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

type Info struct { base; discord.SlashCommandCreate }

func (i *Info) Init(add util.InteractionRegister) {
	i.Name = "info"
	i.Description = "just returns bot's info and states"
	i.Options = []discord.ApplicationCommandOption{
		util.HiddenOpt,
		discord.ApplicationCommandOptionString{
			Name:         "about",
			Description:  "Get help about some topic",
			Required:     false,
			Autocomplete: true,
		},
	}

	add.Autocomplete(i, i.autocomplete)
}

func (i *Info) Run(ctx *util.CommandContext) error {
	hidden := ctx.SlashCommandInteractionData().Bool("hidden")
	// if key, ok := ctx.SlashCommandInteractionData().OptString("about"); ok {
	// 	return HandleAbout(ctx, key, hidden)
	// }
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
	var cpu float32
	if err != nil {
		ctx.Logger.Error("falied to get cpu percent usage", "err", err)
	} else {
			for _, core := range prc {
				cpu+=float32(core)
			}
			cpu/=float32(len(prc))
	}

	usage := fmt.Sprint(
		"```ansi\n",
		"Goroutines: ", runtime.NumGoroutine(),
		"\nMemory: ", GenUsageBar(float32(ms.Sys), float32(ms.TotalAlloc), 10),
		"\nCPU:    ", GenUsageBar(100, cpu, 10),
		"```",
	)

	_, err = ctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetAuthorName(cu.Tag()).
			SetAuthorIcon(*cu.AvatarURL()).
			AddFields(
				discord.EmbedField{Name: "Counts", Value: counts, Inline: json.Ptr(true)},
				discord.EmbedField{Name: "Stats", Value: usage, Inline: json.Ptr(false)},
			).
			SetFooterTextf("Version: %s", ctx.Orb.Version).
			SetColor(util.HIGHBLUE).
		Build()).
	Build())
	return err
}

func (i Info) autocomplete(ac *util.AutocompleteContext) (choices []discord.AutocompleteChoice) {
	return nil
	// query, ok := ac.Data.OptString("about")
	// if !ok { return nil }
	// if len(infos) <= 0 {
	// 	for _, c := range Commands {
	// 		if i := c.Info(); i != nil {
	// 			infos[i.Name] = i
	// 		}
	// 	}
	// }
	// for k, a := range infos {
	// 	if query != "" {
	// 		if strings.Contains(a.Name, query) {
	// 			choices = append(choices, discord.AutocompleteChoiceString{
	// 				Name:  a.Name,
	// 				Value: k,
	// 			})
	// 		}
	// 	} else {
	// 		choices = append(choices, discord.AutocompleteChoiceString{
	// 			Name:  a.Name,
	// 			Value: k,
	// 		})
	// 	}
	// }
	// return
}

// func HandleAbout(ctx *util.CommandContext, key string, hidden bool) error {
// 	if len(infos) <= 0 {
// 		for _, c := range Commands {
// 			if i := c.Info(); i != nil {
// 				infos[i.Name] = i
// 			}
// 		}
// 	}
// 	a := infos[key]
// 	return ctx.CreateMessage(
// 		discord.NewMessageCreateBuilder().
// 			AddEmbeds(discord.NewEmbedBuilder().
// 				SetTitle(a.Name).
// 				SetDescription(a.Description).
// 				AddFields(a.Details...).
// 				SetColor(util.LOWBLUE).
// 				Embed).
// 			SetEphemeral(hidden).
// 			MessageCreate)
// }

func FormatSpacing(keyvals ...any) (s string) {
	max := 0
	for i, v := range keyvals {
		if i&1 == 0 {
			if len(v.(string)) > max {
				max = len(v.(string))
			}
		}
	}
	for i, v := range keyvals {
		if i&1 == 0 {
			addings := max - len(v.(string)) + 1
			s += fmt.Sprintf("%v:%s", v, strings.Repeat(" ", addings))
		} else {
			s += fmt.Sprintf("%v\n", v)
		}
	}
	return
}
// https://go.dev/play/p/VdsqNKtr9GX
func GenUsageBar(total, used float32, size int) string {
	color := 32
	fill, empty := "\u25b0", "\u25b1"
	percent := float32((100*used) / total)
	if percent >= 66.666666 { color = 31 } else
	if percent >= 33.333333 { color = 33 }
	amount := int(percent / (100 / float32(size)))
	return fmt.Sprintf("\x1b[%dm%s%s\x1b[0m %.2f%%",
		color,
		strings.Repeat(fill, amount),
		strings.Repeat(empty, size-amount),
		percent,
	)	
}

