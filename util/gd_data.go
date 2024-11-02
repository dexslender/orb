package util

import (
	"reflect"
	"strconv"
	"strings"
)

func DecodeData[T any](data string) *T {
	spl := strings.Split(data, "#")
	RobFormat := strings.Split(spl[0], ":")

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
	for i := 0; i < t.NumField(); i++ {
		if tag, ok := t.Field(i).Tag.Lookup("rob"); ok {
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
	UserName      string `rob:"1"`
	UserID        int    `rob:"2"`
	Stars         int    `rob:"3"`
	Demons        int    `rob:"4"`
	Ranking       int    `rob:"6"`
	CreatorPoints int    `rob:"8"`
	IconID        int    `rob:"9"`
	Color         int    `rob:"10"`
	Color2        int    `rob:"11"`
	SecretCoins   int    `rob:"13"`
	IconType      int    `rob:"14"`
	Special       int    `rob:"15"`
	AccountID     int    `rob:"16"`
	UserCoins     int    `rob:"17"`
	Moons         int    `rob:"52"`
}

type User struct {
	PartialUser `rob:"-"`
	AccountHighlight    int    `rob:"7"`
	MessageState        int    `rob:"18"` // 0: All, 1: Only friends, 2: None
	FriendsState        int    `rob:"19"` // 0: All, 1: None
	YouTube             string `rob:"20"`
	AccIcon             int    `rob:"21"`
	AccShip             int    `rob:"22"`
	AccBall             int    `rob:"23"`
	AccBird             int    `rob:"24"`
	AccDart             int    `rob:"25"` // Wave
	AccRobot            int    `rob:"26"`
	AccStreak           int    `rob:"27"`
	AccGlow             int    `rob:"28"`
	IsRegistered        int    `rob:"29"`
	GlobalRank          int    `rob:"30"`
	FriendState         int    `rob:"31"` // 0: None, 1: already is friend, 3: send request to target, but target haven't accept, 4: target send request, but haven't accept
	Messages            int    `rob:"38"`
	FriendRequests      int    `rob:"39"`
	NewFriends          int    `rob:"40"`
	NewFriendRequest    bool   `rob:"41"`
	Age                 string `rob:"42"`
	AccSpider           int    `rob:"43"`
	Twitter             string `rob:"44"`
	Twitch              string `rob:"45"`
	Diamonds            int    `rob:"46"`
	AccExplosion        int    `rob:"48"`
	ModLevel            int    `rob:"49"` // 0: None, 1: Normal Mod(yellow), 2: Elder Mod(orange)
	CommentHistoryState int    `rob:"50"` // 0: All, 1: Only friends, 2: None
	Color3              int    `rob:"51"`
	AccSwing            int    `rob:"53"`
	AccJetpack          int    `rob:"54"`
	Demonsf             string `rob:"55"` // format {easy},{medium},{hard}.{insane},{extreme},{easyPlatformer},{mediumPlatformer},{hardPlatformer},{insanePlatformer},{extremePlatformer},{weekly},{gauntlet}
	ClassicLevels       string `rob:"56"` // format {auto},{easy},{normal},{hard},{harder},{insane},{daily},{gauntlet}
	PlatformerLevels    string `rob:"57"` // {auto},{easy},{normal},{hard},{harder},{insane}
}
