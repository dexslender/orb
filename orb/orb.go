package orb

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/dexslender/orb/util"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
)

func New(logger *log.Logger, config *util.Config) *Orb {
	return &Orb{
		Config: config,
		Log:    logger,
	}
}

type Orb struct {
	bot.Client
	util.InteractionManager
	util.ActivityManager
	Config *util.Config
	Log    *log.Logger
}

func (o *Orb) Setup() {
	var err error
	o.Client, err = disgo.New(o.Config.Bot.Token,
		bot.WithGatewayConfigOpts(
			func(config *gateway.Config) {
				if o.Config.ActivityManager.OnlineMobil {
					config.Browser = "Discord Android"
				}
			},
			gateway.WithIntents(
				gateway.IntentsNonPrivileged,
				gateway.IntentGuilds,
			),
			gateway.WithCompress(true),
			o.SetupActivity,
		),
		bot.WithLogger(slog.New(o.Log)),
		bot.WithEventListeners(listeners(o)),
	)
	if err != nil {
		o.Log.Fatal("client error", "err", err)
	}
	o.SetupCommands(o)
	o.StartNLock()
}

func (o *Orb) SetCommandManager(m util.InteractionManager) {
	o.InteractionManager = m
}

func (o *Orb) SetActivityManager(m util.ActivityManager) {
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
