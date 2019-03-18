package config

import (
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql" // mysql connect
	"github.com/joho/godotenv"
)

// Connection ...
type Connection struct {
	Driver          string
	DBConnectString string
}

var once sync.Once
var pkgConnection *Connection

// GetConnection ...
func GetConnection() *Connection {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		driveName := os.Getenv("DB_CONNECTION")
		database := os.Getenv("DB_DATABASE")
		userName := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		dataSource := userName + ":" + password + "@/" + database + "?charset=utf8&parseTime=True&loc=Local"

		pkgConnection = &Connection{
			Driver:          driveName,
			DBConnectString: dataSource,
		}
	})

	return pkgConnection
}
