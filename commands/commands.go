package commands

import (
	"github.com/dexslender/orb/orb"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var List = []discord.ApplicationCommandCreate{
	ping,
	purge,
	setupCmd,
	gdCmd,
	aboutCmd,
}

func Setup(o *orb.Orb, router handler.Router) {
	router.Command("/about/{sub}", runAbout(o))
	router.Command("/gd/{sub}", runGD)
	router.Command("/setup/{sub}", runSetup)

	router.Group(func(r handler.Router) {
		r.Command("/purge", runPurge)
		r.Command("/ping", runPing)
	})


	router.ButtonComponent("/ping/refresh", runPingRefresh)
}
