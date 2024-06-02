# RecoveryAdviserAPI

recovery-adviser-api は、部品情報とジョブ情報を管理する。

## ディレクトリ構造

```
recovery-adviser-api/
├── Makefile
├── README.md
├── cmd
│   └── server
│       └── main.go
├── config
│   └── config.go
├── config.dev.json
├── config.prod.json
├── docker
│   ├── Dockerfile.dev
│   ├── Dockerfile.prod
│   ├── docker
│   │   └── mysql
│   │       └── init
│   ├── docker-compose.dev.yml
│   ├── docker-compose.prod.yml
│   └── mysql
│       └── init
│           ├── 01_create_tables.sql
│           └── 02_insert_sample_data.sql
├── domain
│   ├── job.go
│   ├── part.go
│   └── repository.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── database.go
│   └── repository
│       ├── job_repository.go
│       └── part_repository.go
├── interface
│   └── handler
│       ├── job_handler.go
│       ├── part_handler.go
│       └── sysdate_handler.go
├── router
│   └── router.go
├── sql
│   ├── mysql
│   │   ├── job_queries.sql
│   │   └── part_queries.sql
│   └── oracle
│       ├── job_queries.sql
│       └── part_queries.sql
└── usecase
    ├── job_usecase.go
    └── part_usecase.go
```

## セットアップ手順

### 前提条件

- Go 1.16 以上がインストールされていること
- MySQL または Oracle データベースがセットアップされていること
- Docker および Docker Compose がインストールされていること

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

環境ごとに適切な `config.json` ファイルを作成し、データベース接続情報を設定します。

- 開発環境用: `config.dev.json`
- 本番環境用: `config.prod.json`

例:

```json
{
  "database": {
    "type": "oracle",
    "dsn": "user:password@oracle-host:1521/sid"
  }
}
```

### データベーススキーマの準備

`docker/mysql/init` ディレクトリ内の SQL スクリプトを使用して、データベーススキーマを作成します。

### Docker を使用した環境のセットアップ

#### 開発環境

開発環境用のコンテナを起動します。

```sh
make up
```

データベースのマイグレーションを実行します。

```sh
make migrate
```

サンプルデータを挿入します。

```sh
make seed
```

環境を停止する場合は以下のコマンドを使用します。

```sh
make down
```

#### 本番環境

本番環境用のコンテナを起動します。

```sh
make up-prod
```

環境を停止する場合は以下のコマンドを使用します。

```sh
make down-prod
```

### サーバーの起動

以下のコマンドでサーバーを起動します。

```sh
go run cmd/server/main.go
```

サーバーはデフォルトで`http://localhost:8080`でリッスンします。

### エンドポイント

- `GET /part/:seppenbuban`: 部品情報を取得します。

  ```sh
  curl -X GET "http://localhost:8080/part/SEPPEN001"
  ```

- `GET /recovery-job-status/:seppenbuban`: リカバリージョブのステータスを取得します。

  ```sh
  curl -X GET "http://localhost:8080/recovery-job-status/SEPPEN001"
  ```

- `GET /job-queue/:process_order`: ジョブキューを取得します。

  ```sh
  curl -X GET "http://localhost:8080/job-queue/PO123"
  ```

- `GET /job-queue/:process_order`: ジョブキューを取得します。

  ```sh
  curl -X GET "http://localhost:8080/job-queue/P0123"
  ```

- `GET /job-queue?seppenbuban=:seppenbuban`: 部品番号でジョブキューを取得します。

  ```sh
  curl -X GET "http://localhost:8080/job-queue?seppenbuban=SEPPEN001"
  ```

- `PUT /job-queue/:process_order`: ジョブキューを更新します。

  ```sh
  curl -X PUT "http://localhost:8080/job-queue/PO123" -H "Content-Type: application/json" -d '{"status": "completed", "host": "host2"}'
  ```

- `GET /job-lock/:process_order`: ジョブロックを取得します。

  ```sh
  curl -X GET "http://localhost:8080/job-lock/PO123"
  ```

- `DELETE /job-lock/:process_order`: ジョブロックを削除します。
  ```sh
  curl -X DELETE "http://localhost:8080/job-lock/PO123"
  ```

## 使用技術

- Go
- Echo (Web フレームワーク)
- MySQL または Oracle データベース
- Docker & Docker Compose
