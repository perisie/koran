package beans

import "fmt"

type Verse struct {
	SurahId      int
	VerseId      int
	Text         string
	Translations map[string]string
}

func (v *Verse) Key() string {
	return fmt.Sprintf("%v:%v", v.SurahId, v.VerseId)
}
