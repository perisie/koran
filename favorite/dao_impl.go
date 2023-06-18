package favorite

import (
	"context"
	"database/sql"

	"github.com/arikama/koran-backend/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type FavDaoImpl struct {
	context *context.Context
	db      *sql.DB
}

func NewFavDaoImpl(db *sql.DB) (*FavDaoImpl, error) {
	context := context.Background()

	return &FavDaoImpl{
		context: &context,
		db:      db,
	}, nil
}

func (f *FavDaoImpl) AddFavVerse(email string, surah, verse int) error {
	fav := models.Fav{
		Email: email,
		Surah: int16(surah),
		Verse: int16(verse),
	}

	err := fav.Insert(*f.context, f.db, boil.Infer())

	if err != nil {
		return err
	}

	return nil
}

func (f *FavDaoImpl) QueryUserFavsByEmail(email string) ([]*models.Fav, error) {
	query := models.Favs(
		qm.Where("email = ?", email),
		qm.OrderBy("surah ASC, verse ASC"),
	)

	favs, err := query.All(*f.context, f.db)

	if err != nil {
		return nil, err
	}

	return favs, nil
}

func (f *FavDaoImpl) DeleteFav(id int) error {
	fav, err := models.FindFav(*f.context, f.db, id)

	if err != nil {
		return err
	}

	_, err = fav.Delete(*f.context, f.db)

	return err
}
