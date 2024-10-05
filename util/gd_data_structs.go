package util

type (
	UserData struct {
		/*1	*/ userName string /*The name of player*/
		/*2	*/ userID int /*The ID of player*/
		/*3	*/ stars int /*The count of stars player have*/
		/*4	*/ demons int /*The count of demons player have*/
		/*6	*/ ranking int /*the global leaderboard position of the player*/
		/*7	*/ accountHighlight int /*The accountID of the player. Is used for highlighting the player on the leaderboards*/
		/*8	*/ creatorpoints int /*The count of creatorpoints player have*/
		/*9	*/ iconID int /*maybe... link*/
		/*10*/ color int /*First color of the player use*/
		/*11*/ color2 int /*Second color of the player use*/
		/*13*/ secretCoins int /*The count of coins player have*/
		/*14*/ iconType int /*The iconType of the player use*/
		/*15*/ special int /*The special number of the player use*/
		/*16*/ accountID int /*The accountid of this player*/
		/*17*/ usercoins int /*The count of usercoins player have*/
		/*18*/ messageState int /*0: All, 1: Only friends, 2: None*/
		/*19*/ friendsState int /*0: All, 1: None*/
		/*20*/ youTube string /*The youtubeurl of player*/
		/*21*/ accIcon int /*The icon number of the player use*/
		/*22*/ accShip int /*The ship number of the player use*/
		/*23*/ accBall int /*The ball number of the player use*/
		/*24*/ accBird int /*The bird number of the player use*/
		/*25*/ accDart int /*The dart(wave) number of the player use*/
		/*26*/ accRobot int /*The robot number of the player use*/
		/*27*/ accStreak int /*The streak of the user*/
		/*28*/ accGlow int /*The glow number of the player use*/
		/*29*/ isRegistered int /*if an account is registered or not*/
		/*30*/ globalRank int /*The global rank of this player*/
		/*31*/ friendstate int /*0: None, 1: already is friend, 3: send request to target, but target haven't accept, 4: target send request, but haven't accept*/
		/*38*/ messages int /*How many new messages the user has (shown in-game as a notification)*/
		/*39*/ friendRequests int /*How many new friend requests the user has (shown in-game as a notificaiton)*/
		/*40*/ newFriends int /*How many new Friends the user has (shown in-game as a notificaiton)*/
		/*41*/ NewFriendRequest bool /*appears on userlist endpoint to show if the friend request is new*/
		/*42*/ age string /*the time since you submitted a levelScore*/
		/*43*/ accSpider int /*The spider number of the player use*/
		/*44*/ twitter string /*The twitter of player*/
		/*45*/ twitch string /*The twitch of player*/
		/*46*/ diamonds int /*The count of diamonds player have*/
		/*48*/ accExplosion int /*The explosion number of the player use*/
		/*49*/ modlevel int /*0: None, 1: Normal Mod(yellow), 2: Elder Mod(orange)*/
		/*50*/ commentHistoryState int /*0: All, 1: Only friends, 2: None*/
		/*51*/ color3 int /*The ID of the player's glow color*/
		/*52*/ moons int /*The amount of moons the player has*/
		/*53*/ accSwing int /*The player's swing*/
		/*54*/ accJetpack int /*The player's jetpack*/
		/*55*/ demonsf string /*"demons" Breakdown of the player's demons, in the format {easy},{medium},{hard}.{insane},{extreme},{easyPlatformer},{mediumPlatformer},{hardPlatformer},{insanePlatformer},{extremePlatformer},{weekly},{gauntlet}*/
		/*56*/ classicLevels string /*Breakdown of the player's classic mode non-demons, in the format {auto},{easy},{normal},{hard},{harder},{insane},{daily},{gauntlet}*/
		/*57*/ platformerLevels string /*Breakdown of the player's platformer mode non-demons, in the format {auto},{easy},{normal},{hard},{harder},{insane}*/
	}
)