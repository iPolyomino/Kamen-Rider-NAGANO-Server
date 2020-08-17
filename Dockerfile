FROM golang:alpine as builder
ENV APPDIR $GOPATH/src/github.com/iPolyomino/Kamen-Rider-NAGANO-Server
ENV GO111MODULE on
RUN \
  apk update --no-cache && \
  mkdir -p $APPDIR
ADD . $APPDIR/
WORKDIR $APPDIR
RUN go build --mod=vendor -ldflags "-s -w" -o Kamen-Rider-NAGANO-Server main.go
RUN mv Kamen-Rider-NAGANO-Server /

FROM alpine
RUN apk add --no-cache ca-certificates
RUN apk add mysql-client
COPY --from=builder /Kamen-Rider-NAGANO-Server ./
ENTRYPOINT ["./Kamen-Rider-NAGANO-Server"]