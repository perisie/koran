package routes

import (
	"github.com/arikama/koran-backend/controllers"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/services"
	"github.com/gin-gonic/gin"
)

func Routes(
	r *gin.Engine,
	quranManager managers.QuranManager,
	googleAuthService services.GoogleAuthService,
	userManager managers.UserManager,
) {
	r.GET("/", controllers.GetRootController(quranManager))
	r.GET("/surah/:surah_id", controllers.GetSurahController(quranManager))
	r.GET("/surah/:surah_id/verse/:verse_id", controllers.GetSurahVerseController(quranManager))
	r.POST("/user/pointer", controllers.PostUserPointerCtrl(userManager))
	r.POST("/auth/google", controllers.PostAuthGoogleController(googleAuthService, userManager))
	r.PATCH("/user/pointer/advance", controllers.PatchUserPointerAdvanceCtrl(userManager))
}
