package beans

type Verse struct {
	SurahId      int               `json:"surah_id"`
	VerseId      int               `json:"verse_id"`
	Text         string            `json:"text"`
	Translations map[string]string `json:"translations"`
}
