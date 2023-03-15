#!/bin/sh

# 以后台方式启动 tidb-server
nohup /tidb-server >/dev/null 2>&1 &

# 等待 tidb-server（MySQL）启动，您可能需要根据实际情况调整端口和主机名
while ! nc -z 127.0.0.1 4000; do
    echo "Waiting for tidb-server to start..."
    sleep 1
done

# 在此处执行 init.sh，以修改 MySQL 密码
# 请根据实际情况调整 init.sh 的路径
/init.sh
