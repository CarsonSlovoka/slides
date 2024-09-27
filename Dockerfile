# docker build -t slides:v0.2.0.alpha .
# docker run -p 8080:8651 --name containers的名稱 實際image名稱:該image使用的tag名稱
# docker run -p 8080:8651 --name slidesDemo slides:v0.2.0.alpha

ARG GO_VERSION=1.23.0
FROM golang:${GO_VERSION} AS build

# 工作目錄不存在會自己建立
WORKDIR /slides

# .dockerignore 寫好，就可以直接COPY . .
COPY . .
# COPY go.mod go.sum ./
# COPY slides.gohtml ./
# COPY *.go .
# COPY app/*.go ./app/
# COPY assets/* ./assets/
# COPY internal/ ./internal/
# COPY md/* ./md/
# COPY help.md ./help.md
# COPY plugin/ ./plugin/
# COPY reveal.js/ ./reveal.js/
RUN go mod tidy
RUN go mod download -x

# 加上GOARCH可能會失敗
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o slides .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -tags tmpl -o slides .
# /slides/slides.exe

# 如果省略了執行環境，只有基底映像: golang:${GO_VERSION}，那麼出來的Image size會相當大，數百MB

# 加上執行環境，這使得最終的映像只包含Alpine基本系統文件、和我們所編譯的slides執行檔，Image size只有幾十MB而已
# 選擇執行環境
FROM alpine:latest

# WORKDIR /bin # 這通常放系統執行檔

WORKDIR /usr/local/bin

# 將build流程中的 /slides/slides.exe 複製到 當前目錄(也就是/usr/local/bin)
COPY --from=build /slides/slides .
