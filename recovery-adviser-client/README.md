# RecoveryAdviserClient

RecoveryAdviserClient は、部品表（K6）と CAD（V6）の構成アンマッチをリカバリーするツールです。

## 機能

- 部品情報の取得
- リカバリーツールの実行
- BOM および CAABatch の更新
- ジョブキューの管理
- ログファイルからのパラメータ抽出

## フォルダ構成

```

recovery-adviser-client/
│
├── dist/
│ ├── RecoveryAdviser.exe
│ ├── bomRecoveryCheckTool.exe
│ └── RecoveryAdviser.zip
│
├── scripts/
│ └── create_zip.py
│
├── utils/
│ ├── api.py
│ ├── log.py
│ ├── prompts.py
│ └── update.py
│
├── docker/
│ ├── Dockerfile
│ └── docker-compose.yml
│
├── main.py
├── recovery_adviser.spec
├── requirements.txt
├── bomRecoveryCheckTool.exe
└── README.md

```

## 必要条件

- `bomRecoveryCheckTool.exe` は同包されています。

## インストール

1. `RecoveryAdviser.zip`をダウンロードし、任意のフォルダに解凍してください。

## 使用方法

コマンドプロンプトまたはターミナルを開き、以下のコマンドを実行してください。

```sh
RecoveryAdviser.exe <SEPPENBUBAN>
```

例:

```sh
RecoveryAdviser.exe 123456
```

これにより、指定された部品番号に対してリカバリーツールが実行されます。

## 開発者向け情報

### Docker 環境での開発

開発環境は Docker コンテナ内でセットアップされます。

#### 開発環境のコンテナ起動手順

1. **Network 作成**：既にあるなら不要

```sh
docker network create recovery-net
```

1. **ビルド**：以下のコマンドを実行して Docker イメージをビルドします。

   ```sh
   make up
   ```

2. **コンテナ内での作業**：以下のコマンドを実行してコンテナ内に入ります。

   ```sh
   make exec
   ```

3. **EXE ファイルの作成**：以下のコマンドを実行して EXE ファイルを作成します。

   ```sh
   make create-exe
   ```

4. **ZIP ファイルの作成**：以下のコマンドを実行して配布用の ZIP ファイルを作成します。
   ```sh
   make create-zip
   ```
