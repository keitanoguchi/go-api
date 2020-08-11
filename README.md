# go-api

`docker-compose` を使用して、API と DB を 同時起動する

## setup

事前に以下をインストールしておく
- `docker`
  - [インストール方法](https://docs.docker.com/engine/install/)
- `docker-compose`
  - [インストール方法](https://docs.docker.com/compose/install/)

## usage

使用するイメージをビルドする
```sh
$ docker-compose build --no-cache
```

デタッチドモードでコンテナを起動する（DB コンテナの初期化に１分程度かかる）
```sh
$ docker-compose up -d

Creating network "go-api_default" with the default driver
Creating todotweet_db ... done
Creating api          ... done
```

コンテナが起動できているか確認する
```sh
$ docker-compose ps

    Name                 Command             State                 Ports              
--------------------------------------------------------------------------------------
api            ./main                        Up      0.0.0.0:8080->8080/tcp           
todotweet_db   docker-entrypoint.sh mysqld   Up      0.0.0.0:3306->3306/tcp, 33060/tcp
```

API サーバにアクセスする
```sh
$ curl -X GET http://127.0.0.1:8080/

# not found が返ってくれば疎通ができている
{"message":"Not Found"}
```

コンテナを削除する
```sh
$ docker-compose down
Stopping api          ... done
Stopping todotweet_db ... done
Removing api          ... done
Removing todotweet_db ... done
Removing network go-api_default
```