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

func Dep_new(dir_tmpl, dir_quran, dir_static, dir_data string) *Dep {
	dep := Dep_test(dir_tmpl, dir_quran, dir_static)
	dep.Mngr_user = user.Mngr_impl_new(dir_data)
	return dep
}

func Dep_test(dir_tmpl, dir_quran, dir_static string) *Dep {
	tmpl := template.Must(template.ParseGlob(dir_tmpl + "/*.html"))
	mngr_quran, err := quran.Mngr_impl_new(dir_quran)
	if err != nil {
		panic(err.Error())
	}
	fs := http.FileServer(http.Dir(dir_static))
	return &Dep{
		Tmpl:       tmpl,
		Mngr_user:  user.Mngr_impl_fake(),
		Mngr_quran: mngr_quran,
		Fs:         fs,
	}
}
