package quran

type Verse struct {
	Surah_id     int
	Verse_id     int
	Text         string
	Translations map[string]string
}
