# database migrate 機能

## ディレクトリ構成

```
.
├── Makefile         # migration 実行用
├── migrate          # migration 実行用
├── README.md        # this file
├── sample           # database 毎の設定
│   ├── config.toml # database 接続設定
│   └── migrations   # migration ファイル
│       └── 20230313033911-create_users.sql
└── goose.sh   # migration 実行コマンド
```

## migration の一括実行

```bash
make migrate
```

## migration コマンド

### status

migration の実行状態を出力

### new

migration ファイルの生成

### up

migration の実行

> dryrun = yes で実行計画の表示

### down limit

`limit`で指定した数値分のバージョンを下げる

> 0 を指定すると初期化
> dryrun = yes で実行計画の表示

### redo

前回のマイグレーションを再適用する

## migration 設定方法

### database 設定

`database名/dbconfig.yml` を作成

### migration ファイルの生成

1. スクリプトの実行

   ```bash
   sh sql-migrate.sh
   ```

1. database の選択

   `database名` を入力する

1. environment の選択

   `database名/dbconfig.yml` で設定した環境を指定する

1. command の選択

   `new` を入力する

1. DSL を記述

   `-- +migrate Up` : migrate up した際の記述

   `-- +migrate Down` : migrate down した際の記述