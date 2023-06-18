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

func Test_PatchUserPointerAdvanceCtrl_401(t *testing.T) {
	r, w, _, _, _, _ := routes.SetupTestRoutes(t)

	req, err := http.NewRequest(http.MethodPatch, "/user/pointer/advance", bytes.NewBuffer([]byte(``)))
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}

func Test_PatchUserPointerAdvanceCtrl_400(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	req, err := http.NewRequest(http.MethodPatch, "/user/pointer/advance", bytes.NewBuffer([]byte(``)))
	req.Header.Add("x-access-token", "token")
	assert.Nil(t, err)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq("token")).
		Return(nil, errors.New("not found"))

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}

func Test_PatchUserPointerAdvanceCtrl_501(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq("token")).
		Return(&beans.User{
			Email: "amir.ariffin@google.com",
		}, nil)

	userManagerMock.EXPECT().
		AdvanceUserCurrentPointer(gomock.Eq("amir.ariffin@google.com"), gomock.Eq("token")).
		Return("", errors.New("error"))

	req, err := http.NewRequest(http.MethodPatch, "/user/pointer/advance", bytes.NewBuffer([]byte(`{"email":"amir.ariffin@google.com"}`)))
	req.Header.Add("x-access-token", "token")
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Result().StatusCode)
}

func Test_PatchUserPointerAdvanceCtrl_200(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq("token")).
		Return(&beans.User{
			Email: "amir.ariffin@google.com",
		}, nil)

	userManagerMock.EXPECT().
		AdvanceUserCurrentPointer(gomock.Eq("amir.ariffin@google.com"), gomock.Eq("token")).
		Return("", nil)

	req, err := http.NewRequest(http.MethodPatch, "/user/pointer/advance", bytes.NewBuffer([]byte(`{"email":"amir.ariffin@google.com"}`)))
	req.Header.Add("x-access-token", "token")
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
}
