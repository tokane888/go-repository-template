# go-repository-template

go関連のrepositoryを作成する際にtemplateとして使用

## 開発環境構築手順

- devcontainer起動
- 下記実行でcommit前git hook登録
  - `lefthook install`

## 設計方針

- goによるbackend構築目的に使用
- ディレクトリ構成は下記に従う
  - [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ja.md#standard-go-project-layout)に従って作成
- 設計はクリーンアーキに従う

## テンプレ使用時のTODO

- devcontainerを使用しない場合
  - .devcontainer ディレクトリ削除
- `cmd/app/`の`app`の部分をプロジェクトに応じて調整
- リポジトリ内から"TODO: "を検索し、修正

## 今後のTODO

- buildしたバイナリをECRへuploadするCI追加
