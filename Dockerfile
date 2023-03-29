# syntax=docker/dockerfile:1
FROM golang:1.20.2
ENV TZ=Asia/Shanghai
RUN mkdir -p "$GOPATH/bin"
WORKDIR $GOPATH/bin
COPY gower .
EXPOSE 8080
RUN chmod +x gower
CMD ["./gower", "run"]
