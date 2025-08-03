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

## あげるものデータ
```sql
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    want TEXT NOT NULL,
    image_paths TEXT[] NOT NULL, -- 複数の画像パス(カンマ区切り)
    user_id INTEGER NOT NULL REFERENCES users(id)
);
```
