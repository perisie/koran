package controllers

import (
	"net/http"
	"strconv"

	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
)

func GetSurahController(quranManager managers.QuranManager) func(c *gin.Context) {
	return func(c *gin.Context) {
		surahId, err := strconv.Atoi(c.Param("surah_id"))
		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}
		surah, err := quranManager.GetSurah(surahId)
		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"surah": surah,
			},
		})
	}
}
