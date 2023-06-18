package favorite

//go:generate mockgen -package=favorite -mock_names=FavManager=FavManagerMock -source=./manager.go -destination=./manager_mock.go
type FavManager interface {
	CreateFav(email string, surah, verse int) error
	GetFavs(email string) ([]*Fav, error)
	DeleteFav(id int) error
}

type Fav struct {
	Id    int
	Surah int
	Verse int
}
