package favorite

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/fs"
	"perisie.com/koran/models"
	"perisie.com/koran/mouse"
	"sort"
)

type FavDaoImpl struct {
	mouse *mouse.Mouse
}

func NewFavDaoImpl() (*FavDaoImpl, error) {
	return &FavDaoImpl{
		mouse: mouse.Mouse_new("./data"),
	}, nil
}

func (f *FavDaoImpl) AddFavVerse(email string, surah, verse int) error {
	favs, err := f.QueryUserFavsByEmail(email)
	if err != nil {
		return err
	}
	fav := &models.Fav{
		Email: email,
		Surah: int16(surah),
		Verse: int16(verse),
	}
	favs = append(favs, fav)
	value, err := mouse.To_byte(favs)
	if err != nil {
		return err
	}
	key := f.key_fav(email)
	err = f.mouse.Put(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (f *FavDaoImpl) QueryUserFavsByEmail(email string) ([]*models.Fav, error) {
	key := f.key_fav(email)
	favs := []*models.Fav{}
	value, err := f.mouse.Get(key)
	if err != nil {
		if _, ok := err.(*fs.PathError); ok {
			value, err = mouse.To_byte(favs)
		}
		if err != nil {
			return nil, err
		}
	}
	err = gob.NewDecoder(bytes.NewReader(value)).Decode(&favs)
	if err != nil {
		return nil, err
	}
	sort.Slice(favs, func(i, j int) bool {
		if favs[i].Surah < favs[j].Surah {
			return true
		}
		return favs[i].Verse < favs[j].Verse
	})
	return favs, nil
}

func (f *FavDaoImpl) DeleteFav(id int) error {
	return nil
}

func (f *FavDaoImpl) key_fav(email string) string {
	return fmt.Sprintf("fav__%v", email)
}
