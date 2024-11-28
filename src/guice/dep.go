package guice

import (
	"html/template"
	"net/http"
	"perisie.com/koran/src/quran"
	"perisie.com/koran/src/user"
)

type Dep struct {
	Tmpl       *template.Template
	Mngr_user  user.Mngr
	Mngr_quran quran.Mngr
	Fs         http.Handler
}

func Dep_new(dir_tmpl, dir_quran, dir_static string) *Dep {
	tmpl := template.Must(template.ParseGlob(dir_tmpl + "/*.html"))
	mngr_user := user.Mngr_impl_new()
	mngr_quran, err := quran.Mngr_impl_new(dir_quran)
	if err != nil {
		panic(err.Error())
	}
	fs := http.FileServer(http.Dir(dir_static))

	return &Dep{
		Tmpl:       tmpl,
		Mngr_user:  mngr_user,
		Mngr_quran: mngr_quran,
		Fs:         fs,
	}
}

func Dep_test(dir_tmpl, dir_quran, dir_static string) *Dep {
	dependency := Dep_new(dir_tmpl, dir_quran, dir_static)
	dependency.Mngr_user = user.Mngr_impl_fake()
	return dependency
}
