package favorite

import (
	"errors"
	"strings"

	"perisie.com/koran/constants"
	"perisie.com/koran/daos"
	"perisie.com/koran/managers"
)

type FavManagerImpl struct {
	favDao  FavDao
	userDao daos.UserDao
}

func NewFavManagerImpl(favDao FavDao, userDao daos.UserDao) (*FavManagerImpl, error) {
	return &FavManagerImpl{
		favDao:  favDao,
		userDao: userDao,
	}, nil
}

func (f *FavManagerImpl) CreateFav(email string, surah, verse int) error {
	if surah < 1 || surah > constants.SurahPointerEnding() {
		return errors.New(managers.ErrSurahDoesNotExist())
	}

	if verse < 1 || verse > constants.SurahPointerVerseEndings()[surah] {
		return errors.New(managers.ErrVerseDoesNotExist())
	}

	user, err := f.userDao.QueryUserByEmail(email)

	if err != nil {
		return err
	}

	err = f.favDao.AddFavVerse(user.Email, surah, verse)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil
		}

		return err
	}

	return nil
}

func (f *FavManagerImpl) GetFavs(email string) ([]*Fav, error) {
	userFavs, err := f.favDao.QueryUserFavsByEmail(email)

	if err != nil {
		return nil, err
	}

	favs := []*Fav{}

	for _, fav := range userFavs {
		favs = append(favs, &Fav{
			Id:    fav.ID,
			Surah: int(fav.Surah),
			Verse: int(fav.Verse),
		})
	}

	return favs, nil
}

func (f *FavManagerImpl) DeleteFav(id int) error {
	return f.favDao.DeleteFav(id)
}
