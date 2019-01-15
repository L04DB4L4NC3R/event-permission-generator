FROM golang

RUN mkdir -p /go/src/angadsharma1016/omega/src

ADD . /go/src/angadsharma1016/omega/src

WORKDIR /go/src/angadsharma1016/omega

RUN go get  -t -v ./...
RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run angadsharma1016/omega/src/participants/cmd -watch angadsharma1016/omega/src/participants
