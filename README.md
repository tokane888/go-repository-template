# go-repository-template

Go モノレポ構成のテンプレートリポジトリ

## ディレクトリ構成

```sh
.
├── services/          # サービス群
│   └── sample/       # サンプルサービス
│       ├── cmd/
│       ├── .env/
│       └── go.mod
├── pkg/              # 共通モジュール
│   └── logger/       # ログ関連
│       └── go.mod
└── README.md
```

## 開発環境構築手順

- devcontainer起動
- 下記実行でcommit前git hook登録
  - `lefthook install`

## 設計方針

- ディレクトリ構成は[Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ja.md#standard-go-project-layout)に従う
- Go モノレポによる複数サービス管理
- 各サービスは独立した go.mod を持つ
- 共通モジュールは `pkg/` ディレクトリに配置
  - replace ディレクティブでローカル参照
- 設計はクリーンアーキテクチャに従う

## テンプレ使用時のTODO

- devcontainerを使用しない場合
  - .devcontainer ディレクトリ削除
- `services/sample/` を実際のサービス名に変更
- 新しいサービス追加時は `services/` 配下に作成
- リポジトリ内から"TODO: "を検索し、修正
- リポジトリ内から"go-repository-template"を検索し、修正
- CLAUDE.mdは削除の上claude内で`/init`で再生成して調整
- claude codeを使用しない場合は下記で関連ファイルを探索して削除
  - `find . -name '*claude*' -not -path './.git/*'`
- services配下の不要なservice, README.mdは適宜削除
- claude codeによるレビューを可能にする場合、`claude`コマンド実行後、下記でgithub appをインストール
  - `/install-github-app`
    - 詳細は[公式](https://docs.anthropic.com/en/docs/claude-code/github-actions)参照

## サービス実行例

```bash
# サンプルサービスの実行
cd services/sample
go run ./cmd/sample
```

## サービスデバッグ実行例

- ctrl+shift+dで"RUN AND DEBUG"メニューを開く
- 上のメニューからデバッグ実行したいserviceを選択
- F5押下でデバッグ実行
