package controllers

import (
	"net/http"
	"strconv"

	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
)

func GetSurahVerseController(quranManager managers.QuranManager) func(c *gin.Context) {
	return func(c *gin.Context) {
		surahId, err := strconv.Atoi(c.Param("surah_id"))
		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}
		verseId, err := strconv.Atoi(c.Param("verse_id"))
		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}
		verse, err := quranManager.GetVerse(surahId, verseId)
		if err != nil {
			utils.JsonError(c, http.StatusNotFound, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"verse":        verse.Text,
				"translations": verse.Translations,
			},
		})
	}
}
