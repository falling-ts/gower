# syntax=docker/dockerfile:1
FROM pingcap/tidb:v6.5.1
RUN echo "https://mirrors.aliyun.com/alpine/v3.14/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.14/community/" >> /etc/apk/repositories
RUN apk add --no-cache mysql-client procps
COPY . .
RUN chmod +x tidb-server.sh
RUN chmod +x init.sh
RUN /tidb-server.sh
CMD ["/tidb-server"]
