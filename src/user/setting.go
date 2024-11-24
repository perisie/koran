package user

type Setting struct {
	Surah_verse          bool `json:"surah_verse"`
	Surah_translation    bool `json:"surah_translation"`
	Bookmark_verse       bool `json:"bookmark_verse"`
	Bookmark_translation bool `json:"bookmark_translation"`
}

func Setting_new() *Setting {
	return &Setting{
		Surah_verse:          true,
		Surah_translation:    true,
		Bookmark_verse:       true,
		Bookmark_translation: true,
	}
}
