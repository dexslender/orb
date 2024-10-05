package util

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/dexslender/orb/orb"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
)

type Amanager struct {
	Logger  *log.Logger
	Config  *orb.Config
	current int
}

func (a *Amanager) SetupActivity() gateway.ConfigOpt {
	return func(c *gateway.Config) {
		if a.Config.ActivityManager.Enabled && a.size() >= 1 {
			a.Logger.Debug("configuring presence in gateway", "current", a.current)
			c.Presence = a.next()
		} else {
			a.Logger.Info("activity disabled")
		}
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

func (a *Amanager) resolveActivity(act orb.Activity) *gateway.MessageDataPresenceUpdate {
	var (
		activity discord.Activity
		status discord.OnlineStatus
	)
	if act.Status == nil { status = discord.OnlineStatusOnline } else 
	{ status = *act.Status }
	if act.Type == nil { activity.Type = discord.ActivityTypeGame } else
	{ activity.Type = *act.Type }

	switch *act.Type {
	case discord.ActivityTypeCustom:
		activity.Name, activity.State = act.Message, &act.Message
	case discord.ActivityTypeStreaming:
		activity.URL = act.URL
	default:
		activity.Name = act.Message
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
