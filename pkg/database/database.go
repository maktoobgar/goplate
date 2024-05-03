package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type (
	Database struct {
		Type     string `yaml:"type"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"db_name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"ssl_mode"`
		TimeZone string `yaml:"time_zone"`
		Charset  string `yaml:"charset"`
	}

	RelationalDatabaseFunction func() (*sql.DB, error)
)

// creates connections and returns connections and main or test database error if anything wrong happened
func New(db Database, debug bool) (RelationalDatabaseFunction, error) {
	connectionCreatorFunction := func(dbType, config string) RelationalDatabaseFunction {
		return func() (*sql.DB, error) {
			c, err := sql.Open(dbType, config)
			if err != nil {
				return nil, err
			}
			return c, nil
		}
	}

	config := ""

	switch strings.ToLower(db.Type) {
	case "mysql":
		config = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.Username, db.Password, db.Host, db.Port, db.DbName)
	case "sqlite3":
		if _, err := os.Stat(db.DbName); err != nil {
			_, err = os.Create(db.DbName)
			if err != nil {
				return nil, err
			}
		}
		config = fmt.Sprintf("file:%s?cache=shared&mode=rw", db.DbName)
	case "postgres":
		config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", db.Host, db.Port, db.Username, db.Password, db.DbName, db.SSLMode)
	case "mssql":
		config = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", db.Host, db.Username, db.Password, db.Port, db.DbName)
	default:
		log.Panicf("db: unrecognizable database type `%s`", db.Type)
	}

	dbFunc := connectionCreatorFunction(db.Type, config)

	return dbFunc, nil
}
