# docker build -t slides:v0.2.0 -f Dockerfile ..

ARG GO_VERSION=1.23.0
FROM golang:${GO_VERSION} AS build

WORKDIR /slides

COPY . .
RUN go mod tidy
RUN go mod download -x
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -tags tmpl -o slides .
# /slides/slides.exe

FROM alpine:latest

WORKDIR /usr/local/bin

# copy form build slides/slides.exe => /usr/local/bin
COPY --from=build /slides/slides .
