package commands

import (
	"github.com/dexslender/orb/commands/about"
	"github.com/dexslender/orb/orb"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	// "github.com/disgoorg/disgo/handler/middleware"
)

var List = []discord.ApplicationCommandCreate{
	ping,
	purge,
	setupCmd,
	gdCmd,
	aboutCmd,
}

func Setup(o *orb.Orb, router handler.Router) {
	// router.Use(middleware.Logger)
	router.Command("/gd/{sub}", runGD)
	router.Command("/setup/{sub}", runSetup)

	router.Group(func(r handler.Router) {
		r.Command("/purge", runPurge)
		r.Command("/ping", runPing)
	})

	router.Group(func(r handler.Router) {
		r.Command("/about/{sub}", runAbout(o))
		r.ButtonComponent("/about/guild/{guildId}", about.UpdateWithGuildStats)
	})

	router.ButtonComponent("/ping/refresh", runPingRefresh)
}
