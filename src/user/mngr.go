package user

type Mngr interface {
	Create(username string, password string) (*User, error)
	Get(username string) (*User, error)
	Update_surah_verse(username string, surah, verse int) error
	Update_setting(username string, name string, value string) error
}
