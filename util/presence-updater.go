package util

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
)

func NewPresenceUpdater(l *log.Logger, c Config) *PresenceUpdater {
	return &PresenceUpdater{
		config: c,
		logger: l,
	}
}

type PresenceUpdater struct {
	config  Config
	logger  *log.Logger
	current int
}

func (p *PresenceUpdater) Setup(gc *gateway.Config) {
	if p.config.PresenceUpdater.Enabled &&
		len(p.config.PresenceUpdater.Presences) >= 1 {
		gc.Presence = p.next()
	} else {
		p.logger.Info("presence updater disabled")
	}
}

func (p *PresenceUpdater) StartUpdater(c bot.Client) {
	if !p.config.PresenceUpdater.Enabled ||
		len(p.config.PresenceUpdater.Presences) <= 1 {
		return
	}
	ticker := time.NewTicker(p.config.PresenceUpdater.Delay)
	for range ticker.C {
		pi := p.next()
		if pi == nil {
			continue
		}
		p.logger.Debug("updating bot presence", "index", p.current)
		c.Gateway().Presence().Activities = pi.Activities
		c.Gateway().Presence().Status = pi.Status
		if err := c.SetPresence(context.Background()); err != nil {
			p.logger.Error("failed to send presence data to gateway", "err", err)
		}
	}
}

func (p *PresenceUpdater) next() (data *gateway.MessageDataPresenceUpdate) {
	size := len(p.config.PresenceUpdater.Presences)

	pi := resolve_presence(p.config.PresenceUpdater.Presences[p.current])
	if pi != nil {
		data = pi
	} else {
		p.logger.Warn("presence requires 'name' or 'state'", "index", p.current)
	}

	if p.current >= size-1 {
		p.current = 0
	} else {
		p.current++
	}
	return
}

func resolve_presence(pi Presence) *gateway.MessageDataPresenceUpdate {
	var (
		status   discord.OnlineStatus
		activity discord.Activity
	)

	if pi.Name == "" && pi.State == "" {
		return nil
	}

	activity.Name = pi.Name
	if pi.State != "" {
		activity.State = &pi.State
		if activity.Name == "" {
			activity.Name = pi.State
		}
	}

	if pi.Status == "" {
		status = discord.OnlineStatusOnline
	} else {
		status = pi.Status
	}

	switch pi.Type {
	case "watching", "watch":
		activity.Type = discord.ActivityTypeWatching
	case "listening", "listen":
		activity.Type = discord.ActivityTypeListening
	case "game", "playing":
		activity.Type = discord.ActivityTypeGame
	case "competing":
		activity.Type = discord.ActivityTypeCompeting
	case "streaming", "stream":
		activity.Type = discord.ActivityTypeStreaming
	case "custom":
		activity.Type = discord.ActivityTypeCustom
	default:
		if pi.State != "" && pi.Name == "" {
			activity.Type = discord.ActivityTypeCustom
		} else {
			activity.Type = discord.ActivityTypeGame
		}
	}

	if activity.Type == discord.ActivityTypeStreaming && pi.URL != "" {
		activity.URL = &pi.URL
	}

	return &gateway.MessageDataPresenceUpdate{
		Status:     status,
		Activities: []discord.Activity{activity},
	}
}
