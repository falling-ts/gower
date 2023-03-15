#!/bin/sh

# 以后台方式启动 tidb-server
nohup /tidb-server >/dev/null 2>&1 &

# 等待 tidb-server（MySQL）启动，您可能需要根据实际情况调整端口和主机名
while ! nc -z 127.0.0.1 4000; do
    echo "Waiting for tidb-server to start..."
    sleep 1
done

# 初始化 SQL
mysql -h 127.0.0.1 -P 4000 -u root < /init.sql
echo "MySQL initialization is complete."

pkill -f tidb-server
