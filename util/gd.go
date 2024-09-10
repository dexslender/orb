package util

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

var GDClient *gdClient = &gdClient{}

const (
	DATABASEURL = ""

	// Secrets
	COMMON_KEY = ""
)

type (
	Endpoint struct { Method, Route string }
	gdClient struct { http.Client }

	//---
	UsersParams struct { 
		Secret string `query:"secret"`
		Query string `query:"str"`
	}
	
	//---
	UserData struct {}
)

func NewEndpoint(method, route string) *Endpoint {
	return &Endpoint{method, route}
}

var (
	Users = NewEndpoint(http.MethodPost, "getGJUsers20.php")
	Scores = NewEndpoint(http.MethodPost, "getGJScores20.php")
	UserInfo = NewEndpoint(http.MethodPost, "getGJUserInfo20.php")

	// Levels
	Dayly = NewEndpoint(http.MethodPost, "getGJDailyLevel.php")
)

func (gd *gdClient) Request(e *Endpoint, v any) (*http.Response, error)  {
	req, err := http.NewRequest(
		e.Method,
		fmt.Sprintf("%s/%s",
			DATABASEURL,
			e.Route,
		),
		strings.NewReader(structToURLValues(v).Encode()),
	)
	if err != nil { return nil, err }
	req.Header.Set("User-Agent", "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return gd.Do(req)
	
}


func structToURLValues(item interface{}) url.Values {
	res := url.Values{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("query")
		field := reflectValue.Field(i).String()
		if tag != "" && tag != "-" {
			res.Add(tag, field)
		}
	}
	return res
}
