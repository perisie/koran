package routes

import (
	"github.com/arikama/koran-backend/controllers"
	"github.com/arikama/koran-backend/favorite"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/middleware"
	"github.com/arikama/koran-backend/services"
	"github.com/gin-gonic/gin"
)

func Routes(
	r *gin.Engine,
	quranManager managers.QuranManager,
	googleAuthService services.GoogleAuthService,
	userManager managers.UserManager,
	favManager favorite.FavManager,
) {
	public := r.Group("/")
	public.GET("/", controllers.GetRootController(quranManager))
	{
		public.GET("/surah/:surah_id", controllers.GetSurahController(quranManager))
		public.GET("/surah/:surah_id/verse/:verse_id", controllers.GetSurahVerseController(quranManager))

		public.POST("/auth/google", controllers.PostAuthGoogleController(googleAuthService, userManager))
	}

	authorized := r.Group("/")
	authorized.Use(middleware.UserAccessMiddleware(userManager))
	{
		authorized.GET("/fav", favorite.GetFavCtrl(favManager, userManager))

		authorized.POST("/fav", favorite.PostFavCtrl(favManager, userManager))
		authorized.POST("/fav/remove", favorite.PostFavRemoveCtrl(favManager, userManager))
		authorized.POST("/user/pointer", controllers.PostUserPointerCtrl(userManager))

		authorized.PATCH("/user/pointer/advance", controllers.PatchUserPointerAdvanceCtrl(userManager))
	}
}
