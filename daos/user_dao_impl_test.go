package daos_test

import (
	"testing"

	"github.com/arikama/koran-backend/beans"
	"github.com/arikama/koran-backend/constants"
	"github.com/arikama/koran-backend/daos"
	"github.com/arikama/koran-backend/utils"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuery(t *testing.T) {
	db, err := utils.GetTestDb()
	assert.Nil(t, err)

	var userDao daos.UserDao
	userDao, err = daos.NewUserDaoImpl(db)
	assert.Nil(t, err)

	faker := faker.New()

	email := faker.Internet().Email()
	name := faker.Person().Name()
	token := faker.Internet().Password()
	picture := faker.Internet().URL()

	user := beans.User{
		Email:          email,
		Name:           name,
		Token:          token,
		Picture:        picture,
		CurrentPointer: constants.StartPointer(),
	}

	err = userDao.CreateUser(user.Email, user.Token)
	assert.Nil(t, err)

	queried, err := userDao.QueryUserByEmail(email)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, queried.Email)

	queried, err = userDao.QueryUserByToken(token)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, queried.Email)
}
