package main

import (
	"github.com/arikama/koran-backend/controllers"
	"github.com/arikama/koran-backend/managers"
	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine, quranManager managers.QuranManager) {
	r.GET("/", controllers.GetRootController())
	r.GET("/surah/:surah_id/verse/:verse_id", controllers.GetSurahVerseController(quranManager))
	r.POST("/auth/google", controllers.PostAuthGoogleController())
}
