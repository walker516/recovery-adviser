package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config 構造体
type Config struct {
	Database struct {
		Type string `json:"type"`
		DSN  string `json:"dsn"`
	} `json:"database"`
	SQLQueryPath string
}

// グローバル変数としての設定
var ConfigData Config

// 設定ファイルの初期化関数
func InitConfig() {
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		log.Fatal("CONFIG_FILE_PATH environment variable is not set")
	}

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	err = json.Unmarshal(data, &ConfigData)
	if err != nil {
		log.Fatalf("Error unmarshaling config data: %s", err)
	}

	// 環境変数からSQLクエリのパスを取得
	ConfigData.SQLQueryPath = os.Getenv("SQL_QUERY_PATH")
	if ConfigData.SQLQueryPath == "" {
		log.Fatal("SQL_QUERY_PATH environment variable is not set")
	}

	fmt.Printf("Config loaded from %s\n", configFilePath)
	fmt.Printf("SQL queries will be loaded from %s\n", ConfigData.SQLQueryPath)
}
