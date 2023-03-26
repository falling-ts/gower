# syntax=docker/dockerfile:1
FROM golang:1.20.2
ENV GO111MODULE=on
RUN mkdir -p "$GOPATH/src/gower"
WORKDIR /go/src/gower
COPY . .
RUN go mod tidy
RUN go test
RUN go test -bench=Benchmark
RUN go install
EXPOSE 8080
CMD gower

