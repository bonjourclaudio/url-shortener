FROM golang:latest


RUN rm -rf /go/src/github.com/claudioontheweb

RUN go get -u github.com/golang/dep/cmd/dep \
&&  mkdir /go/src/github.com/claudioontheweb \
&&  git clone https://github.com/claudioontheweb/url-shortener /go/src/github.com/claudioontheweb/url-shortener

LABEL maintainer="Claudio Weckherlin <claudio.weckherlin@gmail.com>"

WORKDIR /go/src/github.com/claudioontheweb/url-shortener

RUN dep ensure -v

RUN go build cmd/*.go

EXPOSE 8080

CMD ["./app"]