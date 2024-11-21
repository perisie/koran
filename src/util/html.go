package util

import (
	"html"
	"net/http"
	"strconv"
)

func Redirect_error_page(w http.ResponseWriter, code int, err error) {
	if err != nil {
		msg := html.EscapeString(err.Error())
		w.Header().Add("HX-Redirect", "/error?code="+strconv.Itoa(code)+"&msg="+msg)
	}
}
