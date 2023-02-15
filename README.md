# インターン生技術課題

以下のコマンドを順に実行するとGo(echo)とMySQLの環境構築をすることができます。

## コンテナ構築編

下記のコマンドを実行すると、それぞれのコンテナを作成することができます。

```
docker-compose build --no-cache
docker-compose up -d
```

## マスターテーブルのレコード編

マスターテーブルについて、下記のコマンドを実行して初期データを与えます。

```
docker cp ./sql internship-task-db-1:/
docker-compose exec db bash
mysql -u root -ppassword
use emonavi_db;
source ./sql/roles.sql;
```