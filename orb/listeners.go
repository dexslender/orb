package orb

import "github.com/disgoorg/disgo/events"

func listeners(o *Orb) *events.ListenerAdapter {
	return &events.ListenerAdapter{
		OnReady: func(r *events.Ready) {
			o.Log.Info("Logged in", "username", r.User.Username)
			o.UPresence.StartUpdater(o)
		},
	}
}
