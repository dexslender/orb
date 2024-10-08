package orb

import "github.com/disgoorg/disgo/events"

func listeners(o *Orb) *events.ListenerAdapter {
	return &events.ListenerAdapter{
		OnReady: func(client *events.Ready) {
			o.Log.Info("logged in", "tag", client.User.Tag())
			go o.StartActivityUpdater(o)
		},
		OnInteraction: func(event *events.InteractionCreate) {
			o.OnInteraction(event)
		},
	}
}
