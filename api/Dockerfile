FROM golang:alpine

RUN apk update && apk add git && rm -rf /var/cache/apk/*

ADD . /go/src/code.olipicus.com/go_rest_api/api
WORKDIR /go/src/code.olipicus.com/go_rest_api/api

RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go build -o api .
CMD ["/go/src/code.olipicus.com/go_rest_api/api/api"]
