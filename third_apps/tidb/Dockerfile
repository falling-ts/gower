# syntax=docker/dockerfile:1
FROM pingcap/tidb:v6.5.1
COPY config.toml .
COPY init.sql .
ENTRYPOINT ["/tidb-server", "--config=/config.toml"]
