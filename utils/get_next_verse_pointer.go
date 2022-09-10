package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arikama/koran-backend/constants"
	"github.com/hooligram/kifu"
)

func GetNextVersePointer(pointer string) string {
	arr := strings.Split(pointer, ":")
	if len(arr) < 2 {
		return constants.StartPointer()
	}

	surahId, err := strconv.Atoi(arr[0])
	if err != nil {
		kifu.Error("Error parsing surah id pointer: %v", err.Error())
		return constants.StartPointer()
	}
	if surahId >= len(constants.SurahPointerVerseEndings()) {
		return constants.StartPointer()
	}

	verseId, err := strconv.Atoi(arr[1])
	if err != nil {
		kifu.Error("Error parsing verse id pointer: %v", err.Error())
		return constants.StartPointer()
	}

	verseEnd := constants.SurahPointerVerseEndings()[surahId]
	if verseId > verseEnd {
		return constants.StartPointer()
	}
	if verseId < verseEnd {
		verseId += 1
	} else {
		verseId = 1
		if surahId >= constants.SurahPointerEnding() {
			surahId = 1
		} else {
			surahId += 1
		}
	}

	return fmt.Sprintf("%v:%v", surahId, verseId)
}
