package handler

import (
	"net/http"
	"perisie.com/koran/src/guice"
)

func Mux(dep *guice.Dep) *http.ServeMux {
	tmpl := dep.Tmpl
	mngr_user := dep.Mngr_user
	mngr_quran := dep.Mngr_quran
	fs := dep.Fs

	mux := http.NewServeMux()

	mux.HandleFunc("/", Home(dep))
	mux.HandleFunc("/surah/{surah_id}", Surah(tmpl, mngr_user, mngr_quran))
	mux.HandleFunc("/surah/{surah_id}/verse/{verse_id}", Verse(tmpl, mngr_user, mngr_quran))

	mux.HandleFunc("/bookmark", Bookmark(tmpl, mngr_user, mngr_quran))
	mux.HandleFunc("/setting", Setting(tmpl, mngr_user))

	mux.HandleFunc("/login", Login(tmpl, mngr_user))
	mux.HandleFunc("/signup", Signup(tmpl, mngr_user))
	mux.HandleFunc("/user", User(tmpl, mngr_user))

	mux.HandleFunc("/error", Error(tmpl))
	mux.HandleFunc("/static/", Static(fs))

	return mux
}
