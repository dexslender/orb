package util

import (
	"cmp"
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
)

type ActivityManager interface {
	SetupActivity(*gateway.Config)
	StartActivityUpdater(bot.Client)
}

var _ ActivityManager = (*Amanager)(nil)

type Amanager struct {
	Logger  *log.Logger
	Config  *Config
	current int
}

func (a *Amanager) SetupActivity(c *gateway.Config) {
	if a.Config.ActivityManager.Enabled && a.size() >= 1 {
		a.Logger.Debug("configuring presence in gateway", "current", a.current)
		c.Presence = a.next()
	} else {
		a.Logger.Info("activity disabled")
	}
}

// Go func (*_*)
func (a *Amanager) StartActivityUpdater(c bot.Client) {
	if !a.Config.ActivityManager.Enabled || a.size() <= 1 {
		return
	}
	ticker := time.NewTicker(a.Config.ActivityManager.Interval)
	for range ticker.C {
		act := a.next()
		if act == nil {
			continue
		}
		a.Logger.Debug("updating presence", "current", a.current)
		c.Gateway().Presence().Activities = act.Activities
		c.Gateway().Presence().Status = act.Status
		if err := c.SetPresence(context.Background()); err != nil {
			a.Logger.Error("failed to send: presence data -> gateway", "err", err)
		}
	}
}

func (a *Amanager) resolveActivity(act Activity) *gateway.MessageDataPresenceUpdate {
	var (
		status   discord.OnlineStatus
		activity discord.Activity
	)

	if act.Name != "" {
		activity.Name = act.Name
	} else if act.State != "" {
		activity.State = &act.State
		activity.Name = act.State
	} else {
		return nil
	}

	status = cmp.Or(act.Status, "online")

	activity.Type = func() discord.ActivityType {
		switch act.Type {
		case "watching", "watch":
			return discord.ActivityTypeWatching
		case "listening", "listen":
			return discord.ActivityTypeListening
		case "game", "playing":
			return discord.ActivityTypeGame
		case "competing":
			return discord.ActivityTypeCompeting
		case "streaming", "stream":
			return discord.ActivityTypeStreaming
		case "custom":
			return discord.ActivityTypeCustom
		default:
			if activity.State != nil && activity.Name != "" {
				return discord.ActivityTypeCustom
			} else {
				return discord.ActivityTypeGame
			}
		}
	}()

	if activity.Type == discord.ActivityTypeStreaming && act.URL != "" {
		activity.URL = &act.URL
	}
	return &gateway.MessageDataPresenceUpdate{
		Status:     status,
		Activities: []discord.Activity{activity},
	}
}

func (a *Amanager) next() (data *gateway.MessageDataPresenceUpdate) {
	act := a.resolveActivity(a.Config.ActivityManager.Activities[a.current])
	if act != nil {
		data = act
	} else {
		a.Logger.Error("activity require 'name' or 'state'", "current", a.current)
	}
	if a.current >= a.size()-1 {
		a.current = 0
	} else {
		a.current += 1
	}
	return
}

func (a *Amanager) size() int {
	return len(a.Config.ActivityManager.Activities)
}
