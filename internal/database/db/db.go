package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	URL           string
	MaxDBConn     int
	EnableTracing bool
}

// Init the connection to DB
func Init(config *Config) error {
	if DB == nil {
		sqlDB, err := sql.Open("postgres", config.URL)
		if err != nil {
			log.Println("Unable to open postges connection. Err:", err)
			return err
		}
		//fetching max connections and lifetime and all the other config variables from dev.json and mapping it with config
		sqlDB.SetMaxIdleConns(config.MaxDBConn)
		sqlDB.SetMaxOpenConns(config.MaxDBConn)
		sqlDB.SetConnMaxLifetime(time.Hour)

		DB, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			log.Println("Unable to open postges gorm connection. Err:", err)
			return err
		}

		log.Println("Successfully established database connection")
	}

	return nil
}

type DBConn struct {
	*gorm.DB
}

func New() *DBConn {
	return &DBConn{
		DB: DB,
	}
}

// not needed as of now , might require in furture toh
// adding args which will be used in sql queries
type Args struct {
	cnt  int
	vals []interface{}
}

func (a *Args) Append(v interface{}) string {
	a.cnt++
	a.vals = append(a.vals, v)
	return "?"
}

func (a *Args) Vals() []interface{} {
	return a.vals
}
