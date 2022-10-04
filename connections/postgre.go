package connections

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbname, port)
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dbConn, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Println(connString)
		log.Println(err)
		return dbConn, err
	}

	if appEnv != "prod" {
		dbConn = dbConn.Debug()
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Println(sqlDB)
		return dbConn, err
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	log.Println("postgre connected!")
	return dbConn, nil
}
