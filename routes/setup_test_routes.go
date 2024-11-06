package routes

import (
	"html/template"
	"net/http/httptest"
	"testing"

	"github.com/arikama/koran-backend/favorite"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
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
