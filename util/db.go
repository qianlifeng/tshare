package util

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path"
	"sync"
)

var (
	dbOnce       sync.Once
	dbConnection *gorm.DB
)

func getDBPath() string {
	return path.Join(GetAppFolder(), "csf.db")
}

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		conn, err := gorm.Open(sqlite.Open(getDBPath()+"?_journal_mode=WAL&_synchronous=OFF"), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("Can't connect to db , path=%s, err=%s ", getDBPath(), err.Error()))
		}
		dbConnection = conn
	})

	return dbConnection
}