FROM golang:alpine

RUN apk add --update git \
 && go get -u -v github.com/constabulary/gb/... \
 && apk del git

ADD . /go/src/github.com/Xe/flurryheart

RUN cd /go/src/github.com/Xe/flurryheart \
 && gb build all \
 && cp ./bin/flurryheart /go/bin

CMD flurryheart
