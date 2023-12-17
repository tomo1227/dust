# dust

## 概要

このリポジトリで作成するアプリケーションの目標はまだ定まっていない。
一番の目的はAWSやGolangのプラクティスである。

## usage

1. 下記のアプリケーションをインストールしてください。

- Visual Studio Code
  - devcontainer
- Docker
- Docker-compose
- Git

2. devcontainerで起動。


## 動作確認

**devcontainer** 内で行う。

### サーバ起動

airでホットリロード可能なサーバを起動する。

```
cd cmd/api
air
```

curl で以下のように確認する。

```
$ curl localhost:8080/health
{"status":"OK!"}
```

> [!TIP]
> **ローカルでの確認方法**
> http://localhost/health にアクセスする。
> `{"status":"OK!"}`と表示されることを確認。

## Go のオニオンアーキテクチャディレクトリ構成例（gRPC、REST、GraphQL）

以下は、オニオンアーキテクチャに基づく Go プロジェクトに gRPC、REST、および GraphQL の実装を含めたディレクトリ構成の例です。(**一部未実装**)

```
.
├── cmd             # アプリケーションのエントリーポイントが含まれます。
│   ├── api         # APIサーバーのエントリーポイント。
│   │   └── main.go
│   └── batch       # バッチ処理のエントリーポイント。
│       └── main.go
├── go.mod          # プロジェクトの依存関係を管理します。
├── internal        # アプリケーションコードの主要部分。
│   ├── adapter
│   │   ├── rest       # RESTful APIエンドポイント（ハンドラー）が含まれます。これらはHTTPリクエストを受け取り、適切なドメインサービスを呼び出し、HTTPレスポンスを生成します。
│   │   └── batch     # バッチ処理ジョブやスケジュールされたタスクが含まれます。これらは定期的に実行される、または大量のデータを処理するためのジョブです。
│   ├── domain
│   │   ├── model       # ドメインモデルが含まれます。これらはビジネスロジックの核心的な部分を表現するためのエンティティや値オブジェクトです。
│   │   ├── repository  # リポジトリインターフェースが定義されています。これは永続化層（通常はデータベース）へのアクセスを抽象化します。
│   │   └── service    # システム外部との統合や外部APIの呼び出しを抽象化します。
│   ├── infrastructure  # インフラストラクチャ層は、技術的な詳細（データベース、メッセージングなど）が含まれています。。
│   ├── injector    # 依存性注入設定が含まれています。
│   ├── shared      # 共有ユーティリティやヘルパー関数が含まれています。
│   └── usecase     # ユースケース層は、特定のビジネスケースを実行するための操作を含む。
├── loadtest        # ロードテスト用のコードが含まれています。
│   └── main.go
└── schema
    ├── oapi       # OpenAPI (以前のSwagger) スキーマが含まれます。これらはRESTful APIのエンドポイントとリクエスト/レスポンスフォーマットを定義します。
    └── graphql   # GraphQLスキーマが含まれます。これらはGraphQL APIのクエリ、ミューテーション、およびサブスクリプションを定義します。
```

この構成において、以下のディレクトリがそれぞれの API 実装を含んでいます。

- `infrastructure/adapter/http/`：REST API 関連の実装。
- `infrastructure/adapter/grpc/`：gRPC 関連の実装。
- `infrastructure/adapter/graphql/`：GraphQL 関連の実装。

このディレクトリ構成は、Go 言語でのオニオンアーキテクチャと gRPC、REST、および GraphQL の実装を組み合わせた一例です。プロジェクトの要件やチームの好みに応じて、適切にディレクトリ構成を調整してください。

https://github.com/golang-standards/project-layout を参考

## merge

main にマージするときは`squash and merge`すること

## commit template の設定

```bash
git config --global commit.template .commit.template
```

## Commit Message ガイドライン

.commit template を参考に記述してください

## Pull Request ガイドライン

`[emoji][type]([scope]): (title)`

- ✨feat: ログイン機能を追加
- 👓fix(a11y): ナビゲーションのアクセシビリティを改善

release-please による自動リリースを行うため、[Conventional Commit messages\.](https://www.conventionalcommits.org/) を採用しています。

最も重要な接頭辞として、覚えておいていただきたいのが、この接頭辞です：

- `fix:` バグの修正を表し、[SemVer](https://semver.org/)パッチと関連します。
- `feat：`新機能を表し、[SemVer](https://semver.org/)のマイナーに対応します。
- `feat!:`、`fix!:`、`refactor!:`などは、破壊的な変更（「`!`」で示す）を表し、[SemVer](https://semver.org/)のメジャーになります。

`fix:`や`feat:`以外の型も可能で、例えば@commitlint/config-conventional（Angular の規約に基づく）は、`build:`, `chore:`, `ci:`, `docs:`, `style:`, `refactor:`, `perf:`, `test:` などを推奨します。

### Pull Request Type:

- fix: 🐛 バグの修正 (SemVer パッチと関連)
- feat: ✨ 新機能を追加 (SemVer のマイナーに対応)
- feat!: 💥 破壊的な新機能 (SemVer のメジャーになります)
- fix!: 💥 破壊的なバグ修正 (SemVer のメジャーになります)
- refactor: ♻️ コードの再構築
- revert: ⏪ 変更を取り消す
- test: 🧪 テストに関連する変更
- docs: 📚 ドキュメンテーションの変更
- style: 🎨 スタイルや書式の変更
- perf: ⚡ パフォーマンス改善
- build: 👷‍♀️ ビルドシステムや外部依存関係の変更
- chore: 🔧 その他の変更
- ci: 🎡 CI/CD パイプラインに関連する変更

### Pull Request Scope:

プロジェクト内の変更の範囲を示すことができます。
変更の範囲が広い場合は、スコープを省略しても構いません。
- ui: ユーザーインターフェースに関する変更
- ux: ユーザーエクスペリエンスに関する変更
- api: API に関連する変更
- db: データベースに関連する変更
- auth: 認証やセキュリティに関連する変更
- i18n: 国際化に関連する変更
- a11y: アクセシビリティに関連する変更
- server: サーバーサイドの変更
- client: クライアントサイドの変更
- infra: インフラストラクチャに関連する変更
- test: テストに関連する変更
- config: 設定ファイルに関連する変更
- deps: 依存関係に関連する変更
- ci: CI/CD パイプラインに関連する変更
