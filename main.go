package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

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
	var quranManager managers.QuranManager
	quranManager, err = InitializeQuranManagerImpl("./qurancsvs")
	if err != nil {
		kifu.Fatal("error initializing quran manager: %v", err.Error())
	}
	s := setupWebServer(quranManager)
	if isTestEnv() {
		go s.Run()
	} else {
		s.Run()
	}
}

func setupWebServer(quranManager managers.QuranManager) *gin.Engine {
	r := gin.Default()
	configureCors(r)
	routes(r, quranManager)
	return r
}

func configureCors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
}

func getDb() (*sql.DB, error) {
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlIp := os.Getenv("MYSQL_IP")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	if isTestEnv() {
		result, _ := mysqltestcontainer.Start("test")
		mysqlUsername = result.Username
		mysqlPassword = result.Password
		mysqlIp = result.Ip
		mysqlPort = result.Port
		mysqlDatabase = result.Database
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
