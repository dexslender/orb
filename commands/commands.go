package commands

import (
	"errors"

	"github.com/dexslender/orb/util"
)

var Commands = []util.Command{
	&Ping{},
}

type base struct {
	util.Command
}

func (c *base) Run(*util.CommandContext) error { return errors.New("missing run function :(") }

func (c *base) Error(cctx *util.CommandContext, err error) {
	cctx.Logger.Error("command malfunction",
		"command", cctx.Data.CommandName(),
		"error", err,
	)
}
