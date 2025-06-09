package orb

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/gateway"
)

func New(ver string, logger *log.Logger, config *Config) *Orb {
	return &Orb{
		Config:  config,
		Log:     logger,
		Version: ver,
	}
}

type Orb struct {
	*bot.Client
	ActivityManager
	Config  *Config
	Log     *log.Logger
	Version string
}

func (o *Orb) Setup() {
	var err error
	o.Client, err = disgo.New(o.Config.Bot.Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentsNonPrivileged,
				gateway.IntentGuilds,
			),
			gateway.WithCompress(true),
			gateway.WithBrowser(func() string {
				if o.Config.ActivityManager.OnlineMobil {
					return "Discord Android"
				}
				return ""
			} ()),
			gateway.WithPresenceOpts(o.SetupActivity()),
		),
		bot.WithLogger(slog.New(o.Log)),
		bot.WithEventListeners(listeners(o)),
		bot.WithCacheConfigOpts(
			cache.WithCaches(cache.FlagGuilds|cache.FlagMembers),
		),
	)
	if err != nil {
		o.Log.Fatal("client error", "err", err)
	}
}

func (o *Orb) SetActivityManager(m ActivityManager) {
	o.ActivityManager = m
}

func (o *Orb) StartNLock() {
	ctx, c := context.WithTimeout(context.Background(), time.Second*10)
	defer c()
	defer func() {
		o.Close(ctx)
		o.Log.Info("client closed, bye...")
	}()

	err := o.OpenGateway(ctx)
	if err != nil {
		o.Log.Fatal("gateway open error", "err", err)
	}

	o.Log.Debug("Bot startup finished, runtime locked")
	k := make(chan os.Signal, 1)
	signal.Notify(k, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-k
}
