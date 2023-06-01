package daos

import (
	"github.com/arikama/koran-backend/beans"
)

//go:generate mockgen -package=daos -mock_names=FavDao=FavDaoMock -source=./fav_dao.go -destination=./fav_dao_mock.go
type FavDao interface {
	AddFavVerse(email string, surah, verse int) error
	QueryUserFavsByEmail(email string) ([]*beans.Fav, error)
}
