package favorite_test

import (
	"errors"
	"testing"

	"github.com/arikama/koran-backend/beans"
	"github.com/arikama/koran-backend/daos"
	"github.com/arikama/koran-backend/favorite"
	"github.com/arikama/koran-backend/models"
	gomock "github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func Test_manager_impl_CreateFav(t *testing.T) {
	ctrl := gomock.NewController(t)

	favDaoMock := favorite.NewFavDaoMock(ctrl)
	userDaoMock := daos.NewUserDaoMock(ctrl)

	favManager, err := favorite.NewFavManagerImpl(favDaoMock, userDaoMock)
	assert.Nil(t, err)

	faker := faker.New()
	email := faker.Internet().Email()

	err = favManager.CreateFav(email, 0, 1)
	assert.NotNil(t, err)
	assert.Equal(t, "surah does not exist", err.Error())

	err = favManager.CreateFav(email, 1, 0)
	assert.NotNil(t, err)
	assert.Equal(t, "verse does not exist", err.Error())

	userDaoMock.EXPECT().
		QueryUserByEmail(gomock.Eq(email)).
		Return(nil, errors.New("user not found"))

	err = favManager.CreateFav(email, 1, 1)
	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error())

	user := beans.User{
		Email: email,
	}

	userDaoMock.EXPECT().
		QueryUserByEmail(gomock.Eq(email)).
		Return(&user, nil)

	favDaoMock.EXPECT().
		AddFavVerse(gomock.Eq(email), gomock.Eq(1), gomock.Eq(1)).
		Return(errors.New("not added"))

	err = favManager.CreateFav(email, 1, 1)
	assert.NotNil(t, err)

	userDaoMock.EXPECT().
		QueryUserByEmail(gomock.Eq(email)).
		Return(&user, nil)

	favDaoMock.EXPECT().
		AddFavVerse(gomock.Eq(email), gomock.Eq(1), gomock.Eq(1)).
		Return(nil)

	err = favManager.CreateFav(email, 1, 1)
	assert.Nil(t, err)
}

func Test_manager_impl_CreateFav_duplicate(t *testing.T) {
	ctrl := gomock.NewController(t)

	favDaoMock := favorite.NewFavDaoMock(ctrl)
	userDaoMock := daos.NewUserDaoMock(ctrl)

	favManager, err := favorite.NewFavManagerImpl(favDaoMock, userDaoMock)
	assert.Nil(t, err)

	faker := faker.New()
	email := faker.Internet().Email()

	user := beans.User{
		Email: email,
	}

	userDaoMock.EXPECT().
		QueryUserByEmail(gomock.Eq(email)).
		Return(&user, nil)

	favDaoMock.EXPECT().
		AddFavVerse(gomock.Eq(email), gomock.Eq(1), gomock.Eq(1)).
		Return(errors.New(" Duplicate entry "))

	err = favManager.CreateFav(email, 1, 1)
	assert.Nil(t, err)
}

func Test_manager_impl_GetFavs(t *testing.T) {
	ctrl := gomock.NewController(t)

	favDaoMock := favorite.NewFavDaoMock(ctrl)
	userDaoMock := daos.NewUserDaoMock(ctrl)

	favManager, err := favorite.NewFavManagerImpl(favDaoMock, userDaoMock)
	assert.Nil(t, err)

	faker := faker.New()
	email := faker.Internet().Email()

	favDaoMock.EXPECT().
		QueryUserFavsByEmail(gomock.Eq(email)).
		Return(nil, errors.New("favs not found"))

	_, err = favManager.GetFavs(email)
	assert.NotNil(t, err)
	assert.Equal(t, "favs not found", err.Error())

	favDaoMock.EXPECT().
		QueryUserFavsByEmail(gomock.Eq(email)).
		Return([]*models.Fav{
			{
				ID:    1,
				Surah: 1,
				Verse: 1,
			},
			{
				ID:    2,
				Surah: 1,
				Verse: 2,
			},
		}, nil)

	favs, err := favManager.GetFavs(email)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(favs))
}

func Test_manager_impl_DeleteFav(t *testing.T) {
	ctrl := gomock.NewController(t)

	favDaoMock := favorite.NewFavDaoMock(ctrl)
	userDaoMock := daos.NewUserDaoMock(ctrl)

	favManager, err := favorite.NewFavManagerImpl(favDaoMock, userDaoMock)
	assert.Nil(t, err)

	favDaoMock.EXPECT().
		DeleteFav(gomock.Eq(1)).
		Return(nil)

	err = favManager.DeleteFav(1)
	assert.Nil(t, err)
}
