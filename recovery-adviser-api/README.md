# RecoveryAdviserAPI

recovery-adviser-api は、部品情報とジョブ情報を管理する Go 言語のアプリケーションです。この README では、プロジェクトのセットアップ手順とディレクトリ構造を説明します。

## ディレクトリ構造

```
recovery-adviser-api/
├── README.md
├── cmd
│   └── server
│       └── main.go
├── config
│   └── config.go
├── config.json
├── domain
│   ├── job.go
│   ├── part.go
│   └── repository.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── database.go
│   ├── repository
│   │   ├── job_repository.go
│   │   └── part_repository.go
│   └── sql
│       ├── mysql
│       │   ├── job_queries.sql
│       │   └── part_queries.sql
│       └── oracle
│           ├── job_queries.sql
│           └── part_queries.sql
├── interface
│   ├── handler
│   │   ├── job_handler.go
│   │   └── part_handler.go
│   └── response
├── router
│   └── router.go
└── usecase
    ├── job_usecase.go
    └── part_usecase.go
```

## セットアップ手順

### 前提条件

- Go 1.16 以上がインストールされていること
- MySQL または Oracle データベースがセットアップされていること

### プロジェクトのクローン

```sh
git clone https://github.com/walker516/recovery-adviser/recovery-adviser-api.git
cd recovery-adviser-api
```

### 依存関係のインストール

```sh
go mod tidy
```

### 設定ファイルの作成

`config.json`ファイルを作成し、データベース接続情報を設定します。例:

```json
{
  "database": {
    "type": "oracle",
    "dsn": "oracle://username:password@hostname:port/sid"
  }
}
```

### データベーススキーマの準備

`sql/mysql`または`sql/oracle`ディレクトリ内の SQL スクリプトを使用して、データベーススキーマを作成します。

### サーバーの起動

以下のコマンドでサーバーを起動します。

```sh
go run cmd/server/main.go
```

サーバーはデフォルトで`http://localhost:8080`でリッスンします。

### エンドポイント

- `GET /part/:seppenbuban`: 部品情報を取得します。
- `GET /recovery-job-status/:seppenbuban`: リカバリージョブのステータスを取得します。
- `GET /job-queue/:process_order`: ジョブキューを取得します。
- `PUT /job-queue/:process_order`: ジョブキューを更新します。
- `GET /job-lock/:process_order`: ジョブロックを取得します。
- `DELETE /job-lock/:process_order`: ジョブロックを削除します。

## 使用技術

- Go
- Echo (Web フレームワーク)
- MySQL または Oracle データベース
