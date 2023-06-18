package controllers_test

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/arikama/koran-backend/beans"
	"github.com/arikama/koran-backend/routes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_PostUserPointerCtrl_401(t *testing.T) {
	r, w, _, _, _, _ := routes.SetupTestRoutes(t)

	req, err := http.NewRequest(http.MethodPost, "/user/pointer", bytes.NewBuffer([]byte(``)))
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}

func Test_PostUserPointerCtrl_400(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	req, err := http.NewRequest(http.MethodPost, "/user/pointer", bytes.NewBuffer([]byte(``)))
	req.Header.Add("x-access-token", "token")
	assert.Nil(t, err)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq("token")).
		Return(nil, errors.New("error"))

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}

func Test_PostUserPointerCtrl_404(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq("token")).
		Return(nil, errors.New("error"))

	req, err := http.NewRequest(http.MethodPost, "/user/pointer", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Add("x-access-token", "token")
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}

func Test_PostUserPointerCtrl_400_EmailMismatch(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq("token")).
		Return(&beans.User{
			Email: "amir.ariffin@google.com",
		}, nil)
	userManagerMock.EXPECT().
		GetUser(gomock.Eq("token")).
		Return(&beans.User{
			Email: "amir.ariffin@google.com",
		}, nil)

	req, err := http.NewRequest(http.MethodPost, "/user/pointer", bytes.NewBuffer([]byte(`{"email":"amir.ariffin@gmail.com"}`)))
	req.Header.Add("x-access-token", "token")
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}

func Test_PostUserPointerCtrl_200(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq("token")).
		Return(&beans.User{
			Email: "amir.ariffin@google.com",
		}, nil)
	userManagerMock.EXPECT().
		GetUser(gomock.Eq("token")).
		Return(&beans.User{
			Email: "amir.ariffin@google.com",
		}, nil)

	req, err := http.NewRequest(http.MethodPost, "/user/pointer", bytes.NewBuffer([]byte(`{"email":"amir.ariffin@google.com"}`)))
	req.Header.Add("x-access-token", "token")
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
}
