GEN=protoc
ROOT=.
GO=GOOS=linux GOARCH=amd64 go

.PHONY: all
all: build

.PHONY: gen
gen:
	${GEN} -I${ROOT} -I ${ROOT}/api \
	-I /usr/local/include \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	-I ${GOPATH}/src \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:${ROOT}/pkg/pb/example \
	--goclay_out=${ROOT}/pkg/pb/example \
	messages.proto

.PHONY: build
build:
	${GO} build -o ${ROOT}/bin/example ${ROOT}/cmd/example/main.go
	docker build .

test:
	${GO} test ./internal/*