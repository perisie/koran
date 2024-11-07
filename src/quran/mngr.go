package quran

const (
	NAME_QURAN      = "quran"
	NAME_PICKTHALL  = "pickthall"
	NAME_CLEARQURAN = "clearquran"
)

type Mngr interface {
	Get_verse(surah_id, verse_id int) (*Verse, error)
	Get_surah(surah_id int) (*Surah, error)
	Get_surah_infos() ([]*Surah_info, error)
}
