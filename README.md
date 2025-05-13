# go-repository-template

go関連のrepositoryを作成する際にtemplateとして使用

## 設計方針

- ディレクトリ構成
  - [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ja.md#standard-go-project-layout)に従って作成

## テンプレ使用時のTODO

- devcontainerを使用しない場合
  - .devcontainer ディレクトリ削除
- devcontainerを使用する場合
  - .vscode ディレクトリ削除
- `cmd/app/`の`app`の部分をプロジェクトに応じて調整
- リポジトリ内から"TODO: "を検索し、修正
