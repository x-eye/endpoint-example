FROM golang

ADD . /go/src/example
WORKDIR /go/src/example

RUN GO111MODULE=on go mod vendor && GO111MODULE=off go build -o ./bin/example ./cmd/example/main.go

CMD bin/example