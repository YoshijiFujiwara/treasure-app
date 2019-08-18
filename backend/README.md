# バックエンドのローカルでの環境構築方法

## 1. まず、API仕様書
```$xslt
/backend/postman_collection.json
```
をPOSTMANでインポートして見てください。

## 2. データベース立ち上げ
/databas ディレクトリに移動して
```$xslt
$ make install
$ docker-compose up -d
```

マイグレーションする
```$xslt
$ make migrate/init
$ make migrate/up
```

マイグレーションし直す(db吹き飛ばして、再構築)
```$xslt
$ make migrate/refresh
```

## 3. APIサーバー
/backend ディレクトリに移動して

.env作成
```$xslt
$ cp .env.example .env
```

立ち上げコマンド
```$xslt
$ make run
```
でいける。

## 4. phpmyadmin
```$xslt
http://localhost:8888
```
パスワードとかは、dockerファイルを読み解く
