package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/arikama/koran-backend/managers"
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

	db, err := NewDb()
	if err != nil {
		kifu.Fatal("error connecting to db: %v", err.Error())
	}

	arctictern.Migrate(db, "./migrations")

	var quranManager managers.QuranManager
	quranManager, err = wireQuranManagerImpl("./qurancsv")
	if err != nil {
		kifu.Fatal("error initializing quran manager: %v", err.Error())
	}

	var googleAuthManager managers.GoogleAuthManager
	googleAuthManager, err = wireGoogleAuthManagerImpl()
	if err != nil {
		kifu.Fatal("error initializing google auth manager: %v", err.Error())
	}

	s := setupWebServer()
	routes(s, quranManager, googleAuthManager)

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
