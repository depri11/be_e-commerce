package connections

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func PostgreConnect(host, user, pass, dbname, port string) (dbConn *sql.DB, err error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	dbConn, err = sql.Open("postgres", connString)
	if err != nil {
		log.Println(connString)
		log.Println(err)
		return dbConn, err
	}

	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)
	log.Println("postgre connected!")
	return dbConn, nil
}

func GormConnect(host, user, pass, dbname, port, appEnv string) (dbConn *gorm.DB, err error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	dbConn, err = gorm.Open("postgres", connString)
	if err != nil {
		log.Println(connString)
		log.Println(err)
		return dbConn, err
	}

	if appEnv != "prod" {
		dbConn = dbConn.Debug()
	}

	dbConn.DB().SetConnMaxLifetime(time.Minute * 3)
	dbConn.DB().SetMaxOpenConns(10)
	dbConn.DB().SetMaxIdleConns(10)
	log.Println("postgre connected!")
	return dbConn, nil
}
