package orb

import "github.com/disgoorg/disgo/events"

func listeners(orb *Orb) *events.ListenerAdapter {
	return &events.ListenerAdapter{
		OnReady: func(client *events.Ready) {
			orb.Logger.Info("logged in", "tag", client.User.Tag())
		},
		OnInteraction: func(event *events.InteractionCreate) {
			orb.OnInteraction(event)
		},
	}
}
