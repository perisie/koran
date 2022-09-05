package beans

type Surah struct {
	surahId  int
	surahLen int
	verses   []*Verse
}

func NewSurah(surahId, surahLen int) *Surah {
	return &Surah{
		surahId:  surahId,
		surahLen: surahLen,
		verses:   []*Verse{},
	}
}
