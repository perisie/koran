package controllers

import (
	"net/http"

	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
)

func GetRootController(quranManager managers.QuranManager) func(c *gin.Context) {
	return func(c *gin.Context) {
		surahInfos, err := quranManager.GetSurahInfos()
		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}
		utils.JsonData(c, http.StatusOK, gin.H{
			"surah_infos": surahInfos,
		})
	}
}
