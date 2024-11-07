package controllers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"perisie.com/koran/managers"
)

func GetRootController(tmpl *template.Template, quranManager managers.QuranManager) func(c *gin.Context) {
	return func(c *gin.Context) {
		surah_infos, _ := quranManager.GetSurahInfos()
		_ = tmpl.ExecuteTemplate(c.Writer, "page_root.html", map[string]interface{}{
			"surah_infos": surah_infos,
		})
	}
}
