# ait-guide-back

![Static Badge](https://img.shields.io/badge/Go-1.21.6-_?logo=go&link=https%3A%2F%2Fgo.dev%2F)
![Static Badge](https://img.shields.io/badge/Gin-1.9.1-_?logo=gin&link=https%3A%2F%2Fpkg.go.dev%2Fgithub.com%2Fgin-gonic%2Fgin)
![Static Badge](https://img.shields.io/badge/Viper-1.18.2-_?link=https%3A%2F%2Fpkg.go.dev%2Fgithub.com%2Fspf13%2Fviper)
![Static Badge](https://img.shields.io/badge/Docker-25.0.3-_?logo=docker&link=https%3A%2F%2Fdocs.docker.com%2F)
![Static Badge](https://img.shields.io/badge/MySQL-8-_?logo=mysql)
![Static Badge](https://img.shields.io/badge/Neo4j-5.17.0-_?logo=neo4j&link=https%3A%2F%2Fneo4j.com%2Fdocs%2Fgo-manual%2Fcurrent%2F)

[AITガイド](https://ait-guide.sysken.net)のバックエンドリポジトリ

## 環境

| 言語・フレームワーク  | バージョン |
| --------------------- | ---------- |
| Go                    | 1.21.6     |
| gin-gonic/gin         | 1.9.1      |
| neo4j-go-driver/v5    | 5.18.0     |
| paulmach/orb          | 0.11.1     |
| spf13/viper           | 1.18.2     |
| MySQL                 | 8.0        |
| Neo4j                 | 5.17.0     |
| Docker                | 25.0.3     |

## makeコマンド

| Make                | 実行する処理               |
| ------------------- | ------------------------ |
| make                | ヘルプの表示               |
| make build          | dockerコンテナビルド       |
| make up             | dockerコンテナ起動         |
| make build-up       | dockerコンテナビルド＆起動  |
| make down           | dockerコンテナ停止         |
| make restart        | コンテナ再起動             |
| make db             | DBコンテナに入る           |
| make go             | goコンテナに入る           |
| make help           | ヘルプ                    |

## 環境変数

シス研のesaを参照

## 実行方法

### 1.環境構築

Dockerをインストールする
[Docker Desktop の Mac へのインストール](https://matsuand.github.io/docs.docker.jp.onthefly/desktop/mac/install/)\
or\
`brew install --cask docker`

[Docker Engine インストール（Ubuntu 向け）](https://matsuand.github.io/docs.docker.jp.onthefly/engine/install/ubuntu/)

### 2.環境変数のファイルを格納する

プロジェクトのルートディレクトリと`./src/conf/environments`にそれぞれファイルを設定
環境変数のファイルについてはesaを確認すること

### 3.実行

makefileがあるのでそれを使う

#### ビルド＆実行

``` bash
make build-up
```
