package favorite_test

import (
	"testing"

	"github.com/arikama/koran-backend/favorite"
	"github.com/arikama/koran-backend/utils"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestAddFavQueryFavs(t *testing.T) {
	db, err := utils.GetTestDb()
	assert.Nil(t, err)

	var favDao favorite.FavDao
	favDao, err = favorite.NewFavDaoImpl(db)
	assert.Nil(t, err)

	faker := faker.New()
	email := faker.Internet().Email()

	for i := 0; i < 3; i++ {
		surah := 3 - i
		verse := 3 - i

		err = favDao.AddFavVerse(email, surah, verse)
		assert.Nil(t, err)
	}

	queried, err := favDao.QueryUserFavsByEmail(email)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(queried))

	for i := 0; i < 3; i++ {
		assert.Equal(t, email, queried[i].Email)
		assert.Equal(t, int16(i+1), queried[i].Surah)
		assert.Equal(t, int16(i+1), queried[i].Verse)
	}
}
