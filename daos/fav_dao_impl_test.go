package daos_test

import (
	"testing"

	"github.com/arikama/koran-backend/beans"
	"github.com/arikama/koran-backend/daos"
	"github.com/arikama/koran-backend/utils"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestAddFavQueryFavs(t *testing.T) {
	db, err := utils.GetTestDb()
	assert.Nil(t, err)

	var favDao daos.FavDao
	favDao, err = daos.NewFavDaoImpl(db)
	assert.Nil(t, err)

	faker := faker.New()
	email := faker.Internet().Email()

	for i := 0; i < 3; i++ {
		surah := 3 - i
		verse := 3 - i

		fav := beans.Fav{
			Email: email,
			Surah: surah,
			Verse: verse,
		}

		err = favDao.AddFavVerse(fav.Email, fav.Surah, fav.Verse)
		assert.Nil(t, err)
	}

	queried, err := favDao.QueryUserFavsByEmail(email)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(queried))

	for i := 0; i < 3; i++ {
		assert.Equal(t, email, queried[i].Email)
		assert.Equal(t, i+1, queried[i].Surah)
		assert.Equal(t, i+1, queried[i].Verse)
	}
}
