package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"recovery-adviser-api/config"
	"recovery-adviser-api/infrastructure"
	"recovery-adviser-api/router"
)

func main() {
	// 設定ファイルの初期化
	config.InitConfig()

	// ログファイルの設定
	setUpLogging()

	// データベース接続の確立
	db, err := infrastructure.OpenDatabaseConnection()
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	}
	defer db.Close()

	// ルータの初期化
	e := router.NewRouter(db)

	// サーバー開始
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("Shutting down the server")
	}
}

// 日付ごとに新しいログファイルを作成
func setUpLogging() {
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			log.Fatalf("Failed to create log directory: %s", err)
		}
	}

	logFile := fmt.Sprintf("%s/%s.log", logDir, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Logging initialized")
}
