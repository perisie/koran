package main

import (
	"flag"
	"github.com/arikama/koran-backend/daos"
	"github.com/arikama/koran-backend/favorite"
	"github.com/arikama/koran-backend/managers"
	ari_mouse "github.com/arikama/koran-backend/mouse"
	"github.com/arikama/koran-backend/routes"
	"github.com/arikama/koran-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hooligram/kifu"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		kifu.Warn(".env: %v", err.Error())
	}
	wd, err := os.Getwd()
	if err != nil {
		kifu.Fatal("error getting wd: %v", err.Error())
	}
	kifu.Info("wd: %v", wd)

	_ = ari_mouse.Mouse_new(wd)

	var quranManager managers.QuranManager
	quranManager, err = managers.NewQuranManagerImpl("./qurancsv")
	if err != nil {
		kifu.Fatal("error initializing quran manager: %v", err.Error())
	}

	var googleAuthService services.GoogleAuthService
	googleAuthService, err = services.NewGoogleAuthServiceImpl()
	if err != nil {
		kifu.Fatal("error initializing google auth service: %v", err.Error())
	}

	userDao, err := daos.NewUserDaoImpl()
	if err != nil {
		kifu.Fatal("error initializing user dao: %v", err.Error())
	}

	var userManager managers.UserManager
	userManager, err = managers.NewUserManagerImpl(userDao)
	if err != nil {
		kifu.Fatal("error initializing user manager: %v", err.Error())
	}

	favDao, err := favorite.NewFavDaoImpl()
	if err != nil {
		kifu.Fatal("error initializing fav dao: %v", err.Error())
	}

	var favManager favorite.FavManager
	favManager, err = favorite.NewFavManagerImpl(favDao, userDao)
	if err != nil {
		kifu.Fatal("error initializing fav manager: %v", err.Error())
	}

	s := setupWebServer()
	routes.Routes(s, quranManager, googleAuthService, userManager, favManager)

	if isTestEnv() {
		go s.Run()
	} else {
		s.Run()
	}
}

func setupWebServer() *gin.Engine {
	r := gin.Default()
	configureCors(r)
	return r
}

func configureCors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"x-access-token"}
	r.Use(cors.New(config))
}

func isTestEnv() bool {
	return flag.Lookup("test.v") != nil
}
