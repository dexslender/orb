package orb

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

type (
	Config struct {
		Bot struct {
			Token          string       `hcl:"token"`
			GuildID        snowflake.ID `hcl:"guild"`
			SetupCommands  bool         `hcl:"setup_commands"`
			GlobalCommands bool         `hcl:"global_commands"`
			LogLevel       log.Level    `hcl:"log_level"`
		} `hcl:"bot,block"`
		ActivityManager struct {
			Enabled     bool          `hcl:"enabled"`
			OnlineMobil bool          `hcl:"online_movil"`
			Interval    time.Duration `hcl:"interval"`
			Activities  []Activity    `hcl:"activity,block"`
		} `hcl:"activity_manager,block"`
	}
	Activity struct {
		Message string                `hcl:",label"`
		Status  *discord.OnlineStatus `hcl:"status"`
		Type    *discord.ActivityType `hcl:"type"`
		URL     *string               `hcl:"url"`
	}
)

func HCLctx(l *log.Logger) *hcl.EvalContext {
	return &hcl.EvalContext{
		Variables: map[string]cty.Value{
			//----LogLevels
			"debug": cty.NumberIntVal(int64(log.DebugLevel)),
			"info":  cty.NumberIntVal(int64(log.InfoLevel)),
			"warn":  cty.NumberIntVal(int64(log.WarnLevel)),
			"error": cty.NumberIntVal(int64(log.ErrorLevel)),
			"fatal": cty.NumberIntVal(int64(log.FatalLevel)),
			//----Activity Type
			"playing":   cty.NumberIntVal(int64(discord.ActivityTypeGame)),
			"streaming": cty.NumberIntVal(int64(discord.ActivityTypeStreaming)),
			"listening": cty.NumberIntVal(int64(discord.ActivityTypeListening)),
			"watching":  cty.NumberIntVal(int64(discord.ActivityTypeWatching)),
			"custom":    cty.NumberIntVal(int64(discord.ActivityTypeCustom)),
			"competing": cty.NumberIntVal(int64(discord.ActivityTypeCompeting)),
			//----Statuses
			"online":    cty.StringVal(string(discord.OnlineStatusOnline)),
			"dnd":       cty.StringVal(string(discord.OnlineStatusDND)),
			"idle":      cty.StringVal(string(discord.OnlineStatusIdle)),
			"invisible": cty.StringVal(string(discord.OnlineStatusInvisible)),
			"offline":   cty.StringVal(string(discord.OnlineStatusOffline)),
		},
		Functions: map[string]function.Function{
			"env": function.New(&function.Spec{
				Params: []function.Parameter{{
					Name:             "key",
					Type:             cty.String,
					AllowDynamicType: true,
				}},
				Type: function.StaticReturnType(cty.String),
				Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
					k := args[0].AsString()
					v := os.Getenv(k)
					if v == "" {
						l.Warn("empty env-var", k, v)
					}
					return cty.StringVal(v), nil
				},
			}),
			"duration": function.New(&function.Spec{
				Params: []function.Parameter{{
					Name:             "format",
					Type:             cty.String,
					AllowDynamicType: true,
				}},
				Type: function.StaticReturnType(cty.Number),
				Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
					out, err := time.ParseDuration(args[0].AsString())
					return cty.NumberIntVal(int64(out)), err
				},
			}),
		},
	}
}
