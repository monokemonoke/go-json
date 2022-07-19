# GO-JSON

これは Golang で書かれた JSON パーサです.
JSON文字列を `map[string]interface{}` に変換します.

## How To Run

1. そのまま実行するとサンプルケースのパース結果を確認できます.
```bash
$ go run main.go
```

2. テストを確認すると, 内部のメソッドが何をしているのかわかりやすいです.
```bash
$ go test -v ./...
```

## Change log
- 2022/07/20
    - とりあえず string, int, bool をパースできるようにした
    - テストが貧弱
    - コードが汚いので改善したい