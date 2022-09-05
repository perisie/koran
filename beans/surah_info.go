package beans

type SurahInfo struct {
	Id                int
	Title             string
	NumberOfVerses    int
	PlaceOfRevelation string
	JuzStart          int
	JuzEnd            int
}

func NewSurahInfo(id int, title string, numberOfVerses int, placeOfRevelation string, juzStart, juzEnd int) *SurahInfo {
	return &SurahInfo{}
}
