package orb

import (
	"context"
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
		Logger: logger,
	}
}

type Orb struct {
	bot.Client
	util.CommandManager
	Config *util.Config
	Logger *log.Logger
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
		),
	)
	if err != nil {
		o.Logger.Fatal("client error", "err", err)
	}
	o.StartNLock()
}

func (o *Orb) SetCommandManager(m util.CommandManager) {
	o.CommandManager = m
}

func (o *Orb) StartNLock() {
	ctx, c := context.WithTimeout(context.Background(), time.Second*10)
	defer c()
	defer func() {
		o.Close(ctx)
		o.Logger.Info("client closed, bye...")
	}()

	err := o.OpenGateway(ctx)
	if err != nil {
		o.Logger.Fatal("gateway open error", "err", err)
	}

	o.Logger.Debug("Bot startup finished, runtime locked")
	k := make(chan os.Signal, 1)
	signal.Notify(k, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-k
}
