FROM golang:1.9.2-stretch

LABEL maintainer="user@mail.com"

RUN apt-get update && apt-get -y upgrade

RUN go get github.com/chrfrantz/imt2681_cloud/cmd/studentdb

RUN go install github.com/chrfrantz/imt2681_cloud/cmd/studentdb

WORKDIR $GOPATH

CMD ["./bin/studentdb"]

