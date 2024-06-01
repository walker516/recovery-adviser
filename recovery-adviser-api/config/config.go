package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config 構造体
type Config struct {
	Database struct {
		Type string `json:"type"`
		DSN  string `json:"dsn"`
	} `json:"database"`
}

// グローバル変数としての設定
var ConfigData Config

// 設定ファイルの初期化関数
func InitConfig() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	err = json.Unmarshal(data, &ConfigData)
	if err != nil {
		log.Fatalf("Error unmarshaling config data: %s", err)
	}
}
