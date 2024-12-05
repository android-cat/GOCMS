# goバージョン
FROM golang:1.23-alpine
# ワーキングディレクトリの設定
WORKDIR /app
# ホストの go.mod と go.sum をコピー
COPY go.mod go.sum ./
# 依存関係をダウンロード
RUN go mod download
# Goを実行
CMD ["go", "run", "./cmd/go-cms"]
