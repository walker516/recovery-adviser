package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"recovery-adviser-api/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/sijms/go-ora/v2"
)

// データベース接続を開く関数
func OpenDatabaseConnection() (*sql.DB, error) {
	var db *sql.DB
	var err error

	if config.ConfigData.Database.Type == "mysql" {
		db, err = sql.Open("mysql", config.ConfigData.Database.DSN)
	} else if config.ConfigData.Database.Type == "oracle" {
		db, err = sql.Open("oracle", config.ConfigData.Database.DSN)
	} else {
		return nil, fmt.Errorf("Unsupported database type: %s", config.ConfigData.Database.Type)
	}

	if err != nil {
		log.Printf("Error opening database: %s", err)
		return nil, err
	}
	return db, nil
}
