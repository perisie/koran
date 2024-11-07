package favorite_test

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"perisie.com/koran/beans"
	"perisie.com/koran/favorite"
	"perisie.com/koran/routes"
)

func Test_controller_PostFavRemoveCtrl_success(t *testing.T) {
	r, w, _, userManagerMock, _, favManagerMock := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav/remove", bytes.NewBuffer([]byte(`{"id":1}`)))
	req.Header.Add("x-access-token", token)

	user := beans.User{
		Email: faker.Internet().Email(),
	}

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(&user, nil)

	favManagerMock.EXPECT().
		DeleteFav(gomock.Eq(1)).
		Return(nil)

	favManagerMock.EXPECT().
		GetFavs(gomock.Eq(user.Email)).
		Return([]*favorite.Fav{
			{
				Id:    2,
				Surah: 1,
				Verse: 2,
			},
		}, nil)

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, `{"data":{"favorites":[{"id":2,"surah":1,"verse":2}]}}`, w.Body.String())
}

func Test_controller_PostFavRemoveCtrl_error_getFavs(t *testing.T) {
	r, w, _, userManagerMock, _, favManagerMock := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav/remove", bytes.NewBuffer([]byte(`{"id":1}`)))
	req.Header.Add("x-access-token", token)

	user := beans.User{
		Email: faker.Internet().Email(),
	}

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(&user, nil)

	favManagerMock.EXPECT().
		DeleteFav(gomock.Eq(1)).
		Return(nil)

	favManagerMock.EXPECT().
		GetFavs(gomock.Eq(user.Email)).
		Return(nil, errors.New(""))

	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Result().StatusCode)
}

func Test_controller_PostFavRemoveCtrl_error_delete(t *testing.T) {
	r, w, _, userManagerMock, _, favManagerMock := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav/remove", bytes.NewBuffer([]byte(`{"id":1}`)))
	req.Header.Add("x-access-token", token)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	favManagerMock.EXPECT().
		DeleteFav(gomock.Eq(1)).
		Return(errors.New(""))

	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Result().StatusCode)
}

func Test_controller_PostFavRemoveCtrl_error_badRequest(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav/remove", bytes.NewBuffer([]byte(`{"id":"1"`)))
	req.Header.Add("x-access-token", token)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)
}

func Test_controller_PostFavRemoveCtrl_error_userNotFound(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav/remove", nil)
	req.Header.Add("x-access-token", token)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, errors.New(""))

	r.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Result().StatusCode)
}

func Test_controller_PostFavRemoveCtrl_error_auth(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav/remove", nil)
	req.Header.Add("x-access-token", token)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, errors.New(""))

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}

func Test_controller_PostFavCtrl_success(t *testing.T) {
	r, w, _, userManagerMock, _, favManagerMock := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav", bytes.NewBuffer([]byte(`{"surah":1,"verse":1}`)))
	req.Header.Add("x-access-token", token)

	user := beans.User{
		Email: faker.Internet().Email(),
	}

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(&user, nil)

	favManagerMock.EXPECT().
		CreateFav(gomock.Eq(user.Email), gomock.Eq(1), gomock.Eq(1)).
		Return(nil)

	favManagerMock.EXPECT().
		GetFavs(gomock.Eq(user.Email)).
		Return([]*favorite.Fav{
			{
				Id:    1,
				Surah: 1,
				Verse: 1,
			},
			{
				Id:    2,
				Surah: 1,
				Verse: 2,
			},
		}, nil)

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, `{"data":{"favorites":[{"id":1,"surah":1,"verse":1},{"id":2,"surah":1,"verse":2}]}}`, w.Body.String())
}

func Test_controller_PostFavCtrl_error_favs(t *testing.T) {
	r, w, _, userManagerMock, _, favManagerMock := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav", bytes.NewBuffer([]byte(`{"surah":1,"verse":1}`)))
	req.Header.Add("x-access-token", token)

	user := beans.User{
		Email: faker.Internet().Email(),
	}

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(&user, nil)

	favManagerMock.EXPECT().
		CreateFav(gomock.Eq(user.Email), gomock.Eq(1), gomock.Eq(1)).
		Return(nil)

	favManagerMock.EXPECT().
		GetFavs(gomock.Eq(user.Email)).
		Return(nil, errors.New(""))

	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Result().StatusCode)
}

func Test_controller_PostFavCtrl_error_create(t *testing.T) {
	r, w, _, userManagerMock, _, favManagerMock := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	user := beans.User{
		Email: faker.Internet().Email(),
	}

	req, _ := http.NewRequest(http.MethodPost, "/fav", bytes.NewBuffer([]byte(`{"surah":1,"verse":1}`)))
	req.Header.Add("x-access-token", token)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(&user, nil)

	favManagerMock.EXPECT().
		CreateFav(gomock.Eq(user.Email), gomock.Eq(1), gomock.Eq(1)).
		Return(errors.New(""))

	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Result().StatusCode)
}

func Test_controller_PostFavCtrl_error_badRequest(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav", bytes.NewBuffer([]byte(`{"surah":"1","verse":"1"}`)))
	req.Header.Add("x-access-token", token)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)
}

func Test_controller_PostFavCtrl_error_userNotFound(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav", bytes.NewBuffer([]byte(`{"surah":1,"verse":1}`)))
	req.Header.Add("x-access-token", token)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, errors.New(""))

	r.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Result().StatusCode)
}

func Test_controller_PostFavCtrl_error_auth(t *testing.T) {
	r, w, _, userManagerMock, _, _ := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodPost, "/fav", bytes.NewBuffer([]byte(`{"surah":1,"verse":1}`)))
	req.Header.Add("x-access-token", token)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, errors.New(""))

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}

func Test_controller_GetFavCtrl(t *testing.T) {
	r, w, _, userManagerMock, _, favManagerMock := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	req, _ := http.NewRequest(http.MethodGet, "/fav", nil)
	req.Header.Add("x-access-token", token)

	user := beans.User{
		Email: faker.Internet().Email(),
	}

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(&user, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(&user, nil)

	favManagerMock.EXPECT().
		GetFavs(gomock.Eq(user.Email)).
		Return([]*favorite.Fav{
			{
				Id:    1,
				Surah: 1,
				Verse: 1,
			},
			{
				Id:    2,
				Surah: 1,
				Verse: 2,
			},
		}, nil)

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, `{"data":{"favorites":[{"id":1,"surah":1,"verse":1},{"id":2,"surah":1,"verse":2}]}}`, w.Body.String())
}

func Test_controller_GetFavCtrl_error(t *testing.T) {
	r, w, _, userManagerMock, _, favManagerMock := routes.SetupTestRoutes(t)

	faker := faker.New()
	token := faker.Internet().Password()

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(nil, errors.New("user not found"))

	req, _ := http.NewRequest(http.MethodGet, "/fav", nil)
	req.Header.Add("x-access-token", token)

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)

	user := beans.User{
		Email: faker.Internet().Email(),
	}

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(&user, nil)

	userManagerMock.EXPECT().
		GetUser(gomock.Eq(token)).
		Return(&user, nil)

	favManagerMock.EXPECT().
		GetFavs(gomock.Eq(user.Email)).
		Return(nil, errors.New(""))

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}
