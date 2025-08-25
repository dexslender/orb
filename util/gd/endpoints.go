package gd

import "net/http"

type Endpoint struct{ Method, Route string }

func NewEndpoint(method, route string) *Endpoint {
	return &Endpoint{method, route}
}


var (
	Users    = NewEndpoint(http.MethodPost, "getGJUsers20.php")
	Scores   = NewEndpoint(http.MethodPost, "getGJScores20.php")
	UserInfo = NewEndpoint(http.MethodPost, "getGJUserInfo20.php")

	// Levels
	Daily         = NewEndpoint(http.MethodPost, "getGJDailyLevel.php")
	DownloadLevel = NewEndpoint(http.MethodPost, "downloadGJLevel22.php")
)

