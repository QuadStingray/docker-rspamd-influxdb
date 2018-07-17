FROM golang:alpine AS build-env

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

RUN apk update && apk upgrade && \
apk add --no-cache bash git openssh

# build directories
ADD . /go/src/quadstingray/rspamd-influxdb
WORKDIR /go/src/quadstingray/rspamd-influxdb

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Build my app
RUN go build -o rspamdInfluxDB *.go

# final stage
FROM alpine
WORKDIR /app

MAINTAINER QuadStingray <rspamd-influxdb@quadstingray.com>

ENV INTERVAL=3600 \
    INFLUXDB_DB="rspamd" \
    INFLUXDB_URL="http://influxdb:8086" \
    INFLUXDB_USER="DEFAULT" \
    INFLUXDB_PWD="DEFAULT" \
    RSPAMD_URL="https://RSPAMD-URL" \
    RSPAMD_PASSWORD="PASSWORD"

RUN apk add ca-certificates
COPY --from=build-env /go/src/quadstingray/rspamd-influxdb/rspamdInfluxDB /app/rspamdInfluxDB
ADD run.sh run.sh
CMD sh run.sh