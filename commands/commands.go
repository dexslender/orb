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
	router.Group(func(r handler.Router) {
		r.Command("/about/{sub}", runAbout(o))
	})
}
