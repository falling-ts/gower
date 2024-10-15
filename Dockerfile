# syntax=docker/dockerfile:1
FROM golang:1.23.2
ENV TZ=Asia/Shanghai
RUN mkdir -p "/go/bin"
WORKDIR /go/bin
COPY gower .
EXPOSE 8080
RUN chmod +x gower
CMD ["./gower", "run"]
