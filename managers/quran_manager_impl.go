package managers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/arikama/koran-backend/beans"
)

type QuranManagerImpl struct {
	verseMap map[string]*beans.Verse
}

func NewQuranManagerImpl(csvDir string) (*QuranManagerImpl, error) {
	names := []string{
		Quran(),
		Pickthall(),
	}
	verseMap := make(map[string]*beans.Verse)
	for _, name := range names {
		filePath := fmt.Sprintf("%v/%v.csv", csvDir, name)
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}
		for _, record := range records {
			surahId, err := strconv.Atoi(strings.Trim(record[0], "\ufeff"))
			if err != nil {
				return nil, err
			}
			verseId, err := strconv.Atoi(record[1])
			if err != nil {
				return nil, err
			}
			key := fmt.Sprintf("%v:%v", surahId, verseId)
			if _, ok := verseMap[key]; !ok {
				verse := beans.Verse{}
				verse.SurahId = surahId
				verse.VerseId = verseId
				verseMap[key] = &verse
			}
			text := strings.Trim(record[2], `"`)
			if name == Quran() {
				verseMap[key].Text = text
			} else {
				if verseMap[key].Translations == nil {
					verseMap[key].Translations = map[string]string{}
				}
				verseMap[key].Translations[name] = text
			}
		}
	}
	return &QuranManagerImpl{
		verseMap: verseMap,
	}, nil
}

func (q *QuranManagerImpl) GetVerse(surahId, verseId int) (*beans.Verse, error) {
	key := fmt.Sprintf("%v:%v", surahId, verseId)
	verse, ok := q.verseMap[key]
	if !ok {
		return nil, errors.New("verse does not exist")
	}
	return verse, nil
}
