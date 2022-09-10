package routes

import (
	"net/http/httptest"
	"testing"

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
) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	ctrl := gomock.NewController(t)
	quranManagerImpl, _ := managers.NewQuranManagerImpl("./qurancsv")
	userManagerMock := managers.NewUserManagerMock(ctrl)
	googleAuthServiceMock := services.NewGoogleAuthServiceMock(ctrl)

	Routes(r, quranManagerImpl, googleAuthServiceMock, userManagerMock)

	return r, w, quranManagerImpl, userManagerMock, googleAuthServiceMock
}
