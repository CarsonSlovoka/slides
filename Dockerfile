# docker build -t slides:v0.2.0.alpha .
# docker run -p 8080:8651 --name containers的名稱 實際image名稱:該image使用的tag名稱
# docker run -p 8080:8651 --name slidesDemo slides:v0.2.0.alpha

ARG GO_VERSION=1.23.0
FROM golang:${GO_VERSION} AS build

# 工作目錄不存在會自己建立
WORKDIR /slides

# COPY . . 全部複製
COPY go.mod go.sum ./
COPY slides.gohtml ./
COPY *.go .
COPY app/*.go ./app/
COPY assets/* ./assets/
COPY internal/ ./internal/
COPY md/* ./md/
COPY help.md ./help.md
COPY plugin/ ./plugin/
COPY reveal.js/ ./reveal.js/
RUN go mod tidy
RUN go mod download -x

# 加上GOARCH可能會失敗
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o slides .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -tags tmpl -o slides .
# /slides/slides.exe

# 選擇執行環境
FROM alpine:latest

WORKDIR /bin

# 將build流程中的 /slides/slides.exe 複製到 當前目錄(也就是bin)
COPY --from=build /slides/slides .
# bin/slides.exe

# COPY --from=build /slides/md/ . # 這是錯的，複製資料夾後面要加上目錄路徑，不然複製不過去
COPY --from=build /slides/md/ ./md/
# bin/md

COPY --from=build /slides/assets/ ./assets/

# docker run -p 8080:8651 --name slidesDemo slides:v0.2.0.alpha
# 曝露給外部的port是8080, 而docker裡面用的是8651
CMD ["./slides", "-md", "md", "-port", "8651", "-host", ""]
