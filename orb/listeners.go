package orb

import "github.com/disgoorg/disgo/events"

func listeners(orb *Orb) *events.ListenerAdapter {
	return &events.ListenerAdapter{}
}