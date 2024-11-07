package quran

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Mngr_impl struct {
	surah_map   map[int]*Surah
	verse_map   map[string]*Verse
	surah_infos []*Surah_info
}

func (m *Mngr_impl) Get_verse(surah_id, verse_id int) (*Verse, error) {
	key := surah_verse_key(surah_id, verse_id)
	verse, ok := m.verse_map[key]
	if !ok {
		return nil, errors.New("verse not found")
	}
	return verse, nil
}

func (m *Mngr_impl) Get_surah(surah_id int) (*Surah, error) {
	surah, ok := m.surah_map[surah_id]
	if !ok {
		return nil, errors.New("surah not found")
	}
	return surah, nil
}

func (m *Mngr_impl) Get_surah_infos() ([]*Surah_info, error) {
	return m.surah_infos, nil
}

func Mngr_impl_new(csv_dir string) (*Mngr_impl, error) {
	names := []string{
		NAME_QURAN,
		NAME_PICKTHALL,
		NAME_CLEARQURAN,
	}
	verse_map := map[string]*Verse{}
	for _, name := range names {
		file_path := fmt.Sprintf("%v/%v.csv", csv_dir, name)
		records, err := read_quran_csv(file_path)
		if err != nil {
			return nil, err
		}
		for _, record := range records {
			surah_id, err_record := strconv.Atoi(record[0])
			if err_record != nil {
				return nil, err_record
			}
			verse_id, err_record := strconv.Atoi(record[1])
			if err_record != nil {
				return nil, err_record
			}
			key := surah_verse_key(surah_id, verse_id)
			if _, ok := verse_map[key]; !ok {
				verse := Verse{}
				verse.Surah_id = surah_id
				verse.Verse_id = verse_id
				verse_map[key] = &verse
			}
			text := strings.Trim(record[2], `"`)
			if name == NAME_QURAN {
				verse_map[key].Text = text
			} else {
				if verse_map[key].Translations == nil {
					verse_map[key].Translations = map[string]string{}
				}
				verse_map[key].Translations[name] = text
			}
		}
	}
	surah_map := map[int]*Surah{}
	surah_infos, err := read_surah_info(fmt.Sprintf("%v/surahs.csv", csv_dir))
	if err != nil {
		return nil, err
	}
	for _, surah_info := range surah_infos {
		surah_map[surah_info.Surah_id] = Surah_new(*surah_info)
		for verseId := 1; verseId <= surah_info.Verses; verseId++ {
			key := surah_verse_key(surah_info.Surah_id, verseId)
			surah_map[surah_info.Surah_id].Verses = append(surah_map[surah_info.Surah_id].Verses, verse_map[key])
		}
	}
	return &Mngr_impl{
		surah_map:   surah_map,
		verse_map:   verse_map,
		surah_infos: surah_infos,
	}, nil
}
