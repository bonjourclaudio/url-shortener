FROM golang:latest

RUN go get -u github.com/golang/dep/cmd/dep \
&& rm -rf /go/src/github.com/claudioontheweb \
&&  mkdir /go/src/github.com/claudioontheweb \
&&  git clone https://github.com/claudioontheweb/url-shortener /go/src/github.com/claudioontheweb/url-shortener

LABEL maintainer="Claudio Weckherlin <claudio.weckherlin@gmail.com>"

WORKDIR /go/src/github.com/claudioontheweb/url-shortener

RUN dep ensure -v

RUN go build cmd/*.go

EXPOSE 8080

CMD ["./app"]