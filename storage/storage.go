package storage

import (
	"sync"
	"todo-api/config"
	"todo-api/helper"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var once sync.Once

// Connect database
func Connect() *gorm.DB {
	once.Do(func() {
		conf := config.GetConnection()
		var err error
		db, err = gorm.Open(conf.Driver, conf.DataSource)
		if err != nil {
			helper.Logging("Storage", "Connect", err.Error())
		}
		db.LogMode(conf.LogMode)
	})

	return db
}
