# go-repository-template

go関連のrepositoryを作成する際にtemplateとして使用。postgresも追加済み

## 使用時TODO

- postgres認証情報修正
  - .devcontainer/.env
- DB初期化sql修正
  - .devcontainer/.init-db/init-tables.sql

## 関連コマンド

- appコンテナからdbコンテナのpostgresへ接続
  - `PGPASSWORD=postgres psql -h db -d postgres --username=postgres`
