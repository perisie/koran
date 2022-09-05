package managers

import "github.com/arikama/koran-backend/beans"

type QuranManager interface {
	GetVerse(surahId, verseId int) (*beans.Verse, error)
	GetSurah(surahId int) (*beans.Surah, error)
	GetSurahInfos() ([]*beans.SurahInfo, error)
}

func Quran() string     { return "quran" }
func Pickthall() string { return "pickthall" }
