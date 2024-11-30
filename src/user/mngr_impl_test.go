package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	mngr := Mngr_impl_fake()

	username := "faithes"
	password := "password"

	user, err := mngr.Get(username)

	assert.NotNil(t, err)
	assert.Equal(t, "", user.Username)
	assert.False(t, user.Ok())

	user, err = mngr.Create(username, password)

	assert.Nil(t, err)
	assert.Equal(t, username, user.Username)
	assert.NotEqual(t, password, user.Password)

	user, err = mngr.Get(username)

	assert.Nil(t, err)
	assert.Equal(t, username, user.Username)
	assert.NotEqual(t, password, user.Password)
	assert.Equal(t, 1, user.Surah)
	assert.Equal(t, 1, user.Verse)

	_ = mngr.Update_surah_verse(user.Username, 2, 3)
	user, _ = mngr.Get(username)

	assert.Equal(t, 2, user.Surah)
	assert.Equal(t, 3, user.Verse)

	assert.True(t, user.Setting.Surah_verse)
	assert.True(t, user.Setting.Surah_translation)
	assert.True(t, user.Setting.Bookmark_verse)
	assert.True(t, user.Setting.Bookmark_translation)

	_ = mngr.Update_setting(user.Username, "surah_verse", "false")
	_ = mngr.Update_setting(user.Username, "surah_translation", "false")
	_ = mngr.Update_setting(user.Username, "bookmark_verse", "false")
	_ = mngr.Update_setting(user.Username, "bookmark_translation", "false")
	user, _ = mngr.Get(username)

	assert.False(t, user.Setting.Surah_verse)
	assert.False(t, user.Setting.Surah_translation)
	assert.False(t, user.Setting.Bookmark_verse)
	assert.False(t, user.Setting.Bookmark_translation)

	_ = mngr.Update_setting(user.Username, "surah_verse", "true")
	_ = mngr.Update_setting(user.Username, "surah_translation", "true")
	_ = mngr.Update_setting(user.Username, "bookmark_verse", "true")
	_ = mngr.Update_setting(user.Username, "bookmark_translation", "true")
	user, _ = mngr.Get(username)

	assert.True(t, user.Setting.Surah_verse)
	assert.True(t, user.Setting.Surah_translation)
	assert.True(t, user.Setting.Bookmark_verse)
	assert.True(t, user.Setting.Bookmark_translation)
}
