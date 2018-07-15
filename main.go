package main

import (
	"database/sql"
	"fmt"
	httpDeliver "go-api/book/delivery/http"
	_bookRepo "go-api/book/repository"
	bookUcase "go-api/book/usecase"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.yml`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(nil)
	}
	if viper.GetBool(`debug`) {
		fmt.Println("Server RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	// val.Add("useSSL", "false")
	//	val.Add("allowPublicKeyRetrieval", "true")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	// golang build in
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer fmt.Println("dbConn is close")
	defer dbConn.Close()

	bookRepo := _bookRepo.NewMysqlBookRepository(dbConn)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	bu := bookUcase.NewBookUsecase(bookRepo, timeoutContext)

	r := gin.Default()
	httpDeliver.NewBookHttpHandler(r, bu)

	r.Run(":" + viper.GetString("server.address"))
}
