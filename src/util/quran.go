package util

import (
	"errors"

	"perisie.com/koran/src/quran"
)

func Move_surah_verse(surah, verse, direction int) (int, int, error) {
	if surah < quran.Surah_start() || surah > quran.Surah_ending() {
		return 0, 0, errors.New("invalid surah")
	}
	if verse < quran.Verse_start() || verse > quran.Surah_verse_endings()[surah] {
		return 0, 0, errors.New("invalid verse")
	}
	if direction > 0 {
		verse += direction
		for verse > quran.Surah_verse_endings()[surah] {
			verse -= quran.Surah_verse_endings()[surah]
			surah += 1
			if surah > quran.Surah_ending() {
				surah = 1
			}
		}
	} else if direction < 0 {
		verse += direction
		for verse < quran.Verse_start() {
			surah -= 1
			if surah < quran.Surah_start() {
				surah = quran.Surah_ending()
			}
			verse += quran.Surah_verse_endings()[surah]
		}
	}
	return surah, verse, nil
}
