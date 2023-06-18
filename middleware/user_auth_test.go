package middleware_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/arikama/koran-backend/routes"
	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func Test_user_auth(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	req, err := http.NewRequest(http.MethodGet, "/fav", nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)

	faker := faker.New()
	token := faker.Internet().Password()

	req.Header.Add("x-access-token", token)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, errors.New("not found"))

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}
