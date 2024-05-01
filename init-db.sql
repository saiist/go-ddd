-- テーブルの全削除
DROP TABLE IF EXISTS users;

-- テーブルの作成
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);