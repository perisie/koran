package daos

import (
	"context"
	"database/sql"

	"github.com/arikama/koran-backend/beans"
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

func (f *FavDaoImpl) QueryUserFavsByEmail(email string) ([]*beans.Fav, error) {
	favs, err := models.Favs(
		qm.Where("email = ?", email),
		qm.OrderBy("surah ASC, verse ASC"),
	).All(*f.context, f.db)
	if err != nil {
		return nil, err
	}
	results := []*beans.Fav{}
	for _, fav := range favs {
		results = append(results, &beans.Fav{
			Email: fav.Email,
			Surah: int(fav.Surah),
			Verse: int(fav.Verse),
		})
	}
	return results, nil
}
