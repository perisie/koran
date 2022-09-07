package controllers

import (
	"net/http"
	"strconv"

	"github.com/arikama/koran-backend/managers"
	"github.com/gin-gonic/gin"
)

func GetSurahVerseController(quranManager managers.QuranManager) func(c *gin.Context) {
	return func(c *gin.Context) {
		surahId, err := strconv.Atoi(c.Param("surah_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		verseId, err := strconv.Atoi(c.Param("verse_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		verse, err := quranManager.GetVerse(surahId, verseId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
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
