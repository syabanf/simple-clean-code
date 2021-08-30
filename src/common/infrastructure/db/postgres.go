package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"golang.org/x/crypto/ssh"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Open(s string) (_ driver.Conn, err error) {
	return pq.DialOpen(self, s)
}

func (self *ViaSSHDialer) Dial(network, address string) (net.Conn, error) {
	return self.client.Dial(network, address)
}

func (self *ViaSSHDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return self.client.Dial(network, address)
}

// NewPgSQLConnection ...
func newPgSQLConnection() *ConnectTo {

	dbConn := ConnectTo{
		PGRead: DBReadPG(),
		DBRead: DBGormPG(),
		DBExec: DBGormPG(),
	}

	return &dbConn
}

var (
	open *sql.DB
	db   *sqlx.DB
	godb *gorm.DB
)

func init() {

	// env Load
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	var (
		dbUser = mustGetenv("PG_USER")
		dbPwd  = mustGetenv("PG_PASSWORD")
		dbHost = mustGetenv("PG_HOST")
		// dbPort  = mustGetenv("PG_PORT")
		dbName = mustGetenv("PG_DBNAME")
	)

	// initConn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbName, dbUser, dbPwd)
	initConn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPwd, dbHost, dbName)
	open, _ = sql.Open("postgres", initConn)

	// See "Important settings" section.
	open.SetConnMaxLifetime(time.Minute * 3)
	open.SetMaxOpenConns(10)
	open.SetMaxIdleConns(10)

}

func DBReadPG() *sqlx.DB {
	db = sqlx.NewDb(open, "postgres")

	if err := db.Ping(); err != nil {
		log.Println("error ping", err)
	}

	return db
}

func DBGormPG() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: open,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error connection postgres by gorm: ", err)
	}

	return db
}
