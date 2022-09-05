package beans

type Surah struct {
	SurahInfo *SurahInfo `json:"surah_info"`
	Verses    []*Verse   `json:"verses"`
}

func NewSurah(surahInfo SurahInfo) *Surah {
	return &Surah{
		SurahInfo: &surahInfo,
		Verses:    []*Verse{},
	}
}
