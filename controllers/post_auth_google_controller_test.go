package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"perisie.com/koran/beans"
	"perisie.com/koran/routes"
	"perisie.com/koran/services"
)

func TestPostAuthGoogleControllerBadRequest(t *testing.T) {
	r, w, _, _, _, _ := routes.SetupTestRoutes(t)

	req := httptest.NewRequest(http.MethodPost, "/auth/google", bytes.NewBuffer([]byte("")))
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Result().StatusCode)
}

func TestPostAuthGoogleControllerNoAuthCode(t *testing.T) {
	r, w, _, _, _, _ := routes.SetupTestRoutes(t)

	req := httptest.NewRequest(http.MethodPost, "/auth/google", bytes.NewBuffer([]byte("{}")))
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Result().StatusCode)

	response, err := io.ReadAll(w.Result().Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"error":"missing \"auth_code\" in body"}`, string(response))
}

func TestPostAuthGoogleControllerBadAuthCode(t *testing.T) {
	r, w, _, _, googleAuthServiceMock, _ := routes.SetupTestRoutes(t)

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

func TestPostAuthGoogleController(t *testing.T) {
	faker := faker.New()
	r, w, _, userManagerMock, googleAuthServiceMock, _ := routes.SetupTestRoutes(t)

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

func Test_PostAuthGoogleController_CreateUserError(t *testing.T) {
	faker := faker.New()
	r, w, _, userManagerMock, googleAuthServiceMock, _ := routes.SetupTestRoutes(t)

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
		Return(nil, errors.New("error"))

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

	assert.Equal(t, 500, w.Result().StatusCode)
}
