package managers_test

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"perisie.com/koran/beans"
	"perisie.com/koran/daos"
	"perisie.com/koran/managers"
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

func TestCreateUserQueryError(t *testing.T) {
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
		Return(nil, errors.New("error"))

	_, err = userManager.CreateUser(email, token)
	assert.Equal(t, "error", err.Error())
}

func Test_CreateUser_AlreadyExist(t *testing.T) {
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
			Email: email,
			Token: token,
		}, nil)

	userDaoMock.EXPECT().
		UpdateUserToken(gomock.Eq(email), gomock.Eq(token)).
		Return(nil)

	result, err := userManager.CreateUser(email, token)
	assert.Nil(t, err)
	assert.Equal(t, email, result.Email)
	assert.Equal(t, token, result.Token)
}

func Test_CreateUser_CreateError(t *testing.T) {
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
		Return(nil, nil)

	userDaoMock.EXPECT().
		CreateUser(gomock.Eq(email), gomock.Eq(token)).
		Return(errors.New("error"))

	_, err = userManager.CreateUser(email, token)
	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}

func Test_CreateUser_QueryError(t *testing.T) {
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
		Return(nil, nil)

	userDaoMock.EXPECT().
		CreateUser(gomock.Eq(email), gomock.Eq(token)).
		Return(nil)

	userDaoMock.EXPECT().
		QueryUserByToken(gomock.Eq(token)).
		Return(nil, errors.New("error"))

	_, err = userManager.CreateUser(email, token)
	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}

func Test_CreateUser_UpdatePointerError(t *testing.T) {
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
		Return(nil, nil)

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
		UpdateUserCurrentPointer(gomock.Eq(email), gomock.Eq("1:1")).
		Return(errors.New("error"))

	_, err = userManager.CreateUser(email, token)
	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
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

func Test_GetUser_Error(t *testing.T) {
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
		}, errors.New("error"))

	_, err = userManager.GetUser(token)
	assert.NotNil(t, err)
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

func Test_AdvanceUserCurrentPointer_QueryError(t *testing.T) {
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
		}, errors.New("error"))

	_, err = userManager.AdvanceUserCurrentPointer(email, token)
	assert.NotNil(t, err)
}

func Test_AdvanceUserCurrentPointer_PointerUpdateError(t *testing.T) {
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
		Return(errors.New("error"))

	_, err = userManager.AdvanceUserCurrentPointer(email, token)
	assert.NotNil(t, err)
}

func Test_AdvanceUserCurrentPointer_TokenMismatch(t *testing.T) {
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
			Token:          token + token,
			CurrentPointer: "1:1",
		}, nil)

	_, err = userManager.AdvanceUserCurrentPointer(email, token)
	assert.NotNil(t, err)
	assert.Equal(t, "error user token mismatch", err.Error())
}
