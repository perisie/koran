package routes

import (
	"html/template"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"perisie.com/koran/favorite"
	"perisie.com/koran/managers"
	"perisie.com/koran/services"
)

func SetupTestRoutes(t *testing.T) (
	*gin.Engine,
	*httptest.ResponseRecorder,
	*managers.QuranManagerImpl,
	*managers.UserManagerMock,
	*services.GoogleAuthServiceMock,
	*favorite.FavManagerMock,
) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	tmpl, _ := template.ParseGlob("../template/*.html")
	ctrl := gomock.NewController(t)
	quranManagerImpl, _ := managers.NewQuranManagerImpl("./qurancsv")
	userManagerMock := managers.NewUserManagerMock(ctrl)
	googleAuthServiceMock := services.NewGoogleAuthServiceMock(ctrl)
	favManagerMock := favorite.NewFavManagerMock(ctrl)

	Routes(r, tmpl, quranManagerImpl, googleAuthServiceMock, userManagerMock, favManagerMock)

	return r, w, quranManagerImpl, userManagerMock, googleAuthServiceMock, favManagerMock
}
