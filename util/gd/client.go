package gd

import (
	"fmt"
	"net/http"
	"strings"
)

var GDClient *Client = &Client{}

const (
	DATABASEURL = ""

	// Secrets
	COMMON_KEY = ""
)


type Client struct {
	http.Client
}

func (gd *Client) Request(e *Endpoint, v any) (*http.Response, error) {
	req, err := http.NewRequest(
		e.Method,
		fmt.Sprintf("%s/%s",
			DATABASEURL,
			e.Route,
		),
		strings.NewReader(structToURLValues(v).Encode()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return gd.Do(req)
}

