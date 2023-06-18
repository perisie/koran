package favorite

import (
	"github.com/arikama/koran-backend/models"
)

//go:generate mockgen -package=favorite -mock_names=FavDao=FavDaoMock -source=./dao.go -destination=./dao_mock.go
type FavDao interface {
	AddFavVerse(email string, surah, verse int) error
	QueryUserFavsByEmail(email string) ([]*models.Fav, error)
	DeleteFav(id int) error
}
