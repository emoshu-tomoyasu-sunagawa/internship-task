# インターン生技術課題

以下のコマンドを順に実行するとGo(echo)とMySQLの環境構築をすることができます。

```
docker-compose build --no-cache
docker-compose up -d
docker cp ./sql 729e23709bec:/
docker-compose exec db bash
mysql -u root -ppassword
source ./sql/roles.sql
```