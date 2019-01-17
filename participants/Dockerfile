FROM golang

RUN mkdir -p /go/src/angadsharma1016/omega

ADD . /go/src/angadsharma1016/omega

WORKDIR /go/src/angadsharma1016/omega/participants

RUN go get -t -v ./...
RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run angadsharma1016/omega/participants/cmd  -watch angadsharma1016/omega/participants
