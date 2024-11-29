package util

import "net/http"

func Cookie_username_token(r *http.Request) (string, string) {
	var username string
	var token string
	for _, cookie := range r.Cookies() {
		if cookie.Name == "username" {
			username = cookie.Value
		}
		if cookie.Name == "token" {
			token = cookie.Value
		}
	}
	return username, token
}
