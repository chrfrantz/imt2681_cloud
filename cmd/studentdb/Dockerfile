FROM golang:1.11.1-stretch as builder

LABEL maintainer "user@mail.com"

RUN go get github.com/chrfrantz/imt2681_cloud/cmd/studentdb

WORKDIR /go/src/github.com/chrfrantz/imt2681_cloud/cmd/studentdb

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o studentdb



FROM scratch

LABEL maintainer="user@mail.com"

WORKDIR /

COPY --from=builder /go/src/github.com/chrfrantz/imt2681_cloud/cmd/studentdb/studentdb .

CMD ["/studentdb"]

