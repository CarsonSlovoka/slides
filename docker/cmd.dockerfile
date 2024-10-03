# docker build -t slides-cmd:v0.2.0 -f cmd.dockerfile ..
# docker run -p 8080:8080 -e PORT=8080 -e MD_DIR="/docs" -e FS_DIRS="/pages,/static,/tmpl" --name SlidesCmd slides-cmd:v0.2.0

ARG GO_VERSION=1.23.0
FROM golang:${GO_VERSION} AS build
WORKDIR /slides
COPY . .
RUN go mod tidy
RUN go mod download -x
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -tags tmpl -o slides .


FROM alpine:latest
WORKDIR /usr/local/bin
COPY --from=build /slides/slides .
COPY --from=build /slides/docker/cmd_entrypoint.sh ./cmd_entrypoint.sh

# 取得工具bash和envsubset 才能認得sh腳本 #!/bin/bash
# alpine預設好像只能 #!/bin/sh
RUN apk add --no-cache bash

# 設置執行權限
# RUN chmod +x ./cmd_entrypoint.sh

ENTRYPOINT ["./cmd_entrypoint.sh"]
