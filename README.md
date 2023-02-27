# インターン生技術課題

以下のコマンドを順に実行するとGo(echo)とMySQLの環境構築をすることができます。

```
docker-compose build --no-cache
docker-compose up -d
```

## CSVを出力

上記の環境構築が完了した後で以下のコマンド操作に従うと、DBの内容をCSVファイルに出力することができます。

```
git add csv.go

# members.csvという名前で出力されるので、下記コマンドで内容を確認可能
cat members.csv
```