package favorite_test

import (
	"testing"

	"github.com/arikama/koran-backend/favorite"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func Test_dao_impl(t *testing.T) {
	var favDao favorite.FavDao
	favDao, err := favorite.NewFavDaoImpl()
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

	favDao.DeleteFav(queried[0].ID)

	queried, err = favDao.QueryUserFavsByEmail(email)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(queried))
}
