package beans

type SurahInfo struct {
	SurahId  int    `json:"surah_id"`
	Title    string `json:"title"`
	Arabic   string `json:"arabic"`
	English  string `json:"english"`
	Verses   int    `json:"verses"`
	City     string `json:"city"`
	JuzStart int    `json:"juz_start"`
	JuzEnd   int    `json:"juz_end"`
}
