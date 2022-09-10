package managers_test

import (
	"errors"
	"testing"

	"github.com/arikama/koran-backend/beans"
	"github.com/arikama/koran-backend/daos"
	"github.com/arikama/koran-backend/managers"
	gomock "github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := daos.NewUserDaoMock(ctrl)

	var userManager managers.UserManager
	userManager, err := managers.NewUserManagerImpl(userDaoMock)
	assert.Nil(t, err)

	faker := faker.New()
	email := faker.Internet().Email()
	token := faker.Internet().Password()

	userDaoMock.EXPECT().
		QueryUserByEmail(gomock.Eq(email)).
		Return(nil, errors.New(daos.ErrSqlNoRowsInResultSet()))

	userDaoMock.EXPECT().
		CreateUser(gomock.Eq(email), gomock.Eq(token)).
		Return(nil)

	userDaoMock.EXPECT().
		QueryUserByToken(gomock.Eq(token)).
		Return(&beans.User{
			Email: email,
			Token: token,
		}, nil)

	userDaoMock.EXPECT().
		UpdateUserCurrentPointer(gomock.Eq(email), "1:1").
		Return(nil)

	result, err := userManager.CreateUser(email, token)
	assert.Nil(t, err)
	assert.Equal(t, email, result.Email)
	assert.Equal(t, token, result.Token)
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := daos.NewUserDaoMock(ctrl)

	var userManager managers.UserManager
	userManager, err := managers.NewUserManagerImpl(userDaoMock)
	assert.Nil(t, err)

	faker := faker.New()
	email := faker.Internet().Email()
	token := faker.Internet().Password()

	userDaoMock.EXPECT().
		QueryUserByToken(gomock.Eq(token)).
		Return(&beans.User{
			Email: email,
			Token: token,
		}, nil)

	result, err := userManager.GetUser(token)
	assert.Nil(t, err)
	assert.Equal(t, email, result.Email)
	assert.Equal(t, token, result.Token)
}

func TestAdvanceUserCurrentPointer(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDaoMock := daos.NewUserDaoMock(ctrl)

	var userManager managers.UserManager
	userManager, err := managers.NewUserManagerImpl(userDaoMock)
	assert.Nil(t, err)

	faker := faker.New()
	email := faker.Internet().Email()
	token := faker.Internet().Password()

	userDaoMock.EXPECT().
		QueryUserByEmail(gomock.Eq(email)).
		Return(&beans.User{
			Email:          email,
			Token:          token,
			CurrentPointer: "1:1",
		}, nil)

	userDaoMock.EXPECT().
		UpdateUserCurrentPointer(gomock.Eq(email), gomock.Eq("1:2")).
		Return(nil)

	result, err := userManager.AdvanceUserCurrentPointer(email, token)
	assert.Nil(t, err)
	assert.Equal(t, "1:2", result)
}
