package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// Database driver choice
const (
	MySQLDriver int = iota
	PostgreDriver
	SQLiteDriver
	MSSQLDriver
)

type ConnectTo struct {
	// MySQL
	DBRead *gorm.DB
	DBExec *gorm.DB
	// PostgreSQL
	PGRead *sqlx.DB
}

// NewDBConnectionFactory to switch db driver at runtime
// TODO Pass DBConfig as parameter
func NewDBConnectionFactory(driver int) *ConnectTo {
	switch driver {
	case MySQLDriver:
		return newMySQLDBConnection()
	case PostgreDriver:
		return newPgSQLConnection()
	case SQLiteDriver:
		return nil
	default:
		log.Fatal("Invalid database driver")
		return nil
	}
}
