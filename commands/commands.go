package commands

import "github.com/dexslender/orb/util"

var Commands = []util.Command{
	&Ping{},
}

type base struct {
	util.Command
}

func (c *base) Error() {
	print("someting command error")
}

// func (c *command) Init() {}
// func (c *command) Run() {}
