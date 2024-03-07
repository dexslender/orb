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

func New(l *log.Logger, c util.Config) *Orb {
	return &Orb{
		Log:    l,
		Config: c,
	}
}

type Orb struct {
	bot.Client
	Config    util.Config
	Log       *log.Logger
	UPresence *util.PresenceUpdater
}

func (o *Orb) Setup() {
	var err error

	o.UPresence = util.NewPresenceUpdater(o.Log, o.Config)

	o.Client, err = disgo.New(o.Config.Bot.Token,
		bot.WithGatewayConfigOpts(
			func(c *gateway.Config) {
				if o.Config.Bot.MobileOs {
					c.Browser = "Discord Android"
					o.UPresence.Setup(c)
				}
			},
			gateway.WithIntents(
				gateway.IntentsNonPrivileged,
			),
		),
		bot.WithLogger(slog.New(o.Log)),
		bot.WithEventListeners(listeners(o)),
	)
	if err != nil {
		o.Log.Fatal("failed to create client", "err", err)
	}
}

func (o *Orb) StartNLock() {
	ctx, c := context.WithTimeout(context.Background(), time.Second*10)
	defer c()

	defer func() {
		o.Close(ctx)
		o.Log.Info("client closed, turning off\n\t(please wait)")
	}()

	if err := o.OpenGateway(ctx); err != nil {
		o.Log.Fatal("Gateway open error: ", err)
	}

	k := make(chan os.Signal, 1)
	signal.Notify(k, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-k
}
