package main

import (
	"log"
	"net/http"

	"recovery-adviser-api/config"
	"recovery-adviser-api/infrastructure"
	"recovery-adviser-api/router"
)

func main() {
	// 設定ファイルの初期化
	config.InitConfig()

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
