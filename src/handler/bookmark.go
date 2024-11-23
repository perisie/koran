package handler

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/quran"
	"perisie.com/koran/src/user"
	"perisie.com/koran/src/util"
)

func Bookmark(tmpl *template.Template, mngr_user user.Mngr, mngr_quran quran.Mngr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, token := util.Cookie_username_token(r.Cookies())
		user, _ := mngr_user.Get(username)
		if !user.Ok() || user.Password != token {
			tmpl.ExecuteTemplate(w, "page_login.html", nil)
			return
		}
		verse, err := mngr_quran.Get_verse(user.Surah, user.Verse)
		if err != nil {
			util.Redirect_error_page(w, http.StatusInternalServerError, err)
			return
		}
		switch r.Method {
		case http.MethodPatch:
			{
				q := r.URL.Query()
				move := q.Get("move")
				direction := 0
				if move == "next" {
					direction += 1
				}
				if move == "prev" {
					direction -= 1
				}
				surah_id, verse_id, _ := util.Move_surah_verse(user.Surah, user.Verse, direction)
				_ = mngr_user.Update_surah_verse(user.Username, surah_id, verse_id)
				verse, err := mngr_quran.Get_verse(surah_id, verse_id)
				if err != nil {
					util.Redirect_error_page(w, http.StatusInternalServerError, err)
					return
				}
				tmpl.ExecuteTemplate(w, "comp_verse.html", verse)
			}
		default:
			{
				tmpl.ExecuteTemplate(w, "page_bookmark.html", map[string]interface{}{
					"user":  user,
					"verse": verse,
				})
			}
		}
	}
}
