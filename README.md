# Github Project Promoter (GHPP)

GitHub Projects をベースにした、プロジェクト進行ワークフローを支援する CLI ツール。

## 機能

- GitHub Projects の Item(Issue) をフェッチしてローカルにキャッシュ
- 定義済みの昇格ルールに基づいて Issue のステータスを自動 Promote
  - **計画フェーズ**: `inbox` → `plan`（上限数付き）
  - **実行フェーズ**: `ready` → `doing`（リポジトリ単位で1つまで）

## セットアップ

```bash
# ビルド
go build -o gh-project-promoter .

# 環境変数の設定（.env ファイルまたは直接指定）
export GH_TOKEN=your_token
export GHPP_OWNER=your_org
export GHPP_PROJECT_NUMBER=1
```

## 環境変数

| 変数名 | 必須 | 説明 |
|---|---|---|
| `GH_TOKEN` | Yes | GitHub API トークン |
| `GHPP_OWNER` | Yes | GitHub Organization / User 名 |
| `GHPP_PROJECT_NUMBER` | Yes | GitHub Projects の番号 |
| `GHPP_STATUS_INBOX` | No | inbox に対応するステータス名 |
| `GHPP_STATUS_PLAN` | No | plan に対応するステータス名 |
| `GHPP_STATUS_READY` | No | ready に対応するステータス名 |
| `GHPP_STATUS_DOING` | No | doing に対応するステータス名 |
| `GHPP_PLAN_LIMIT` | No | 計画フェーズの昇格上限数 |

## ライセンス

MIT
