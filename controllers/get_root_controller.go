package controllers

import (
	"net/http"

	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
)

func GetRootController(quranManager managers.QuranManager) func(c *gin.Context) {
	return func(c *gin.Context) {
		surahInfos, _ := quranManager.GetSurahInfos()
		utils.JsonData(c, http.StatusOK, gin.H{
			"surah_infos": surahInfos,
		})
	}
}
