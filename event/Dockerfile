FROM golang

RUN mkdir -p /go/src/angadsharma1016/omega

ADD . /go/src/angadsharma1016/omega

WORKDIR /go/src/angadsharma1016/omega/event

RUN go get -t -v ./...

EXPOSE 3000

CMD ["go", "run", "main.go"]
