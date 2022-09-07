package utils

import "fmt"

func GetSurahVerseKey(surahId, verseId int) string {
	return fmt.Sprintf("%v:%v", surahId, verseId)
}
