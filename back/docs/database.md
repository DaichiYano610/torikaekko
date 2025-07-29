# データベース (postgreSQL)について


## ユーザデータ
ログインと認証・認可を行うためのテーブル  
id，ユーザ名，ハッシュ化したパスワードを保存
```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username TEXT NOT NULL UNIQUE,
  password BYTEA NOT NULL
);
```
