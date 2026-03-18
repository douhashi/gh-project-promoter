# Promote ライフサイクル

## ステータスフロー

```
inbox → plan → ready → doing
```

GHPP は上記フローのうち、以下2つの昇格を自動化する。

## 1. 計画フェーズ（inbox → plan）

Issue を `inbox` から `plan` ステータスに昇格させる。

### 動作

- `inbox` ステータスの Issue を取得し、`plan` ステータスに変更する
- 昇格した Issue の一覧を JSON で返す

### 制約

- 一度に昇格する個数に上限を設ける（環境変数 `GHPP_PLAN_LIMIT` で上書き可能）

## 2. 実行フェーズ（ready → doing）

Issue を `ready` から `doing` ステータスに昇格させる。

### 動作

- `ready` ステータスの Issue を取得し、`doing` ステータスに変更する

### 制約

- **リポジトリ単位で1つまで**: `doing` に昇格できるのは、各リポジトリにつき1つの Issue のみ
- すでに同リポジトリの Issue が `doing` にある場合、昇格しない
- リポジトリの判定は Issue URL から `owner/repository` を抽出して行う
