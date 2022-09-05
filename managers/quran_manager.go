package managers

import "github.com/arikama/koran-backend/beans"

type QuranManager interface {
	GetVerse(surahId, verseId int) (*beans.Verse, error)
}

func Quran() string     { return "quran" }
func Pickthall() string { return "pickthall" }
