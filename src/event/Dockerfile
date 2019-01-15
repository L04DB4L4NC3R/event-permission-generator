FROM golang

RUN mkdir -p /go/src/angadsharma1016/omega/src

ADD . /go/src/angadsharma1016/omega/src

WORKDIR /go/src/angadsharma1016/omega

RUN go get -t -v ./...

EXPOSE 3000

CMD ["go", "run", "src/event/main.go"]
