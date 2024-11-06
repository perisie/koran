package routes

import (
	"github.com/arikama/koran-backend/controllers"
	"github.com/arikama/koran-backend/favorite"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/middleware"
	"github.com/arikama/koran-backend/services"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Routes(
	r *gin.Engine,
	tmpl *template.Template,
	quranManager managers.QuranManager,
	googleAuthService services.GoogleAuthService,
	userManager managers.UserManager,
	favManager favorite.FavManager,
) {
	public := r.Group("/")
	public.GET("/", controllers.GetRootController(tmpl, quranManager))
	{
		public.GET("/surah/:surah_id", controllers.GetSurahController(quranManager))
		public.GET("/surah/:surah_id/verse/:verse_id", controllers.GetSurahVerseController(quranManager))

		public.POST("/auth/google", controllers.PostAuthGoogleController(googleAuthService, userManager))
	}

	authorized := r.Group("/")
	authorized.Use(middleware.UserAuth(userManager))
	{
		authorized.GET("/fav", favorite.GetFavCtrl(favManager, userManager))

		authorized.POST("/fav", favorite.PostFavCtrl(favManager, userManager))
		authorized.POST("/fav/remove", favorite.PostFavRemoveCtrl(favManager, userManager))
		authorized.POST("/user/pointer", controllers.PostUserPointerCtrl(userManager))

		authorized.PATCH("/user/pointer/advance", controllers.PatchUserPointerAdvanceCtrl(userManager))
		authorized.PATCH("/user/pointer/reverse", controllers.PatchUserPointerReverseCtrl(userManager))
	}
}
