package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/arikama/koran-backend/constants"
	"github.com/hooligram/kifu"
)

func GetNextVersePointer(pointer string, direction int) string {
	splits := strings.Split(pointer, ":")
	surah, err := strconv.Atoi(splits[0])
	if err != nil {
		kifu.Error("Error parsing surah id pointer: %v", err.Error())
		return ""
	}
	verse, err := strconv.Atoi(splits[1])
	if err != nil {
		kifu.Error("Error parsing verse id pointer: %v", err.Error())
		return ""
	}
	surahNext, verseNext, err := MoveSurahVerse(surah, verse, direction)
	if err != nil {
		kifu.Error("Error moving surah verse: surah=%v, verse=%v", surah, verse)
		return ""
	}
	return fmt.Sprintf("%v:%v", surahNext, verseNext)
}

func MoveSurahVerse(surah, verse, direction int) (int, int, error) {
	if surah < constants.SurahPointerStart() || surah > constants.SurahPointerEnding() {
		return 0, 0, errors.New("invalid surah")
	}
	if verse < constants.VersePointerStart() || verse > constants.SurahPointerVerseEndings()[surah] {
		return 0, 0, errors.New("invalid verse")
	}
	if direction > 0 {
		verse += direction
		for verse > constants.SurahPointerVerseEndings()[surah] {
			verse -= constants.SurahPointerVerseEndings()[surah]
			surah += 1
			if surah > constants.SurahPointerEnding() {
				surah = 1
			}
		}
	} else if direction < 0 {
		verse += direction
		for verse < constants.VersePointerStart() {
			surah -= 1
			if surah < constants.SurahPointerStart() {
				surah = constants.SurahPointerEnding()
			}
			verse += constants.SurahPointerVerseEndings()[surah]
		}
	}
	return surah, verse, nil
}
