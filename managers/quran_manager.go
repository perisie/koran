package managers

import "perisie.com/koran/beans"

type QuranManager interface {
	GetVerse(surahId, verseId int) (*beans.Verse, error)
	GetSurah(surahId int) (*beans.Surah, error)
	GetSurahInfos() ([]*beans.SurahInfo, error)
}

func Quran() string      { return "quran" }
func Pickthall() string  { return "pickthall" }
func ClearQuran() string { return "clearquran" }

func ErrVerseDoesNotExist() string { return "verse does not exist" }
func ErrSurahDoesNotExist() string { return "surah does not exist" }
