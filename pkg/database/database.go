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
func New(dbs map[string]Database, debug bool) (cons map[string]RelationalDatabaseFunction, db RelationalDatabaseFunction, err error) {
	cons = map[string]RelationalDatabaseFunction{}
	mainOrTest := "test"
	if !debug {
		mainOrTest = "main"
	}

	connectionCreatorFunction := func(dbType, config string) RelationalDatabaseFunction {
		return func() (*sql.DB, error) {
			var c *sql.DB
			c, err = sql.Open(dbType, config)
			if err != nil {
				return nil, err
			}
			return c, nil
		}
	}

	for k, v := range dbs {
		config := ""

		switch strings.ToLower(v.Type) {
		case "mysql":
			config = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", v.Username, v.Password, v.Host, v.Port, v.DbName)
		case "sqlite3":
			if _, err = os.Stat(v.DbName); err != nil {
				_, err = os.Create(v.DbName)
				if err != nil {
					return
				}
			}
			config = fmt.Sprintf("file:%s?cache=shared&mode=rw", v.DbName)
		case "postgres":
			config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", v.Host, v.Port, v.Username, v.Password, v.DbName, v.SSLMode)
		case "mssql":
			config = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", v.Host, v.Username, v.Password, v.Port, v.DbName)
		default:
			log.Fatalf("db: unrecognizable database type `%s`", v.Type)
		}

		dbFunction := connectionCreatorFunction(v.Type, config)
		if mainOrTest == k {
			db = dbFunction
		}

		cons[k] = dbFunction
	}

	return
}

func CloseDBs(cons map[string]*sql.DB) {
	for _, con := range cons {
		con.Close()
	}
}
