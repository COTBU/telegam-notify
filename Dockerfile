FROM golang:alpine AS build_base
RUN set -ex && apk update && apk add -q  \
  git unzip build-base autoconf libtool

WORKDIR $GOPATH/src/SOTBI/telegram-notifier

COPY . .

RUN go mod tidy && \
  go build -a -installsuffix cgo -o telegram-notify . && \
  mv telegram-notify /telegram-notify

# Start fresh from a smaller image
FROM alpine:latest

WORKDIR /root/
COPY --from=build_base telegram-notify .
ENTRYPOINT ["./telegram-notify","-config", "/config/notifier.yaml"]
