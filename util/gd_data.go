package util

import (
	"reflect"
	"strconv"
	"strings"
)

func DecodeGDData[T any](data string) *T {
	spl := strings.Split(data, "#")
	RobFormat := strings.Split(spl[0], ":")
	if len(RobFormat) <= 1 { return nil }

	gd := make(map[string]string)
	for i, key := range RobFormat {
		if i&1 == 0 {
			gd[key] = RobFormat[i+1]
		}
	}

	var target T
	tS := reflect.TypeOf(target)
	if tS.Kind() != reflect.Struct {
		return nil
	}
	vS := reflect.ValueOf(&target)
	updateStruct(gd, tS, vS.Elem())
	return &target
}

func updateStruct(data map[string]string, t reflect.Type, v reflect.Value) {
	for i := range t.NumField() {
		if tag, ok := t.Field(i).Tag.Lookup("prop"); ok {
			field := v.Field(i)
			if field.CanSet() {
				robvalue := data[tag]
				switch field.Kind() {
				case reflect.String:
					field.SetString(robvalue)
				case reflect.Int:
					vi, _ := strconv.Atoi(robvalue)
					field.Set(reflect.ValueOf(vi))
				case reflect.Bool:
					vb, _ := strconv.ParseBool(robvalue)
					field.SetBool(vb)
				case reflect.Struct:
					tfield := field.Type()
					updateStruct(data, tfield, field)
				}
			}
		}
	}
}

type PartialUser struct {
	UserName      string `prop:"1"`
	UserID        int    `prop:"2"`
	Stars         int    `prop:"3"`
	Demons        int    `prop:"4"`
	Ranking       int    `prop:"6"`
	CreatorPoints int    `prop:"8"`
	IconID        int    `prop:"9"`
	Color         int    `prop:"10"`
	Color2        int    `prop:"11"`
	SecretCoins   int    `prop:"13"`
	IconType      int    `prop:"14"`
	Special       int    `prop:"15"`
	AccountID     int    `prop:"16"`
	UserCoins     int    `prop:"17"`
	Moons         int    `prop:"52"`
}

type User struct {
	PartialUser `prop:"-"`
	AccountHighlight    int    `prop:"7"`
	MessageState        int    `prop:"18"` // 0: All, 1: Only friends, 2: None
	FriendsState        int    `prop:"19"` // 0: All, 1: None
	YouTube             string `prop:"20"`
	AccIcon             int    `prop:"21"`
	AccShip             int    `prop:"22"`
	AccBall             int    `prop:"23"`
	AccBird             int    `prop:"24"`
	AccDart             int    `prop:"25"` // Wave
	AccRobot            int    `prop:"26"`
	AccStreak           int    `prop:"27"`
	AccGlow             int    `prop:"28"`
	IsRegistered        int    `prop:"29"`
	GlobalRank          int    `prop:"30"`
	FriendState         int    `prop:"31"` // 0: None, 1: already is friend, 3: send request to target, but target haven't accept, 4: target send request, but haven't accept
	Messages            int    `prop:"38"`
	FriendRequests      int    `prop:"39"`
	NewFriends          int    `prop:"40"`
	NewFriendRequest    bool   `prop:"41"`
	Age                 string `prop:"42"`
	AccSpider           int    `prop:"43"`
	Twitter             string `prop:"44"`
	Twitch              string `prop:"45"`
	Diamonds            int    `prop:"46"`
	AccExplosion        int    `prop:"48"`
	ModLevel            int    `prop:"49"` // 0: None, 1: Normal Mod(yellow), 2: Elder Mod(orange)
	CommentHistoryState int    `prop:"50"` // 0: All, 1: Only friends, 2: None
	Color3              int    `prop:"51"`
	AccSwing            int    `prop:"53"`
	AccJetpack          int    `prop:"54"`
	Demonsf             string `prop:"55"` // format {easy},{medium},{hard}.{insane},{extreme},{easyPlatformer},{mediumPlatformer},{hardPlatformer},{insanePlatformer},{extremePlatformer},{weekly},{gauntlet}
	ClassicLevels       string `prop:"56"` // format {auto},{easy},{normal},{hard},{harder},{insane},{daily},{gauntlet}
	PlatformerLevels    string `prop:"57"` // {auto},{easy},{normal},{hard},{harder},{insane}
}

type Level struct {}
