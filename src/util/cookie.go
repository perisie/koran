package util

import "net/http"

func Cookie_username_token(cookies []*http.Cookie) (string, string) {
	var username string
	var token string
	for _, cookie := range cookies {
		if cookie.Name == "username" {
			username = cookie.Value
		}
		if cookie.Name == "token" {
			token = cookie.Value
		}
	}
	return username, token
}
