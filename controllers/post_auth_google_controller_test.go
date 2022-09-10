package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arikama/koran-backend/beans"
	"github.com/arikama/koran-backend/controllers"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestPostAuthGoogleControllerBadRequest(t *testing.T) {
	r, w, _, _ := setupPostAuthGoogleController(t)

	req := httptest.NewRequest(http.MethodPost, "/auth/google", bytes.NewBuffer([]byte("")))
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Result().StatusCode)
}

func TestPostAuthGoogleControllerNoAuthCode(t *testing.T) {
	r, w, _, _ := setupPostAuthGoogleController(t)

	req := httptest.NewRequest(http.MethodPost, "/auth/google", bytes.NewBuffer([]byte("{}")))
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Result().StatusCode)

	response, err := io.ReadAll(w.Result().Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"error":"missing \"auth_code\" in body"}`, string(response))
}

func TestPostAuthGoogleControllerBadAuthCode(t *testing.T) {
	r, w, googleAuthServiceMock, _ := setupPostAuthGoogleController(t)

	googleAuthServiceMock.EXPECT().
		AuthUserCode(gomock.Eq("bad")).
		Return(nil, errors.New("bad auth code"))

	req := httptest.NewRequest(http.MethodPost, "/auth/google", bytes.NewBuffer([]byte(`{"auth_code":"bad"}`)))
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Result().StatusCode)

	response, err := io.ReadAll(w.Result().Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"error":"bad auth code"}`, string(response))
}

func XTestPostAuthGoogleController(t *testing.T) {
	faker := faker.New()
	r, w, googleAuthServiceMock, userManagerMock := setupPostAuthGoogleController(t)

	googleUser := &services.GoogleUser{
		Email: faker.Internet().Email(),
		Token: faker.Internet().Password(),
	}
	authCode := faker.Internet().Password()
	googleAuthServiceMock.EXPECT().
		AuthUserCode(gomock.Eq(authCode)).
		Return((googleUser), nil)

	userManagerMock.EXPECT().
		CreateUser(gomock.Eq(googleUser.Email), gomock.Eq(googleUser.Token)).
		Return(&beans.User{
			Email: googleUser.Email,
			Token: googleUser.Token,
		}, nil)

	type JsonBody struct {
		AuthCode string `json:"auth_code"`
	}
	jsonBody := JsonBody{
		AuthCode: authCode,
	}
	buf, err := json.Marshal(jsonBody)
	assert.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/auth/google", bytes.NewBuffer(buf))
	r.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Result().Body)
	assert.Nil(t, err)
}

func setupPostAuthGoogleController(t *testing.T) (
	*gin.Engine,
	*httptest.ResponseRecorder,
	*services.GoogleAuthServiceMock,
	*managers.UserManagerMock,
) {
	ctrl := gomock.NewController(t)
	googleAuthServiceMock := services.NewGoogleAuthServiceMock(ctrl)
	userManagerMock := managers.NewUserManagerMock(ctrl)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.POST("/auth/google", controllers.PostAuthGoogleController(googleAuthServiceMock, nil))

	return r, w, googleAuthServiceMock, userManagerMock
}
