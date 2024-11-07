package quran

type Surah struct {
	SurahInfo *Surah_info
	Verses    []*Verse
}

func Surah_new(surah_info Surah_info) *Surah {
	return &Surah{
		SurahInfo: &surah_info,
		Verses:    []*Verse{},
	}
}
