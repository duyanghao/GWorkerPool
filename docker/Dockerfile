# Building stage
FROM golang:1.11-alpine3.9

RUN apk add --update make

WORKDIR /go/src/github.com/duyanghao/GWorkerPool

# Source code, building tools and dependences
COPY cmd /go/src/github.com/duyanghao/GWorkerPool/cmd
COPY worker /go/src/github.com/duyanghao/GWorkerPool/worker
COPY Makefile /go/src/github.com/duyanghao/GWorkerPool

COPY go.mod /go/src/github.com/duyanghao/GWorkerPool
COPY go.sum /go/src/github.com/duyanghao/GWorkerPool
COPY vendor /go/src/github.com/duyanghao/GWorkerPool/vendor

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE=on

ENV TIMEZONE "Asia/Shanghai"

RUN make src.build
RUN mv build/GWorkerPool/GWorkerPool /go/bin

# Production stage
FROM alpine:3.9
WORKDIR /go/bin

# copy the go binaries from the building stage
COPY --from=0 /go/bin /go/bin

# copy the config files from the current working dir
COPY examples /GWorkerPool/examples

ENTRYPOINT ["./GWorkerPool", "-config", "/GWorkerPool/examples/simple.yml", "-v", "5", "-logtostderr", "true"]
