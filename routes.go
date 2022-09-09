package main

import (
	"github.com/arikama/koran-backend/controllers"
	"github.com/arikama/koran-backend/managers"
	"github.com/gin-gonic/gin"
)

func routes(
	r *gin.Engine,
	quranManager managers.QuranManager,
	googleAuthManager managers.GoogleAuthManager,
) {
	r.GET("/", controllers.GetRootController(quranManager))
	r.GET("/surah/:surah_id", controllers.GetSurahController(quranManager))
	r.GET("/surah/:surah_id/verse/:verse_id", controllers.GetSurahVerseController(quranManager))
	r.GET("/user/pointer", controllers.GetUserPointerCtrl())
	r.POST("/auth/google", controllers.PostAuthGoogleController(googleAuthManager))
	r.PATCH("/user/pointer/advance", controllers.PatchUserPointerAdvanceCtrl())
}
