package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQLDBConnection ...
func newMySQLDBConnection() *ConnectTo {

	// env Load
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	db := dbReadConn()
	db2 := dbExecConn()

	dbConn := ConnectTo{
		DBRead: db,
		DBExec: db2,
	}

	return &dbConn
}

func dbReadConn() *gorm.DB {

	// Connect to database
	var (
		dbUser = mustGetenv("MYSQL_USER")
		dbPwd  = mustGetenv("MYSQL_PASSWORD")
		dbHost = mustGetenv("MYSQL_HOST")
		dbPort = mustGetenv("MYSQL_PORT")
		dbName = mustGetenv("MYSQL_DBNAME")
	)

	// socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	// if !isSet {
	// 	socketDir = "/cloudsql"
	// }

	var dsn string
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName)
	// dsn = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connection: ", err)
	}

	return db
}

func dbExecConn() *gorm.DB {

	// Connect to database
	var (
		dbUser = mustGetenv("MYSQL_USER")
		dbPwd  = mustGetenv("MYSQL_PASSWORD")
		dbHost = mustGetenv("MYSQL_HOST")
		dbPort = mustGetenv("MYSQL_PORT")
		dbName = mustGetenv("MYSQL_DBNAME")
	)

	// socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	// if !isSet {
	// 	socketDir = "/cloudsql"
	// }

	var dsn string
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName)
	// dsn = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connection: ", err)
	}

	return db
}

func dbReadTCP() *gorm.DB {

	// Connect to database
	var (
		dbUser = mustGetenv("MYSQL_USER")
		dbPwd  = mustGetenv("MYSQL_PASSWORD")
		dbHost = mustGetenv("MYSQL_HOST")
		dbPort = mustGetenv("MYSQL_PORT")
		dbName = mustGetenv("MYSQL_DBNAME")
	)

	var dsn string
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connection: ", err)
	}

	return db
}

func dbExecTCP() *gorm.DB {

	// Connect to database
	var (
		dbUser = mustGetenv("MYSQL_USER")
		dbPwd  = mustGetenv("MYSQL_PASSWORD")
		dbHost = mustGetenv("MYSQL_HOST")
		dbPort = mustGetenv("MYSQL_PORT")
		dbName = mustGetenv("MYSQL_DBNAME")
	)

	var dsn string
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connection: ", err)
	}

	return db
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
