package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/arikama/koran-backend/daos"
	"os"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/arikama/koran-backend/favorite"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/routes"
	"github.com/arikama/koran-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hooligram/kifu"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		kifu.Warn(".env: %v", err.Error())
	}

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

func NewDb() (*sql.DB, error) {
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlIp := os.Getenv("MYSQL_IP")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	if isTestEnv() {
		result, _ := mysqltestcontainer.Create("test")
		mysqlUsername = result.GetDbInfo().Username
		mysqlPassword = result.GetDbInfo().Password
		mysqlIp = result.GetDbInfo().Ip
		mysqlPort = result.GetDbInfo().Port
		mysqlDatabase = result.GetDbInfo().DbName
	}

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", mysqlUsername, mysqlPassword, mysqlIp, mysqlPort, mysqlDatabase)
	dataSourceName += "?charset=utf8mb4"
	dataSourceName += "&collation=utf8mb4_unicode_ci"
	dataSourceName += "&parseTime=true"

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func isTestEnv() bool {
	return flag.Lookup("test.v") != nil
}
