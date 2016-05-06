FROM alpine:latest

RUN apk update && apk add curl git mercurial bzr go && rm -rf /var/cache/apk/*

ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

ADD . /gopath/src/code.olipicus.com/go_rest_api/api
WORKDIR /gopath/src/code.olipicus.com/go_rest_api/api

RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go build -o api .
CMD ["/gopath/src/code.olipicus.com/go_rest_api/api/api"]
