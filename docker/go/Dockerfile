# goバージョン
FROM golang:1.19.1-alpine

# アップデートとgit、alpine-sdkのインストール
RUN apk update && apk add git && apk add alpine-sdk

# 環境変数の設定
ENV APP_ROOT /app

# ディレクトリの作成
RUN mkdir $APP_ROOT

# ワーキングディレクトリの設定
WORKDIR $APP_ROOT

# ルートディレクトリを作業ディレクトリにコピー
COPY . $APP_ROOT

EXPOSE 8080
