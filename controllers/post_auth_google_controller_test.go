package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arikama/koran-backend/controllers"
	"github.com/arikama/koran-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestPostAuthGoogleControllerBadRequest(t *testing.T) {
	r, w, _ := setupPostAuthGoogleController(t)

	req := httptest.NewRequest(http.MethodPost, "/auth/google", bytes.NewBuffer([]byte("")))
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Result().StatusCode)
}

func TestPostAuthGoogleControllerNoAuthCode(t *testing.T) {
	r, w, _ := setupPostAuthGoogleController(t)

	req := httptest.NewRequest(http.MethodPost, "/auth/google", bytes.NewBuffer([]byte("{}")))
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Result().StatusCode)

	response, err := io.ReadAll(w.Result().Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"error":"missing \"auth_code\" in body"}`, string(response))
}

func TestPostAuthGoogleControllerBadAuthCode(t *testing.T) {
	r, w, manager := setupPostAuthGoogleController(t)

	manager.EXPECT().
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
	r, w, googleAuthManagerMock := setupPostAuthGoogleController(t)

	googleUser := &services.GoogleUser{
		Email: "amir.ariffin@google.com",
		Name:  "Amir",
	}
	authCode := faker.Internet().Password()
	googleAuthManagerMock.EXPECT().AuthUserCode(gomock.Eq(authCode)).Return((googleUser), nil)

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

	body, err := io.ReadAll(w.Result().Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"data":{"email":"amir.ariffin@google.com","name":"Amir","token":"","picture":""}}`, string(body))
}

func setupPostAuthGoogleController(t *testing.T) (*gin.Engine, *httptest.ResponseRecorder, *services.GoogleAuthServiceMock) {
	ctrl := gomock.NewController(t)
	googleAuthManagerMock := services.NewGoogleAuthServiceMock(ctrl)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.POST("/auth/google", controllers.PostAuthGoogleController(googleAuthManagerMock))

	return r, w, googleAuthManagerMock
}
