# コーディング規約

## フォーマット

- `gofmt` / `goimports` を適用する
- CI でフォーマットチェックを行う

## Lint

- `golangci-lint` を使用する

## 命名規則

- Go の標準的な命名規則に従う（MixedCaps, 短い変数名）
- パッケージ名は小文字の単語1つ

## エラーハンドリング

- エラーは呼び出し元に返す。握りつぶさない
- `fmt.Errorf` で文脈を付与する: `fmt.Errorf("failed to fetch items: %w", err)`

## テスト

- テストファイルは対象と同じパッケージに配置する（`_test.go`）
- テーブル駆動テストを基本とする
